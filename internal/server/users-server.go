package server

import (
	"context"
	"github.com/antonioalfa22/go-grpc-template/internal/services"
	grpc "github.com/antonioalfa22/go-grpc-template/proto"
)

type usersCRUDServer struct {}

func NewUsersCRUDServer() grpc.UsersCRUDServer{
	return &usersCRUDServer{}
}

func (u usersCRUDServer) CreateUser(ctx context.Context, input *grpc.UserInput) (*grpc.UserResponse, error) {
	s := services.GetUserService()
	return s.CreateUser(input)
}

func (u usersCRUDServer) ReadUser(ctx context.Context, id *grpc.ID) (*grpc.User, error) {
	s := services.GetUserService()
	return s.ReadUser(id)
}

func (u usersCRUDServer) ListUsers(ctx context.Context, req *grpc.ListUserReq) (*grpc.ListUserRes, error) {
	s := services.GetUserService()
	return s.ListUsers()
}


func (u usersCRUDServer) UpdateUser(ctx context.Context, input *grpc.UserInput) (*grpc.UserResponse, error) {
	s := services.GetUserService()
	return s.UpdateUser(input)
}

func (u usersCRUDServer) DeleteUser(ctx context.Context, id *grpc.ID) (*grpc.UserResponse, error) {
	s := services.GetUserService()
	return s.DeleteUser(id)
}