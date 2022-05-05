package main

import (
	"strateegy/user-service/database"
	"strateegy/user-service/grpc"
	"strateegy/user-service/server"
)

func main() {
	go grpc.StartServer()
	s := server.NewServer()
	database.StartDB()
	s.Run()
}
