package api

import (
	"Chess/service"
	"Chess/util/resps"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CreateRoom(c *gin.Context) {
	token := c.GetHeader("Authorization")
	roomName := c.PostForm("room_name")
	num, err := service.InitRoom(token, roomName, c)
	if err != nil {
		log.Println(err)
		resps.NormErr(c, 400, err.Error())
	}
	log.Println(num)
	resps.RespOK(c)
}

// ConnectRoom 建立websocket链接
func ConnectRoom(c *gin.Context) {
	token := c.GetHeader("Authorization")
	roomIDStr := c.Query("room_id")
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		log.Println(err)
		resps.NormErr(c, 400, err.Error())
	}
	err = service.ConnectRoom(token, roomID, c)
	if err != nil {
		log.Println(err)
		resps.NormErr(c, 400, err.Error())
	}
}
