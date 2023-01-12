package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lukegriffith/worldservices/internal/render"
	"github.com/lukegriffith/worldservices/internal/server"
	pb "github.com/lukegriffith/worldservices/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	var ctx = context.Background()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8082))
	if err != nil {
		log.Fatalf("Failed to listen on 8082")
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	worldServer := server.NewServer()
	ctx = context.WithValue(ctx, "server", worldServer)

	pb.RegisterWorldServiceServer(grpcServer, worldServer)
	go grpcServer.Serve(lis)
	//terminal.SetupAndRun(ctx)
	render.Render("blah")
}
