package service

import "Chess/model"

func move(msg model.WebsocketMessage, c *model.Client) {
	checkerBoard := c.UserClient.Room.Checkerboard
	x1, y1, x2, y2 := msg.MoveBefore[0], msg.MoveBefore[1], msg.MoveAfter[0], msg.MoveAfter[1]
	checkerBoard.Checkerboard[x1][y1][0], checkerBoard.Checkerboard[x2][y2][0] = model.Void, checkerBoard.Checkerboard[x1][y1][0]
	checkerBoard.Checkerboard[x1][y1][1], checkerBoard.Checkerboard[x2][y2][1] = model.Void, checkerBoard.Checkerboard[x1][y1][1]
	CalculateThreaten(c.UserClient.Color, checkerBoard)
}
