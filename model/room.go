package model

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	Conn       *websocket.Conn
	PlayerName string
	Send       chan []byte
	UserClient *UserClient
}

type UserClient struct {
	Client     *sync.Map
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

type Room struct {
	ID           int    `json:"id"`
	RoomName     string `json:"roomName" `
	PlayerA      string `json:"playerA"`
	PlayerB      string `json:"playerB"`
	Checkerboard Chess  `json:"checkerboard"`
	UserClient   []UserClient
}
