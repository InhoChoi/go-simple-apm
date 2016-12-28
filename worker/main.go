package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	rpc "github.com/inhochoi/go-simple-apm/go-simple-apm"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpc.RegisterAPMServer(s, &Handler{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
