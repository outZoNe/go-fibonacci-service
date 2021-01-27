package main

import (
	"github.com/joho/godotenv"
	"github.com/outZoNe/go-fibonacci-service/pkg/api/fibonacci"
	fibonacciService "github.com/outZoNe/go-fibonacci-service/pkg/api/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	server := grpc.NewServer()
	srv := &fibonacci.GrpcFibonacci{}
	fibonacciService.RegisterFibonacciServer(server, srv)

	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
