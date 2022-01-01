package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/matheusr42/go-grpc/pb/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
//  AddUserVerbose(*User, UserService_AddUserVerboseServer) error
//  AddUsers(UserService_AddUsersServer) error
//  AddUserStreamBoth(UserService_AddUserStreamBothServer) error
// }

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// inserting to db

	// returning
	return req, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 2)

	stream.Send(&pb.UserResultStream{
		Status: "Saving in DB",
		User:   req,
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Done",
		User:   req,
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatal("Error while receiving stream", err.Error())
		}

		users = append(users, req)

		fmt.Println("Adding user", req.GetId())
	}
}

func (*UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal("Error while receiving stream", err.Error())
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})

		fmt.Println("Added user", req.GetId())

		if err != nil {
			log.Fatal("Error to send stream to the client: ", err.Error())
		}
	}
}
