package main

import (
	"log"
	"net"

	"github.com/matheusr42/go-grpc/pb/pb"
	"github.com/matheusr42/go-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
