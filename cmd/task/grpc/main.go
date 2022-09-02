package main

import (
	"context"
	"net"
	apiPkg "tasks/internal/api/task"
	"tasks/internal/config"
	"tasks/internal/pkg/core/logger"
	poolPkg "tasks/internal/pkg/core/pool"
	producerPkg "tasks/internal/pkg/core/producer"
	taskPkg "tasks/internal/pkg/core/task"
	"tasks/internal/pkg/core/task/consumer"
	"tasks/internal/pkg/core/task/repository/postgres"
	pb "tasks/pkg/api/task"
	"tasks/pkg/contract/kafka"

	"github.com/Shopify/sarama"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"

	_ "tasks/internal/pkg/core/tracing"

	"google.golang.org/grpc"
)

func main() {
	var task = taskPkg.New(postgres.New(poolPkg.GetInstance()))
	go runConsumers(task)
	runGRPCServer(task)
}

func runGRPCServer(task taskPkg.Interface) {
	listener, err := net.Listen(config.Config.Task.Grpc.ServerNetwork, config.Config.Task.Grpc.ServerAddress)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterAdminServer(grpcServer, apiPkg.New(task))

	if err = grpcServer.Serve(listener); err != nil {
		logger.Logger.Sugar().Fatal(err)
	}
}

func runConsumers(taskService taskPkg.Interface) {
	cg, err := sarama.NewConsumerGroup(
		[]string{config.Config.Kafka.Broker0},
		kafka.ConsumerGroupTask,
		sarama.NewConfig(),
	)
	if err != nil {
		logger.Logger.Sugar().Fatal(err)
	}
	ctx := context.Background()

	producerConfig := sarama.NewConfig()
	producerConfig.Producer.Return.Successes = true
	producer := producerPkg.New([]string{config.Config.Kafka.Broker0}, producerConfig)

	handler := consumer.New(producer, taskService)
	for {
		if err := cg.Consume(ctx, []string{kafka.TopicTaskAllRequest}, handler); err != nil {
			logger.Logger.Sugar().Error(err)
		}
	}
}
