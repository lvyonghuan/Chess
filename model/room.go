package model

import (
	"sync"
)

type UserClient struct {
	Client     *sync.Map
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	User       User //用户信息，目前没用，留着将来（迫真）用
	IsReady    bool
	Color      int   //参照颜色const。playerA执白旗，B执黑旗。
	Room       *Room //我感觉这一步多少有点......不好评价
}

type Room struct {
	ID            int        `json:"id"`
	RoomName      string     `json:"roomName" `
	PlayerA       UserClient `json:"playerA"`
	PlayerB       UserClient `json:"playerB"`
	NextStep      int        `json:"nextStep"` //1表示下一步A下，2表示下一步B下
	Checkerboard  *Chess     `json:"checkerboard"`
	ReadyNum      int        `json:"readNum"`
	ViewersClient []UserClient
}
