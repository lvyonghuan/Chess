package api

import (
	"Chess/service"
	"Chess/util/resps"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateRoom(c *gin.Context) {
	token := c.GetHeader("Authorization")
	err := service.CheckExp(token)
	if err != nil {
		log.Println(err)
		resps.NormErr(c, 400, err.Error())
	}
}
