package services

import (
	"errors"
	"fmt"
	"github.com/antonioalfa22/go-grpc-template/internal/models"
	"github.com/antonioalfa22/go-grpc-template/internal/repository"
	"github.com/antonioalfa22/go-grpc-template/pkg/crypto"
	grpc "github.com/antonioalfa22/go-grpc-template/proto"
)

type UserService struct {}
var userService *UserService

func GetUserService() *UserService {
	if userService == nil {
		userService = &UserService{}
	}
	return userService
}

func (u UserService) CreateUser(input *grpc.UserInput) (*grpc.UserResponse, error) {
	r := repository.GetUserRepository()
	user := models.User{
		Username: input.Username,
		Hash: crypto.HashAndSalt([]byte(input.Password)),
	}
	if err := r.Add(&user); err != nil {
		return nil, errors.New("user create error")
	} else {
		return &grpc.UserResponse{
			Id: user.ID,
		}, nil
	}
}

func (u UserService) ReadUser(id *grpc.ID) (*grpc.User, error) {
	r := repository.GetUserRepository()
	if user, err := r.Get(id.Id); err != nil {
		fmt.Println(err)
		return nil, errors.New("record not found")
	} else {
		return user.ToGRPC(), nil
	}
}

func (u UserService) ListUsers() (*grpc.ListUserRes, error) {
	r := repository.GetUserRepository()
	q := models.User{}
	if users, err := r.Query(&q); err != nil {
		fmt.Println(err)
		return nil, errors.New("record not found")
	} else {
		var res []*grpc.User
		for _, u := range *users {
			res = append(res, u.ToGRPC())
		}
		return &grpc.ListUserRes{Users: res}, nil
	}
}

func (u UserService) UpdateUser(input *grpc.UserInput) (*grpc.UserResponse, error) {
	r := repository.GetUserRepository()
	if user, err := r.Get(input.Id); err != nil {
		fmt.Println(err)
		return nil, errors.New("record not found")
	} else {
		user.Username = input.Username
		user.Hash = crypto.HashAndSalt([]byte(input.Password))
		if err := r.Update(user); err != nil {
			return nil, errors.New("user update error")
		} else {
			return &grpc.UserResponse{
				Id: user.ID,
			}, nil
		}
	}
}

func (u UserService) DeleteUser(id *grpc.ID) (*grpc.UserResponse, error) {
	r := repository.GetUserRepository()
	if user, err := r.Get(id.Id); err != nil {
		fmt.Println(err)
		return nil, errors.New("record not found")
	} else {
		if err := r.Delete(user); err != nil {
			return nil, errors.New("user delete error")
		} else {
			return &grpc.UserResponse{
				Id: id.Id,
			}, nil
		}
	}
}
