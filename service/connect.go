package service

import (
	"Chess/model"
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
	var userClient model.UserClient
	userClient.Client.Store(client, true)
	room.UserClient = append(room.UserClient, userClient)
	return nil
}
