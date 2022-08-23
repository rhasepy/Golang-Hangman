package msg

import (
	"context"
	"fmt"
	"os"
)

type Server struct {
}

func (s *Server) Drawing(ctx context.Context, msg *Message) (*Message, error) {
	fmt.Fprintf(os.Stdout, "Received message body from client: %s", msg.Body)
	return &Message{Body: "Response from the Server!"}, nil
}
