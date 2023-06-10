package api

import (
	"Chess/service"
	"Chess/util/resps"
	"github.com/gin-gonic/gin"
	"log"
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
}
