package main

import (
	"context"
	"net"
	"net/http"
	apiPkg "tasks/internal/api/backup"
	"tasks/internal/config"
	backupPkg "tasks/internal/pkg/core/backup"
	"tasks/internal/pkg/core/backup/consumer"
	"tasks/internal/pkg/core/backup/interceptor"
	"tasks/internal/pkg/core/backup/repository/postgres"
	"tasks/internal/pkg/core/logger"
	pb "tasks/pkg/api/backup"
	"tasks/pkg/contract/kafka"
	"time"

	"github.com/Shopify/sarama"

	cachePkg "github.com/patrickmn/go-cache"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"google.golang.org/grpc"

	_ "tasks/internal/pkg/core/tracing"

	producerPkg "tasks/internal/pkg/core/producer"
)

func main() {
	go startMetricServer()
	cache := cachePkg.New(5*time.Minute, 10*time.Minute)

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true
	producer := producerPkg.New([]string{config.Config.Kafka.Broker0}, saramaConfig)

	var backup = backupPkg.New(postgres.New(), cache, producer)
	go runConsumers(cache)

	runGRPCServer(backup)
}

func runGRPCServer(backup backupPkg.Interface) {
	listener, err := net.Listen(config.Config.Backup.Grpc.ServerNetwork, config.Config.Backup.Grpc.ServerAddress)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_opentracing.UnaryServerInterceptor(),
			),
			interceptor.ServerMetricInterceptor(),
		),
	)
	pb.RegisterAdminServer(grpcServer, apiPkg.New(backup))

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

func startMetricServer() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(config.Config.Backup.Metric.HttpAddress, nil); err != nil {
		panic(err)
	}
}

func runConsumers(cache *cachePkg.Cache) {
	saramaConfig := sarama.NewConfig()
	// timeout to wait kafka. Otherwise, backup container is stopped when docker-compose up
	saramaConfig.Admin.Timeout = 15 * time.Second
	cg, err := sarama.NewConsumerGroup(
		[]string{config.Config.Kafka.Broker0},
		kafka.ConsumerGroupBackup,
		saramaConfig,
	)
	if err != nil {
		logger.Logger.Sugar().Fatal(err)
	}
	ctx := context.Background()
	handler := consumer.New(cache)
	for {
		if err := cg.Consume(ctx, []string{kafka.TopicTaskAllResponse}, handler); err != nil {
			logger.Logger.Sugar().Error(err)
		}
	}
}
