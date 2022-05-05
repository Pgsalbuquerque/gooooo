package grpc

import (
	"context"
	"log"
	"net"
	repository "strateegy/user-service/repositories/mongo"

	grpc "google.golang.org/grpc"
)

type Server struct {
	UnimplementedChangePlanServer
}

func (service *Server) RequestBilling(ctx context.Context, req *Billing) (*Status, error) {
	ID := req.GetID()
	Plan := req.GetPlan()

	repo := &repository.UserRepository{}
	err := repo.ChangePlan(ID, Plan)
	if err != nil {
		return &Status{
			Status: false,
		}, err
	}

	return &Status{
		Status: true,
	}, nil
}
func (service *Server) mustEmbedUnimplementedChangePlanServer() {}

func StartServer() {
	grpcServer := grpc.NewServer()

	RegisterChangePlanServer(grpcServer, &Server{})

	port := ":3339"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Grpc Server running at port: %v", port)

	grpc_Error := grpcServer.Serve(listener)
	if grpc_Error != nil {
		log.Fatal(grpc_Error)
	}
}
