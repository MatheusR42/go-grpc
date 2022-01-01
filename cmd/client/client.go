package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/matheusr42/go-grpc/pb/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %s", err.Error())
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	// AddUser(client)
	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {
	name := "Matheus"

	req := &pb.User{
		Id:    "123",
		Name:  &name,
		Email: "matheus@email.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %s", err.Error())
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	name := "Matheus"

	req := &pb.User{
		Id:    "123",
		Name:  &name,
		Email: "matheus@email.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %s", err.Error())
	}
	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not make gRPC request: %s", err.Error())
		}

		fmt.Println("Status:", stream.Status, stream.GetUser())
	}
}
