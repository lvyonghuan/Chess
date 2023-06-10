package model

import "github.com/gorilla/websocket"

type WebsocketMessage struct {
	Type int   `json:"type"` //消息类型。1表示准备（重复发送退出准备状态），2表示移动，3表示认输（和棋）
	Move Chess `json:"move"` //复刻一份棋盘
}

type Client struct {
	Conn       *websocket.Conn
	PlayerName string
	Send       chan []byte
	UserClient *UserClient
}
