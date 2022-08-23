package main

import (
	"fmt"
	"game/source/engine"
	"game/source/network"
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

func handleRequest(socket network.Socket) {

	socket.Connection.Write(getHelloMsg())

	engine.GameRoutine(socket)

	mutex.Lock()
	activeConnection -= 1
	socket.Connection.Close()
	mutex.Unlock()
}

func startServerActivity() {

	fmt.Fprintf(os.Stdout, "[%s] Server Creating...\n", util.GetCurrentTime())

	client := network.Socket{
		HOST: network.HOST,
		PORT: network.PORT,
		TYPE: network.TYPE,
	}

	server := network.CreateServer(network.TYPE, network.HOST, network.PORT)
	defer server.Close()

	fmt.Fprintf(os.Stdout, "[%s] Server Created...\n", util.GetCurrentTime())
	for {
		client.Connection = network.AcceptConnection(server)

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
