package service

import (
	"Chess/model"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
)

var RoomMap map[int]*model.Room

func InitRoom(token, roomName string, c *gin.Context) (roomNumber int, err error) {
	err, _ = CheckExp(token, tokenSecret)
	if err != nil {
		return 0, err
	}
	var room model.Room
	room.RoomName = roomName
	room.PlayerA = model.UserClient{}
	room.PlayerB = model.UserClient{}
	room.ReadyNum = 0
	crateRoom(&room)
	//白棋先走
	room.NextStep = model.White
	return room.ID, nil
}

func crateRoom(room *model.Room) {
	room.ID = generateRoomID()
	room.Checkerboard = &model.Chess{}
	initCheckerBoard(room.Checkerboard)
	if RoomMap == nil { // 如果RoomMap是nil，则先进行初始化
		RoomMap = make(map[int]*model.Room)
	}
	RoomMap[room.ID] = room //将房间加入map
}

// 随机生成房间号 虽然撞的概率很小很小，但是我懒得管了，写完再说
func generateRoomID() (id int) {
	rand.Seed(time.Now().Unix())
	id = rand.Intn(90000000) + 10000000
	return id
}

// 特别抽象的初始化棋盘
func initCheckerBoard(checkerBoard *model.Chess) {
	for i := 0; i < 8; i++ {
		//初始化走卒
		if i == 1 || i == 6 {
			for j := 0; j < 8; j++ {
				checkerBoard.Checkerboard[i][j][0] = model.Pawn
				if i == 1 {
					checkerBoard.Checkerboard[i][j][1] = model.White
				} else {
					checkerBoard.Checkerboard[i][j][1] = model.Black
				}
			}
		} else if i == 0 || i == 7 { //初始化特殊
			for j := 0; j < 8; j++ {
				if j == 0 || j == 7 {
					checkerBoard.Checkerboard[i][j][0] = model.Rook
					checkerBoard.Checkerboard[i][j][4] = 0
				} else if j == 1 || j == 6 {
					checkerBoard.Checkerboard[i][j][0] = model.Knight
				} else if j == 2 || j == 5 {
					checkerBoard.Checkerboard[i][j][0] = model.Bishop
				} else if j == 3 {
					checkerBoard.Checkerboard[i][j][0] = model.Queen
				} else if j == 4 {
					checkerBoard.Checkerboard[i][j][0] = model.King
					checkerBoard.Checkerboard[i][j][4] = 0
				}
				if i == 0 {
					checkerBoard.Checkerboard[i][j][1] = model.White
				} else {
					checkerBoard.Checkerboard[i][j][1] = model.Black
				}
			}
		} else {
			for j := 0; j < 8; j++ { //初始化空格
				checkerBoard.Checkerboard[i][j][0] = model.Void
				checkerBoard.Checkerboard[i][j][1] = model.Void
			}
		}
	}
	//初始化棋盘威胁度
	CalculateThreaten(model.White, checkerBoard)
	CalculateThreaten(model.Black, checkerBoard)
	log.Println(checkerBoard)
}
