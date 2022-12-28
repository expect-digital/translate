package main

import (
	"log"
	"net"

	pb "github.com/expect-digital/translate/pkg/server/translate/v1"
	"google.golang.org/grpc"
)

type TranslateServiceServer struct {
	pb.UnimplementedTranslateServiceServer
}

func main() {
	server := grpc.NewServer()
	pb.RegisterTranslateServiceServer(server, &TranslateServiceServer{})

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("error in listening on port :8080", err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatal("unable to start server", err)
	}
}
