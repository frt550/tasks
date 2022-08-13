package main

import (
	"context"
	"log"
	"tasks/internal/config"
	pb "tasks/pkg/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conns, err := grpc.Dial(config.Config.Grpc.ClientTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewAdminClient(conns)

	ctx := context.Background()
	response, err := client.TaskAll(ctx, &pb.TaskAllRequest{})
	if err != nil {
		panic(err)
	}

	log.Printf("response: [%v]", response)
}
