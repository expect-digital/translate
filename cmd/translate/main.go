package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
	"github.com/expect-digital/translate/pkg/translate"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TranslateServiceServer struct {
	pb.UnimplementedTranslateServiceServer
}

const serverAddr = "localhost:8080"

func main() {
	// create new gRPC server
	grpcSever := grpc.NewServer()
	pb.RegisterTranslateServiceServer(grpcSever, translate.New())
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux()
	// setting up a dail up for gRPC service by specifying endpoint/target url
	err := pb.RegisterTranslateServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		serverAddr,
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
	l, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(l)
	// a different listener for HTTP1
	httpL := m.Match(cmux.HTTP1Fast())
	// a different listener for HTTP2 since gRPC uses HTTP2
	grpcL := m.Match(cmux.HTTP2())

	go func() {
		if err := server.Serve(httpL); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := grpcSever.Serve(grpcL); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m.Serve(); err != nil {
		log.Fatal(err)
	}
}
