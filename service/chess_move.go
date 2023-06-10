package service

import "Chess/model"

func move(msg model.WebsocketMessage, c *model.Client) {
	checkerBoard := c.UserClient.Room.Checkerboard
	x1, y1, x2, y2 := msg.MoveBefore[0], msg.MoveBefore[1], msg.MoveAfter[0], msg.MoveAfter[1]
	checkerBoard.Checkerboard[x1][y1][0], checkerBoard.Checkerboard[x2][y2][0] = model.Void, checkerBoard.Checkerboard[x1][y1][0]
	checkerBoard.Checkerboard[x1][y1][1], checkerBoard.Checkerboard[x2][y2][1] = model.Void, checkerBoard.Checkerboard[x1][y1][1]
	//升变检测
	c.UserClient.Room.Upgrade.IsUpgrade = false
	if checkerBoard.Checkerboard[x2][y2][0] == model.Pawn {
		c.UserClient.Room.Upgrade.IsUpgrade = true
		c.UserClient.Room.Upgrade.UpgradeCoordinate[0] = x2
		c.UserClient.Room.Upgrade.UpgradeCoordinate[1] = y2
	}
	CalculateThreaten(model.White, checkerBoard)
	CalculateThreaten(model.Black, checkerBoard)
}

func upgrade(msg model.WebsocketMessage, c *model.Client) {
	x, y := c.UserClient.Room.Upgrade.UpgradeCoordinate[0], c.UserClient.Room.Upgrade.UpgradeCoordinate[1]
	if msg.Upgrade == model.Queen {
		c.UserClient.Room.Checkerboard.Checkerboard[x][y][0] = model.Queen
	} else if msg.Upgrade == model.Rook {
		c.UserClient.Room.Checkerboard.Checkerboard[x][y][0] = model.Rook
	} else if msg.Upgrade == model.Bishop {
		c.UserClient.Room.Checkerboard.Checkerboard[x][y][0] = model.Bishop
	} else if msg.Upgrade == model.Knight {
		c.UserClient.Room.Checkerboard.Checkerboard[x][y][0] = model.Knight
	}
	CalculateThreaten(c.UserClient.Color, c.UserClient.Room.Checkerboard)
}
