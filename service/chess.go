package service

import (
	"Chess/model"
	"log"
)

// 检查着点的合法性
func checkMove(c *model.Client, msg model.WebsocketMessage) (bool, string) {
	checkerBoard := c.UserClient.Room.Checkerboard
	x1 := msg.MoveBefore[0]
	y1 := msg.MoveBefore[1]
	x2 := msg.MoveAfter[0]
	y2 := msg.MoveAfter[1]
	piece := checkerBoard.Checkerboard[x1][y1][0]
	color := checkerBoard.Checkerboard[x1][y1][1]
	if color != c.UserClient.Color {
		return false, "棋色都错了"
	}
	//边界检查
	if x2 > 7 || y2 > 7 || x2 < 0 || y2 < 0 {
		return false, "边界都超了"
	}
	switch piece {
	case model.King:
		return checkKingMove(x1, y1, x2, y2, color, checkerBoard)
	case model.Queen:

	case model.Rook:

	case model.Bishop:

	case model.Knight:

	case model.Pawn:

	default:
		log.Println("什么鬼")
		return false, "什么鬼"
	}
	return true, ""
}

// 合法性检查大全
// 国王检查
func checkKingMove(x1, y1, x2, y2, color int, checkerBoard *model.Chess) (bool, string) {
	var flag int
	if color == model.White {
		flag = 2
	} else {
		flag = 3
	}
	//TODO:威胁检查器
	if checkerBoard.Checkerboard[x2][y2][flag] == 1 {
		return false, "王有威胁"
	}
	//TODO：王车易位判断

	if checkerBoard.Checkerboard[x2][y2][1] == color {
		return false, "怎么还吃自己子的？"
	}
	return true, ""
}

func checkQueenMove(x1, y1, x2, y2 int) (bool, string) {
	return true, ""
}

// 对马撇脚的计算
func isBlocked(x, y int, offset []int, checkerBoard *model.Chess) bool {
	blockerX, blockerY := 0, 0

	if offset[0] > 1 {
		blockerX = x + 1
	} else if offset[0] < -1 {
		blockerX = x - 1
	} else {
		blockerX = x
	}

	if offset[1] > 1 {
		blockerY = y + 1
	} else if offset[1] < -1 {
		blockerY = y - 1
	} else {
		blockerY = y
	}

	if blockerX >= 0 && blockerX <= 7 && blockerY >= 0 && blockerY <= 7 {
		return checkerBoard.Checkerboard[blockerX][blockerY][0] != 0
	}
	return false
}
