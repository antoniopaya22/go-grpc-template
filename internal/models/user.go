package models

import grpc "github.com/antonioalfa22/go-grpc-template/proto"

type User struct {
	ID        uint64   `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	Username  string   `gorm:"column:username;not null;unique_index:username" json:"username" form:"username"`
	Hash      string   `gorm:"column:hash;not null;" json:"hash"`
}

func (u User) ToGRPC() *grpc.User {
	return &grpc.User{
		Id: u.ID,
		Username: u.Username,
		Hash: u.Hash,
	}
}