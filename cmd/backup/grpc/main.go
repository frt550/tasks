package main

import (
	"net"
	apiPkg "tasks/internal/api/backup"
	"tasks/internal/config"
	backupPkg "tasks/internal/pkg/core/backup"
	"tasks/internal/pkg/core/backup/repository/postgres"
	pb "tasks/pkg/api/backup"

	"google.golang.org/grpc"
)

func main() {
	var backup = backupPkg.New(postgres.New())
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
