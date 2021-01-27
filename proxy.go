package main

import (
	"context"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	gw "github.com/outZoNe/go-fibonacci-service/pkg/api/pb" // Update
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	grpcServerEndpoint := flag.String("grpc-server-endpoint", os.Getenv("GRPC_ADDR")+":"+os.Getenv("GRPC_PORT"), "gRPC server endpoint")
	err := gw.RegisterFibonacciHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":"+os.Getenv("HTTP_REST_PORT"), mux)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
