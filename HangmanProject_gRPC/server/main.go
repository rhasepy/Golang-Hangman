package main

import (
	"fmt"
	"hangman_grpc/api"
	"hangman_grpc/util"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "[%s] Server Listening @: %s\n", util.GetCurrentTime(), "8080")

	grpc_server := grpc.NewServer()

	gs := api.GameServer{}

	api.RegisterServicesServer(grpc_server, &gs)

	err = grpc_server.Serve(lis)
	if err != nil {
		fmt.Fprintf(os.Stdout, "[%s] Failed to start gRPC Server..", util.GetCurrentTime())
	}
}
