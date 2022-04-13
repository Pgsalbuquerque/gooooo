package main

import (
	"strateegy/user-service/database"
	"strateegy/user-service/server"
)

func main() {
	s := server.NewServer()
	database.StartDB()
	s.Run()
}
