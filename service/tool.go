package service

import (
	"Chess/model"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func copyChessBoard(chessBoard *model.Chess) model.Chess {
	newBoard := model.Chess{
		Checkerboard: chessBoard.Checkerboard,
		King:         chessBoard.King,
	}
	return newBoard
}

func returnFlag(color int) int {
	if color == model.White {
		return 2
	} else {
		return 3
	}
}
