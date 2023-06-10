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
	userClient.Register = make(chan *model.Client)
	userClient.Unregister = make(chan *model.Client)
	userClient.User = user
	userClient.Room = room
	userClient.IsReady = false
	//用户填入房间
	if room.PlayerA != (model.UserClient{}) {
		userClient.Color = model.White
		room.PlayerA = userClient
	} else if room.PlayerB != (model.UserClient{}) {
		userClient.Color = model.Black
		room.PlayerB = userClient
	} else {
		return errors.New("房间对战用户已满")
	}
	go control(&userClient)
	userClient.Register <- &client
	return nil
}

// 控制各个客户端
func control(user *model.UserClient) {
	for {
		select {
		case client := <-user.Register:
			user.Client.Store(client, true)
		case client := <-user.Unregister:
			//TODO:注销
		case message := <-user.Broadcast:
			//TODO:传播消息
		}
	}
}

func Read(c *model.Client) {
	defer func() {
		log.Println("好似")
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
				} else {
					c.UserClient.IsReady = false
					c.UserClient.Room.ReadyNum -= 1
				}
			case 2:
				if c.UserClient.Room.ReadyNum != 2 {
					log.Println("局都还没开")
					continue
				}
				//TODO：检查合法性
			}
		default:
			log.Println("不支持的消息类型")
			continue
		}
	}
}

func Write(c *model.Client) {

}
