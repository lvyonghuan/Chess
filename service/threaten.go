package service

import "Chess/model"

//专为国王设计的威胁标记器

/*
要考虑的：
1.计算所有威胁的辐射范围
2.如果被棋子挡住，则截断
3.棋子移动后，重新计算威胁
4.理论上使用一些技巧可以尽可能地减少计算量。但是我没这个时间去设计一套算法了，所以每次移动之后重新计算全图威胁。
5.什么情况算王将死了？王的四周全是威胁和不能移动的格子。
6.awsl
7.nmd，这分明也是移动检测
*/

func CalculateThreaten(color int, checkerBoard *model.Chess) {
	var flag int
	if color == model.White {
		flag = 2
	} else {
		flag = 3
	}
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if checkerBoard.Checkerboard[i][j][1] == color {
				switch checkerBoard.Checkerboard[i][j][0] {
				case model.King:
					calculateKing(i, j, flag, checkerBoard)
				case model.Queen:
					calculateQueen(i, j, flag, checkerBoard)
				case model.Rook:
					calculateRock(i, j, flag, checkerBoard)
				case model.Bishop:
					calculateBishop(i, j, flag, checkerBoard)
				case model.Knight:
					calculateKnight(i, j, flag, checkerBoard)
				case model.Pawn:
					calculatePawn(i, j, flag, color, checkerBoard)
				default: //空格
					continue
				}
			}
		}
	}
}

// 计算各个子的威胁格子
// 国王的
func calculateKing(x, y, flag int, checkerBoard *model.Chess) {
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}
	for _, d := range directions {
		newX, newY := x+d[0], y+d[1]
		if newX >= 0 && newX <= 7 && newY >= 0 && newY <= 7 {
			checkerBoard.Checkerboard[newX][newY][flag] = 1
		}
	}
}

// 计算后的
func calculateQueen(x, y, flag int, checkerBoard *model.Chess) {
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}
	for _, d := range directions {
		newX, newY := x+d[0], y+d[1]
		for newX >= 0 && newX <= 7 && newY >= 0 && newY <= 7 {
			checkerBoard.Checkerboard[newX][newY][flag] = 1
			if checkerBoard.Checkerboard[newX][newY][0] != 0 {
				break
			}
			newX += d[0]
			newY += d[1]
		}
	}
}

// 车
func calculateRock(x, y, flag int, checkerBoard *model.Chess) {
	directions := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	for _, d := range directions {
		newX, newY := x+d[0], y+d[1]
		for newX >= 0 && newX <= 7 && newY >= 0 && newY <= 7 {
			checkerBoard.Checkerboard[newX][newY][flag] = 1
			if checkerBoard.Checkerboard[newX][newY][0] != 0 {
				break
			}
			newX += d[0]
			newY += d[1]
		}
	}
}

func calculateBishop(x, y, flag int, checkerBoard *model.Chess) {
	directions := [][]int{
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	for _, d := range directions {
		newX, newY := x+d[0], y+d[1]

		for newX >= 0 && newX <= 7 && newY >= 0 && newY <= 7 {
			checkerBoard.Checkerboard[newX][newY][flag] = 1
			if checkerBoard.Checkerboard[newX][newY][0] != 0 {
				break
			}
			newX += d[0]
			newY += d[1]
		}
	}
}

func calculateKnight(x, y, flag int, checkerBoard *model.Chess) {
	offsets := [][]int{
		{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1},
	}
	for _, o := range offsets {
		newX, newY := x+o[0], y+o[1]
		if newX >= 0 && newX <= 7 && newY >= 0 && newY <= 7 {
			checkerBoard.Checkerboard[newX][newY][flag] = 1
		}
	}
}

func calculatePawn(x, y, flag, color int, checkerBoard *model.Chess) {
	//反正white在下面
	if color == model.White {
		if x+1 <= 7 {
			if y-1 >= 0 {
				checkerBoard.Checkerboard[x+1][y-1][flag] = 1
			}
			if y+1 <= 7 {
				checkerBoard.Checkerboard[x+1][y+1][flag] = 1
			}
		}
	} else {
		if x-1 >= 0 {
			if y-1 >= 0 {
				checkerBoard.Checkerboard[x-1][y-1][flag] = 1
			}
			if y+1 <= 7 {
				checkerBoard.Checkerboard[x-1][y+1][flag] = 1
			}
		}
	}
}
