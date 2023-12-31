package model

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
	Checkerboard [8][8][5]int `json:"checkerboard"` //x、y轴与附加信息——格子落子/威胁/王车易位前提判断。附加信息：索引0代表棋子类型，1代表了棋色。参考上面的定义常量。索引2为白方威胁，3为黑方威胁——为国王设计的。威胁参考常量。索引4判断王车是否移动过。1表示已经移动，只针对王和车。
	King         [2][2]int    //国王A/B位置
}

// Upgrade 升变记录
type Upgrade struct {
	IsUpgrade         bool
	UpgradeCoordinate [2]int
}
