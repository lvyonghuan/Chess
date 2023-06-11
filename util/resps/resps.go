package resps

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var ok = respTemplate{
	Status: 200,
	Info:   "success",
}

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, ok)
}

func NormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}

func RespRoomNum(c *gin.Context, roomID int) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   roomID,
	})
}

func RespToken(c *gin.Context, token, refreshToken string) {
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func NormResp(c *gin.Context, any interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   any,
	})
}
