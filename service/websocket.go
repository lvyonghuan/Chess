package service

import (
	"Chess/database"
	"Chess/model"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ConnectRoom(token string, roomID int, c *gin.Context) error {
	err, userID := CheckExp(token, tokenSecret)
	if err != nil {
		return err
	}
	user, err := database.FindUserByUid(userID)
	if err != nil {
		log.Println("数据库查询失败,", err)
		return err
	}
	room := RoomMap[roomID]
	conn, err := Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("websocket升级失败:", err)
		c.JSON(http.StatusInternalServerError, "websocket协议升级失败")
		return err
	}
	client := model.Client{
		Conn: conn,
		Send: make(chan []byte, 1024),
	}
	var userClient model.UserClient
	client.UserClient = &userClient
	client.UserClient.Client = &sync.Map{}
	userClient.Register = make(chan *model.Client)
	userClient.Unregister = make(chan *model.Client)
	userClient.Broadcast = make(chan []byte)
	userClient.User = user
	userClient.Room = room
	userClient.IsReady = false
	//用户填入房间
	if room.PlayerA == (model.UserClient{}) {
		userClient.Color = model.White
		room.PlayerA = userClient
	} else if room.PlayerB == (model.UserClient{}) {
		userClient.Color = model.Black
		room.PlayerB = userClient
	} else {
		return errors.New("房间对战用户已满")
	}
	if room.ViewersClient == nil {
		room.ViewersClient = make([]model.UserClient, 0)
	}
	room.ViewersClient = append(room.ViewersClient, userClient)
	go control(&userClient)
	userClient.Register <- &client
	go Read(&client)
	go Write(&client)
	return nil
}

// 控制各个客户端
func control(user *model.UserClient) {
	for {
		select {
		case client := <-user.Register:
			user.Client.Store(client, true)
		case client := <-user.Unregister:
			client.UserClient.Room.ReadyNum -= 1
			if *client.UserClient == client.UserClient.Room.PlayerA {
				client.UserClient.Room.PlayerA = model.UserClient{}
			} else {
				client.UserClient.Room.PlayerB = model.UserClient{}
			}
			err := client.Conn.Close()
			if err != nil {
				log.Println(err)
			}
		case message := <-user.Broadcast:
			for _, client := range user.Room.ViewersClient {
				client.Client.Range(func(key, value interface{}) bool {
					client, ok := key.(*model.Client)
					if !ok {
						log.Println("Failed to convert key to Client:", key)
					}
					client.Send <- message
					return true
				})
			}
		}
	}
}

func Read(c *model.Client) {
	defer func() {
		log.Println("好似")
		c.UserClient.Unregister <- c
	}()
	for {
		msgType, msgByte, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read msg failed,err info:", err)
			break
		}
		switch msgType {
		case websocket.TextMessage:
			var msg model.WebsocketMessage
			err := json.Unmarshal(msgByte, &msg)
			if err != nil {
				log.Println("信息处理失败")
				continue
			}
			//消息类型。1表示准备（重复发送退出准备状态），2表示移动，3表示认输（和棋）
			switch msg.Type {
			case 1:
				if !c.UserClient.IsReady {
					c.UserClient.IsReady = true
					c.UserClient.Room.ReadyNum += 1
					marshal, err := json.Marshal(c.UserClient.User.Name + "准备好了")
					if err != nil {
						log.Println("that's messed up,", err)
					}
					c.UserClient.Broadcast <- marshal
				} else {
					c.UserClient.IsReady = false
					c.UserClient.Room.ReadyNum -= 1
					marshal, err := json.Marshal(c.UserClient.User.Name + "退出了准备状态")
					if err != nil {
						log.Println("that's messed up,", err)
					}
					c.UserClient.Broadcast <- marshal
				}
			case 2:
				if c.UserClient.Room.ReadyNum != 2 {
					log.Println("局都还没开")
					continue
				}
				if c.UserClient.Room.NextStep != c.UserClient.Color {
					log.Println("急什么")
					continue
				}
				isLegitimate, errStr := checkMove(c, msg)
				if !isLegitimate {
					log.Println(errStr)
					continue //避免顺路执行下去了
				}
				move(msg, c)
				if c.UserClient.Room.Upgrade.IsUpgrade {
					c.UserClient.Room.NextStep = c.UserClient.Color
				} else {
					if c.UserClient.Color == model.White {
						c.UserClient.Room.NextStep = model.Black
					} else {
						c.UserClient.Room.NextStep = model.White
					}
				}
				log.Println(c.UserClient.Room.Checkerboard.Checkerboard) //打印不可名状的地图到cmd
				marshal, err := json.Marshal(c.UserClient.Room.Checkerboard.Checkerboard)
				if err != nil {
					log.Println("that's messed up,", err)
				}
				c.UserClient.Broadcast <- marshal
				if !CheckWin(c) {
					marshal, err := json.Marshal(c.UserClient.User.Name + "赢得了本局")
					if err != nil {
						log.Println("that's messed up,", err)
					}
					c.UserClient.Broadcast <- marshal
				}
			case 3:
				//TODO：认输
			case 4:
				if c.UserClient.Room.NextStep != c.UserClient.Color || !c.UserClient.Room.Upgrade.IsUpgrade {
					log.Println("升个锤子")
					continue
				}
				upgrade(msg, c)
				log.Println(c.UserClient.Room.Checkerboard.Checkerboard) //打印不可名状的地图到cmd
				marshal, err := json.Marshal(c.UserClient.Room.Checkerboard.Checkerboard)
				if err != nil {
					log.Println("that's messed up,", err)
				}
				c.UserClient.Broadcast <- marshal
				if !CheckWin(c) {
					marshal, err := json.Marshal(c.UserClient.User.Name + "赢得了本局")
					if err != nil {
						log.Println("that's messed up,", err)
					}
					c.UserClient.Broadcast <- marshal
				}
			}
		default:
			log.Println("不支持的消息类型")
			continue
		}
	}
}

func Write(c *model.Client) {
	defer func() {
		//我寻思read读到应该就行了
	}()
	for {
		select {
		case msg := <-c.Send:
			err := c.Conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("write message error,that's messed up,", err)
				return
			}
		}
	}
}
