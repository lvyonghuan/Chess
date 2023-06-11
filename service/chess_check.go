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
		return checkRookMove(x1, y1, x2, y2, color, checkerBoard)
	case model.Bishop:
		return checkBishopMove(x1, y1, x2, y2, color, checkerBoard)
	case model.Knight:
		return checkKnightMove(x1, y1, x2, y2, color, checkerBoard)
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
		// 执行临时移动
		tempBoard := copyChessBoard(checkerBoard)
		tempBoard.Checkerboard[x1][y1][0], tempBoard.Checkerboard[x2][y2][0] = model.Void, tempBoard.Checkerboard[x1][y1][0]
		tempBoard.Checkerboard[x1][y1][1], tempBoard.Checkerboard[x2][y2][1] = model.Void, tempBoard.Checkerboard[x1][y1][1]
		CalculateThreaten(color, &tempBoard)
		if tempBoard.Checkerboard[tempBoard.King[color][0]][tempBoard.King[color][1]][flag] == 1 {
			return false, "王有威胁"
		}
	}
	return true, ""
}

func checkRookMove(x1, y1, x2, y2, color int, checkerBoard *model.Chess) (bool, string) {
	var flag = returnFlag(color)
	dx := x2 - x1
	dy := y2 - y1

	if dx != 0 && dy != 0 {
		return false, "车只能沿水平或垂直线移动"
	}

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
	// 执行临时移动
	tempBoard := copyChessBoard(checkerBoard)
	tempBoard.Checkerboard[x1][y1][0], tempBoard.Checkerboard[x2][y2][0] = model.Void, tempBoard.Checkerboard[x1][y1][0]
	tempBoard.Checkerboard[x1][y1][1], tempBoard.Checkerboard[x2][y2][1] = model.Void, tempBoard.Checkerboard[x1][y1][1]
	CalculateThreaten(color, &tempBoard)

	if tempBoard.Checkerboard[tempBoard.King[color][0]][tempBoard.King[color][1]][flag] == 1 {
		return false, "王有威胁"
	}

	return true, ""
}

func checkBishopMove(x1, y1, x2, y2, color int, checkerBoard *model.Chess) (bool, string) {
	var flag = returnFlag(color)
	dx := x2 - x1
	dy := y2 - y1

	if abs(dx) != abs(dy) {
		return false, "象只能沿对角线移动"
	}
	stepX := dx / abs(dx) // 确定 x 方向
	stepY := dy / abs(dy) // 确定 y 方向

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
	// 执行临时移动
	tempBoard := copyChessBoard(checkerBoard)
	tempBoard.Checkerboard[x1][y1][0], tempBoard.Checkerboard[x2][y2][0] = model.Void, tempBoard.Checkerboard[x1][y1][0]
	tempBoard.Checkerboard[x1][y1][1], tempBoard.Checkerboard[x2][y2][1] = model.Void, tempBoard.Checkerboard[x1][y1][1]
	CalculateThreaten(color, &tempBoard)

	if tempBoard.Checkerboard[tempBoard.King[color][0]][tempBoard.King[color][1]][flag] == 1 {
		return false, "王有威胁"
	}
	return true, ""
}

func checkKnightMove(x1, y1, x2, y2, color int, checkerBoard *model.Chess) (bool, string) {
	var flag = returnFlag(color)
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)

	if (dx != 2 || dy != 1) && (dx != 1 || dy != 2) {
		return false, "马的移动方式为 'L' 形"
	}

	if checkerBoard.Checkerboard[x2][y2][0] != 0 && checkerBoard.Checkerboard[x2][y2][1] == color {
		return false, "目标位置存在己方棋子"
	}
	// 执行临时移动
	tempBoard := copyChessBoard(checkerBoard)
	tempBoard.Checkerboard[x1][y1][0], tempBoard.Checkerboard[x2][y2][0] = model.Void, tempBoard.Checkerboard[x1][y1][0]
	tempBoard.Checkerboard[x1][y1][1], tempBoard.Checkerboard[x2][y2][1] = model.Void, tempBoard.Checkerboard[x1][y1][1]
	CalculateThreaten(color, &tempBoard)
	if tempBoard.Checkerboard[tempBoard.King[color][0]][tempBoard.King[color][1]][flag] == 1 {
		return false, "王有威胁"
	}

	return true, ""
}

func checkPawnMove(x1, y1, x2, y2, color int, checkerBoard *model.Chess) (bool, string) {
	var flag = returnFlag(color)
	dx := x2 - x1
	dy := y2 - y1
	isFirstMove := (color == model.White && y1 == 1) || (color == model.Black && y1 == 6)

	// Check for one square move
	if dx == 0 && ((dy == -1 && color == model.Black) || (dy == 1 && color == model.White)) {
		if checkerBoard.Checkerboard[x2][y2][0] == 0 {
			tempBoard := copyChessBoard(checkerBoard)
			tempBoard.Checkerboard[x1][y1][0], tempBoard.Checkerboard[x2][y2][0] = model.Void, tempBoard.Checkerboard[x1][y1][0]
			tempBoard.Checkerboard[x1][y1][1], tempBoard.Checkerboard[x2][y2][1] = model.Void, tempBoard.Checkerboard[x1][y1][1]
			CalculateThreaten(color, &tempBoard)
			if tempBoard.Checkerboard[tempBoard.King[color][0]][tempBoard.King[color][1]][flag] == 1 {
				return false, "王有威胁"
			}
			return true, ""
		}
		return false, "不能将己方棋子移动到目标位置"
	}

	// Check for two square move on first move
	if isFirstMove && dx == 0 && ((dy == -2 && color == model.Black) || (dy == 2 && color == model.White)) {
		if checkerBoard.Checkerboard[x2][y2][0] == 0 && checkerBoard.Checkerboard[x2][y2-1][0] == 0 {
			tempBoard := copyChessBoard(checkerBoard)
			tempBoard.Checkerboard[x1][y1][0], tempBoard.Checkerboard[x2][y2][0] = model.Void, tempBoard.Checkerboard[x1][y1][0]
			tempBoard.Checkerboard[x1][y1][1], tempBoard.Checkerboard[x2][y2][1] = model.Void, tempBoard.Checkerboard[x1][y1][1]
			CalculateThreaten(color, &tempBoard)
			if tempBoard.Checkerboard[tempBoard.King[color][0]][tempBoard.King[color][1]][flag] == 1 {
				return false, "王有威胁"
			}
			return true, ""
		}
		return false, "不能将己方棋子移动到目标位置"
	}

	// Check for diagonal captures
	if abs(dx) == 1 && ((dy == -1 && color == model.Black) || (dy == 1 && color == model.White)) {
		if checkerBoard.Checkerboard[x2][y2][0] != 0 && checkerBoard.Checkerboard[x2][y2][1] != color {
			tempBoard := copyChessBoard(checkerBoard)
			tempBoard.Checkerboard[x1][y1][0], tempBoard.Checkerboard[x2][y2][0] = model.Void, tempBoard.Checkerboard[x1][y1][0]
			tempBoard.Checkerboard[x1][y1][1], tempBoard.Checkerboard[x2][y2][1] = model.Void, tempBoard.Checkerboard[x1][y1][1]
			CalculateThreaten(color, &tempBoard)
			if tempBoard.Checkerboard[tempBoard.King[color][0]][tempBoard.King[color][1]][flag] == 1 {
				return false, "王有威胁"
			}
			return true, ""
		}
		return false, "兵需要对角线上有一个棋子才能斜线移动"
	}

	return false, "兵的移动方式不符合规则"

}

// CheckWin 判断胜负。特别朴素，做不到让对面再多走一步。不是说不能做到，我先摆一会。
func CheckWin(c *model.Client) bool {
	var color = c.UserClient.Color
	//反着来的。A move则判断B
	var flag1, flag2 int
	if color == model.White {
		flag1, flag2 = 3, 1
	} else {
		flag1, flag2 = 2, 0
	}
	kingX, kingY := c.UserClient.Room.Checkerboard.King[flag2][0], c.UserClient.Room.Checkerboard.King[flag2][1]
	//只要王现在没被将死，就还能动
	if c.UserClient.Room.Checkerboard.Checkerboard[kingX][kingY][flag1] == 0 {
		return true
	}
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}
	var count = 0 //定义计数变量，==8时（即周围一圈要么被堵上了要么被将死了）+王自己本身被将死了判输
	//TODO:下一步救场
	for _, d := range directions {
		newX, newY := kingX+d[0], kingY+d[1]
		if c.UserClient.Room.Checkerboard.Checkerboard[newX][newY][flag1] == 1 {
			count++
		}
	}
	return count == 8
}
