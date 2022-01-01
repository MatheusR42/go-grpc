package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

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
	// AddUserVerbose(client)
	AddUsers(client)
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
		log.Fatal("Could not make gRPC request", err.Error())
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
			log.Fatal("Could not make gRPC request", err.Error())
		}

		fmt.Println("Status:", stream.Status, stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{}
	name := "Matheus "
	email := "test@test.com"

	for i := 1; i <= 5; i++ {
		index := strconv.Itoa(i)
		currentName := name + index
		currentEmail := email + index

		reqs = append(reqs, &pb.User{
			Id:    index,
			Name:  &currentName,
			Email: currentEmail,
		})
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatal("Error to create request", err.Error())
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error to receive message", err.Error())
	}

	fmt.Println(res)
}
