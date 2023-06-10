package model

import "sync"

const (
	King   = 1
	Queen  = 2
	Rook   = 3
	Bishop = 4
	Knight = 5
	Pawn   = 6
)

const (
	White = 1
	Black = 2
)

const Void = 0 //没棋子的格子

type Chess struct {
	Checkerboard [8][8][2]int //x、y轴与格子落子。落子的索引1代表棋子类型，2代表了棋色。参考上面的定义常量。
	Mute         sync.Mutex   //防止并发问题
}
