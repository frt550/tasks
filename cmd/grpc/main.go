package main

import (
	"net"
	apiPkg "tasks/internal/api"
	"tasks/internal/config"
	taskPkg "tasks/internal/pkg/core/task"
	pb "tasks/pkg/api"

	"google.golang.org/grpc"
)

func main() {
	var task = taskPkg.New()
	runGRPCServer(task)
}

func runGRPCServer(task taskPkg.Interface) {
	listener, err := net.Listen(config.Config.Grpc.ServerNetwork, config.Config.Grpc.ServerAddress)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New(task))

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
