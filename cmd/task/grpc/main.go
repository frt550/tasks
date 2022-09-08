package main

import (
	"context"
	"net"
	"net/http"
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
	"time"

	"github.com/bradfitz/gomemcache/memcache"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Shopify/sarama"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"

	_ "tasks/internal/pkg/core/tracing"

	cachedRepositoryPkg "tasks/internal/pkg/core/task/repository/cached/memcache"

	"google.golang.org/grpc"

	_ "tasks/internal/pkg/core/task/metric"
)

func main() {
	go startMetricServer()

	repository := postgres.New(poolPkg.GetInstance())
	cacheClient := memcache.New(config.Config.Memcached.Address)
	cachedRepository := cachedRepositoryPkg.New(repository, cacheClient)
	var task = taskPkg.New(cachedRepository)
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

func startMetricServer() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(config.Config.Task.Metric.HttpAddress, nil); err != nil {
		panic(err)
	}
}

func runConsumers(taskService taskPkg.Interface) {
	saramaConfig := sarama.NewConfig()
	// timeout to wait kafka. Otherwise, task container is stopped when docker-compose up
	saramaConfig.Admin.Timeout = 15 * time.Second
	cg, err := sarama.NewConsumerGroup(
		[]string{config.Config.Kafka.Broker0},
		kafka.ConsumerGroupTask,
		saramaConfig,
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
