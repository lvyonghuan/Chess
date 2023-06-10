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
			switch checkerBoard.Checkerboard[i][j][0] {
			case model.King:
				calculateKing(i, j, flag, checkerBoard)
			case model.Queen:
				calculateQueen(i, j, flag, checkerBoard)
			case model.Rook:
				calculateRock(i, j, flag, checkerBoard)

			}
		}
	}
}

// 计算各个子的威胁格子
// 国王的
func calculateKing(x, y, flag int, checkerBoard *model.Chess) {
	if x-1 > 0 {
		checkerBoard.Checkerboard[x-1][y][flag] = 1
		if y+1 < 7 {
			checkerBoard.Checkerboard[x][y+1][flag] = 1
			checkerBoard.Checkerboard[x-1][y+1][flag] = 1
		}
		if y-1 > 0 {
			checkerBoard.Checkerboard[x][y-1][flag] = 1
			checkerBoard.Checkerboard[x-1][y-11][flag] = 1
		}
	}
	if x+1 < 7 {
		checkerBoard.Checkerboard[x+1][y][flag] = 1
		if y+1 > 7 {
			checkerBoard.Checkerboard[x+1][y+1][flag] = 1
		}
		if y-1 > 0 {
			checkerBoard.Checkerboard[x][y-1][flag] = 1
		}
	}
}

// 计算王后的
func calculateQueen(x, y, flag int, checkerBoard *model.Chess) {
	//向下遍历
	var (
		flagXl   = true //左
		flagXr   = true //右
		flagDown = true //下
		i        = x - 1
		j        = y - 1
		k        = y + 1
	)
	for ; i >= 0; i-- {
		if flagDown {
			checkerBoard.Checkerboard[i][y][flag] = 1
		}
		if checkerBoard.Checkerboard[i][y][0] != 0 && flagDown {
			flagDown = false
		}
		if j >= 0 {
			if flagXl {
				checkerBoard.Checkerboard[i][j][flag] = 1
			}
			if checkerBoard.Checkerboard[i][j][0] != 0 && flagXl {
				flagXl = false
			} else if flagXr {
				j -= 1
			}
		}
		if k <= 7 {
			if flagXr {
				checkerBoard.Checkerboard[i][k][flag] = 1
			}
			if checkerBoard.Checkerboard[i][k][0] != 0 && flagXr {
				flagXr = false
			} else if flagXr {
				k += 1
			}
		}
	}
	//向上遍历
	flagXr = true
	flagXl = true
	var flagUp = true
	i = x + 1
	j = y - 1
	k = y + 1
	for ; i <= 7; i++ {
		if flagUp {
			checkerBoard.Checkerboard[i][y][flag] = 1
		}
		if checkerBoard.Checkerboard[i][y][0] != 0 && flagUp {
			flagUp = false
		}
		if j >= 0 {
			if flagXl {
				checkerBoard.Checkerboard[i][j][flag] = 1
			}
			if checkerBoard.Checkerboard[i][j][0] != 0 && flagXl {
				flagXl = false
			} else if flagXr {
				j -= 1
			}
		}
		if k <= 7 {
			if flagXr {
				checkerBoard.Checkerboard[i][k][flag] = 1
			}
			if checkerBoard.Checkerboard[i][k][0] != 0 && flagXr {
				flagXr = false
			} else if flagXr {
				k += 1
			}
		}
	}
}

// 车
func calculateRock(x, y, flag int, checkerBoard *model.Chess) {
	//向下
	for i := x; i >= 0; i-- {
		checkerBoard.Checkerboard[i][y][flag] = 1
		if checkerBoard.Checkerboard[i][y][0] != 0 {
			break
		}
	}
	//向上
	for i := x; i <= 7; i++ {
		checkerBoard.Checkerboard[i][y][flag] = 1
		if checkerBoard.Checkerboard[i][y][0] != 0 {
			break
		}
	}
	//向左
	for j := x; j >= 0; j-- {
		checkerBoard.Checkerboard[x][j][flag] = 1
		if checkerBoard.Checkerboard[x][j][0] != 0 {
			break
		}
	}
	//向右
	for j := x; j >= 0; j++ {
		checkerBoard.Checkerboard[x][j][flag] = 1
		if checkerBoard.Checkerboard[x][j][0] != 0 {
			break
		}
	}
}
