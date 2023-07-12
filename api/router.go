package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register", Register)
		user.GET("/login", Login)
		user.GET("/login/refresh", RefreshToken)
	}
	room := r.Group("/room")
	{
		room.POST("/create", CreateRoom)
		room.GET("/connect", ConnectRoom)
	}
	r.Run()
}
