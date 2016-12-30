package collector

import (
	"context"
	"log"
	"time"

	rpc "github.com/inhochoi/go-simple-apm/go-simple-apm"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var tick = time.Second * 5

func Start() {
	go run()
}

func run() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewAPMClient(conn)

	for {
		cpu := GetCpuStatus()
		memory := GetMemoryStatus()
		c.SendCpuStatus(context.Background(), &cpu)
		c.SendMemoryStatus(context.Background(), &memory)
		time.Sleep(time.Second * 5)
	}
}
