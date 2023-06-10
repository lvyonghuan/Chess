package model

import (
	"sync"
)

type UserClient struct {
	Client     *sync.Map
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	User       User
	IsReady    bool
	Room       *Room //我感觉这一步多少有点......不好评价
}

type Room struct {
	ID            int        `json:"id"`
	RoomName      string     `json:"roomName" `
	PlayerA       UserClient `json:"playerA"`
	PlayerB       UserClient `json:"playerB"`
	Checkerboard  *Chess     `json:"checkerboard"`
	ReadyNum      int        `json:"readNum"`
	ViewersClient []UserClient
}
