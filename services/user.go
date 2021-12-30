package services

import (
	"context"

	"github.com/matheusr42/go-grpc/pb/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
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
	return &pb.User{
		Id:    "123",
		Name:  nil,
		Email: req.GetEmail(),
	}, nil
}
