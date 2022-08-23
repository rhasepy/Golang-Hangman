package network

import (
	"fmt"
	"net"
	"os"
	"weather/util"
)

const HOST = "localhost"
const PORT = "9001"
const TYPE = "tcp"

type Socket struct {
	Connection net.Conn
	PORT       string
	HOST       string
	TYPE       string
	CID        int
}

func CreateServer(_type string, _host string, _port string) net.Listener {

	listen, err := net.Listen(_type, _host+":"+_port)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}
	return listen
}

func AcceptConnection(server net.Listener) net.Conn {

	conn, err := server.Accept()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}
	return conn
}

func (s *Socket) ReadSock() []byte {

	sockBuffer := make([]byte, 1024)
	_, err := s.Connection.Read(sockBuffer)
	if err != nil {
		fmt.Fprintf(os.Stdout, "[%s] Read File Error: %s\n", util.GetCurrentTime(), err)
	}
	fmt.Fprintf(os.Stdout, "[%s] Client - %d Msg: %s\n", util.GetCurrentTime(), s.CID, sockBuffer)

	return sockBuffer
}

func (s *Socket) WriteSock(response string) {
	s.Connection.Write([]byte(response))
}
