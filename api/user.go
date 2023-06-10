package api

import (
	"Chess/service"
	"Chess/util/resps"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		resps.NormErr(c, 10000, "未输入账号密码")
		return
	}
	err := service.Register(username, password)
	if err != nil {
		resps.NormErr(c, 400, err.Error())
		log.Println(err)
		return
	}
	resps.RespOK(c)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	token, refreshToken, err := service.Login(username, password)
	if err != nil {
		resps.NormErr(c, 400, err.Error())
		log.Println(err)
		return
	}
	log.Println(token, refreshToken)
	resps.RespOK(c)
}

func RefreshToken(c *gin.Context) {
	refreshToken := c.Query("refresh_token")
	token, refreshToken, err := service.CheckRefreshTokenAndReturnToken(refreshToken)
	if err != nil {
		resps.NormErr(c, 400, err.Error())
		log.Println(err)
		return
	}
	log.Println(token, refreshToken)
	resps.RespOK(c)
}
