package grpc

import (
	"log"

	grpcgo "google.golang.org/grpc"
)

func GetConn() *grpcgo.ClientConn {
	conn, err := grpcgo.Dial("localhost:3335", grpcgo.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
