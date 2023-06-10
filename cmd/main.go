package main

import (
	"Chess/api"
	"Chess/database"
)

func main() {
	database.InitDB()
	api.InitRouter()
}
