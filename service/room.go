package service

import (
	"Chess/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var RoomMap map[int]*model.Room

func CrateRoom(room *model.Room) {
	room.ID = generateRoomID()
	initCheckerBoard(&room.Checkerboard)
	RoomMap[room.ID] = room //将房间加入map
}

func ConnectRoom(room *model.Room, c *gin.Context) error {
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
	room.UserClient[0].Client.Store(client, true)
	return nil
}

// 随机生成房间号 虽然撞的概率很小很小，但是我懒得管了，写完再说
func generateRoomID() (id int) {
	rand.Seed(time.Now().Unix())
	id = rand.Intn(90000000) + 10000000
	return id
}

// 特别抽象的初始化棋盘
func initCheckerBoard(checkerBoard *model.Chess) {
	for i := 0; i < 8; i++ {
		//初始化走卒
		if i == 1 || i == 6 {
			for j := 0; j < 8; j++ {
				checkerBoard.Checkerboard[i][j][0] = model.Pawn
				if i == 1 {
					checkerBoard.Checkerboard[i][j][1] = model.White
				} else {
					checkerBoard.Checkerboard[i][j][1] = model.Black
				}
			}
		} else if i == 0 || i == 7 { //初始化特殊
			for j := 0; j < 8; j++ {
				if j == 0 || j == 7 {
					checkerBoard.Checkerboard[i][j][0] = model.Rook
				} else if j == 1 || j == 6 {
					checkerBoard.Checkerboard[i][j][0] = model.Bishop
				} else if j == 2 || j == 5 {
					checkerBoard.Checkerboard[i][j][0] = model.Knight
				} else if j == 4 {
					if i == 0 {
						checkerBoard.Checkerboard[i][j][0] = model.Queen
					} else {
						checkerBoard.Checkerboard[i][j][0] = model.King
					}
				} else if j == 5 {
					if i == 7 {
						checkerBoard.Checkerboard[i][j][0] = model.Queen
					} else {
						checkerBoard.Checkerboard[i][j][0] = model.King
					}
				}
				if i == 0 {
					checkerBoard.Checkerboard[i][j][1] = model.White
				} else {
					checkerBoard.Checkerboard[i][j][1] = model.Black
				}
			}
		} else {
			for j := 0; j < 8; j++ { //初始化空格
				checkerBoard.Checkerboard[i][j][0] = model.Void
				checkerBoard.Checkerboard[i][j][1] = model.Void
			}
		}
	}
}
