package main

import (
	"net"
	apiPkg "tasks/internal/api/backup"
	"tasks/internal/config"
	backupPkg "tasks/internal/pkg/core/backup"
	pb "tasks/pkg/api/backup"

	"google.golang.org/grpc"
)

func main() {
	var backup = backupPkg.New()
	runGRPCServer(backup)
}

func runGRPCServer(backup backupPkg.Interface) {
	listener, err := net.Listen(config.Config.Backup.Grpc.ServerNetwork, config.Config.Backup.Grpc.ServerAddress)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New(backup))

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
