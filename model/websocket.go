package model

import "github.com/gorilla/websocket"

type WebsocketMessage struct {
	Type       int    `json:"type"`        //消息类型。1表示准备（重复发送退出准备状态），2表示移动，3表示认输（和棋），4为升变
	MoveBefore [2]int `json:"move_before"` //移动前棋子所在的位置。两个位置分别为x、y轴。
	MoveAfter  [2]int `json:"move_after"`  //移动后棋子所在的位置
	Upgrade    int    `json:"upgrade"`     //升变，参考常量。升变条件由系统进行检测。
}

type Client struct {
	Conn       *websocket.Conn
	PlayerName string
	Send       chan []byte
	UserClient *UserClient
}
