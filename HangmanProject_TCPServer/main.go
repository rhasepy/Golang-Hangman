package main

import (
	"fmt"
	"game/source/Network"
	"game/source/engine"
	"game/source/util"
	"os"
	"sync"
)

var activeConnection int = 0
var mutex sync.Mutex

func getHelloMsg() []byte {
	return []byte("\n******* ****** ****** ****** ****** ****** ******\n" +
		"*** *** *** Welcome the Hangman Game! *** *** ***\n" +
		"*** *** *** * * Game Starting * * *** *** *** ***\n" +
		"******* ****** ****** ****** ****** ****** ******\n\n")
}

func handleRequest(socket Network.Socket) {

	socket.Connection.Write(getHelloMsg())

	engine.GameRoutine(socket)

	mutex.Lock()
	activeConnection -= 1
	socket.Connection.Close()
	mutex.Unlock()
}

func startServerActivity() {

	fmt.Fprintf(os.Stdout, "[%s] Server Creating...\n", util.GetCurrentTime())

	client := Network.Socket{
		HOST: Network.HOST,
		PORT: Network.PORT,
		TYPE: Network.TYPE,
	}

	server := Network.CreateServer(Network.TYPE, Network.HOST, Network.PORT)
	defer server.Close()

	fmt.Fprintf(os.Stdout, "[%s] Server Created...\n", util.GetCurrentTime())
	for {
		client.Connection = Network.AcceptConnection(server)

		mutex.Lock()
		activeConnection += 1
		client.CID = activeConnection
		fmt.Fprintf(os.Stdout, "[%s] Server accept connection. Active Connection: %d\n", util.GetCurrentTime(), activeConnection)
		mutex.Unlock()

		go handleRequest(client)
	}
}

func main() {
	startServerActivity()
}
