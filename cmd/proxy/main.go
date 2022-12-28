package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux()
	// setting up a dail up for gRPC service by specifying endpoint/target url
	err := pb.RegisterTranslateServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:8080",
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatal(err)
	}
	// Creating a normal HTTP server
	server := http.Server{
		Handler:           mux,
		ReadHeaderTimeout: time.Minute,
	}
	// creating a listener for server
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	// start server
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
