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
		return checkQueenMove(x1, y1, x2, y2, color, checkerBoard)
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
	if checkerBoard.Checkerboard[x2][y2][flag] == 1 {
		return false, "王有威胁"
	}

	//王车易位
	if checkerBoard.Checkerboard[x2][y2][0] == model.Rook && checkerBoard.Checkerboard[x2][y2][1] == color && checkerBoard.Checkerboard[x1][y1][4] == 0 && checkerBoard.Checkerboard[x2][y2][4] == 0 && checkerBoard.Checkerboard[x2][y2][flag] == 0 && checkerBoard.Checkerboard[x1][x1][flag] == 0 {
		direction := 1 //易位方向判断
		if y2 < y1 {
			direction = -1
		}
		y := y1 + direction
		for ; y != y1; y += direction {
			if checkerBoard.Checkerboard[x1][y][0] != 0 || checkerBoard.Checkerboard[x1][y][flag] != 0 {
				return false, "易个锤子"
			}
		}
		return true, ""
	} else if checkerBoard.Checkerboard[x2][y2][0] == model.Rook {
		return false, "易个锤子"
	}

	if checkerBoard.Checkerboard[x2][y2][1] == color {
		return false, "怎么还吃自己子的？"
	}
	return true, ""
}

func checkQueenMove(x1, y1, x2, y2, color int, checkerBoard *model.Chess) (bool, string) {
	var flag = returnFlag(color)
	dx := x2 - x1
	dy := y2 - y1

	if dx == 0 || dy == 0 || abs(dx) == abs(dy) { // 判断是否沿水平、垂直或对角线移动
		stepX := 0
		if dx != 0 {
			stepX = dx / abs(dx) // 确定 x 方向
		}

		stepY := 0
		if dy != 0 {
			stepY = dy / abs(dy) // 确定 y 方向
		}

		x := x1 + stepX
		y := y1 + stepY

		for x != x2 || y != y2 {
			if checkerBoard.Checkerboard[x][y][0] != 0 {
				return false, "移动路径被阻挡"
			}

			x += stepX
			y += stepY
		}

		if checkerBoard.Checkerboard[x2][y2][0] != 0 && checkerBoard.Checkerboard[x2][y2][1] == color {
			return false, "目标位置存在己方棋子"
		}

		tempBoard := copyChessBoard(checkerBoard)
		// 执行临时移动。
		tempBoard.Checkerboard[x1][y1][0], tempBoard.Checkerboard[x2][y2][0] = model.Void, tempBoard.Checkerboard[x1][y1][0]
		tempBoard.Checkerboard[x1][y1][1], tempBoard.Checkerboard[x2][y2][1] = model.Void, tempBoard.Checkerboard[x1][y1][1]
		CalculateThreaten(color, &tempBoard)
		if tempBoard.Checkerboard[tempBoard.King[color][0]][tempBoard.King[color][1]][flag] == 1 {
			return false, "王有威胁"
		}
	}
	return true, ""
}
