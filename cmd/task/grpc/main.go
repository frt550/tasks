package main

import (
	"net"
	apiPkg "tasks/internal/api/task"
	"tasks/internal/config"
	taskPkg "tasks/internal/pkg/core/task"
	pb "tasks/pkg/api/task"

	"google.golang.org/grpc"
)

func main() {
	var task = taskPkg.New()
	runGRPCServer(task)
}

func runGRPCServer(task taskPkg.Interface) {
	listener, err := net.Listen(config.Config.Task.Grpc.ServerNetwork, config.Config.Task.Grpc.ServerAddress)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New(task))

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
