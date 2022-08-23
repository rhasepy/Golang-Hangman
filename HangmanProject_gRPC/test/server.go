package main

import (
	"fmt"
	"google.golang.org/grpc"
	"hangman_grpc/test/msg"
	"hangman_grpc/util"
	"net"
	"os"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	s := msg.Server{}

	grpcServer := grpc.NewServer()
	fmt.Fprintf(os.Stdout, "[%s] Server Created...\n", util.GetCurrentTime())

	msg.RegisterGameServiceServer(grpcServer, &s)

	fmt.Fprintf(os.Stdout, "[%s] Server Starting...\n", util.GetCurrentTime())
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
