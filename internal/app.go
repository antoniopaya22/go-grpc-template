package internal

import (
	"fmt"
	"github.com/antonioalfa22/go-grpc-template/config"
	"github.com/antonioalfa22/go-grpc-template/internal/server"
	pb "github.com/antonioalfa22/go-grpc-template/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	config.SetupDB()
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "./cmd/server/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()

	fmt.Println("Go gRPC Running on port " + conf.Server.Port)
	fmt.Println("==================>")

	// Host grpc service
	listen, err := net.Listen("tcp", ":"+conf.Server.Port)
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterUsersCRUDServer(s, server.NewUsersCRUDServer())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
