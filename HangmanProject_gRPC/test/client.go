package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hangman_grpc/test/msg"
	"os"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := msg.NewGameServiceClient(conn)
	message := msg.Message{
		Body: "Hello from the client!",
	}

	for {
		response, err := c.Drawing(context.Background(), &message)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stdout, "Response from the Server: %s\n", response.Body)
	}
}
