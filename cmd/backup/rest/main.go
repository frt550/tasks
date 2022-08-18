package main

import (
	"context"
	"log"
	"net/http"
	"tasks/internal/config"
	pb "tasks/pkg/api/backup"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	runREST()
}

func runREST() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	if err := mux.HandlePath("GET", "/v1/swagger", serveSwaggerFile); err != nil {
		panic(err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, mux, config.Config.Backup.Grpc.ServerAddress, opts); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(config.Config.Backup.Rest.ServerAddress, mux); err != nil {
		panic(err)
	}
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	log.Println("Serving swagger-file: pkg/api/api.swagger.json")
	http.ServeFile(w, r, "pkg/api/api.swagger.json")
}
