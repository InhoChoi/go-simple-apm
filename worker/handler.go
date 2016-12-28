package main

import (
	"fmt"

	rpc "github.com/inhochoi/go-simple-apm/go-simple-apm"
	"golang.org/x/net/context"
)

type Handler struct{}

func (h *Handler) SendCpuStatus(ctx context.Context, in *rpc.CpuStatus) (*rpc.Result, error) {
	fmt.Println(in)
	return &rpc.Result{Result: true}, nil
}
