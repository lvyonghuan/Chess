package model

type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	Avatar        string `json:"avatar"`
	Administrator bool   `json:"administrator"`
}
