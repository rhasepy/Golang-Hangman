package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hangman_grpc/api"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type clientHandle struct {
	stream     api.Services_GameServiceClient
	clientName string
}

func WeatherClient() {
	addr := "localhost:8080"
	// Encrypted with TLS
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := api.NewWeatherServiceClient(conn)
	ctx := context.Background()

	resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Cities: ")
	for _, city := range resp.Items {
		fmt.Printf("\t%s: %s", city.GetCityCode(), city.CityName)
	}

	stream, err := client.QueryWeather(ctx, &api.WeatherRequest{
		CityCode: "tr_ank",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("\nWeather in Ankara:\n")
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("\tTemperature: %.2f\n", msg.GetTemperature())
		time.Sleep(time.Second)
	}
	fmt.Println("Server Stopped Sending...")
}

func (ch *clientHandle) clientConfig() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Your Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	ch.clientName = strings.Trim(name, "\r\n")
}

//send message
func (ch *clientHandle) sendMessage() {

	// create a loop
	for {
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf(" Failed to read from console :: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")

		clientMessageBox := &api.FromClient{
			Name: ch.clientName,
			Body: clientMessage,
		}
		err = ch.stream.Send(clientMessageBox)

		if err != nil {
			fmt.Fprintf(os.Stdout, "Error while sending message to server :: %v", err)
		}
	}
}

//receive message
func (ch *clientHandle) receiveMessage() {

	//create a loop
	for {
		mssg, err := ch.stream.Recv()
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error in receiving message from server :: %v", err)
		}
		//print message to console
		fmt.Printf("%s", mssg.Body)
	}
}

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewServicesClient(conn)
	stream, err := client.GameService(context.Background())
	if err != nil {
		panic(err)
	}

	// implement communication with gRPC server
	ch := clientHandle{stream: stream}
	ch.clientConfig()
	go ch.sendMessage()
	go ch.receiveMessage()

	// blocker
	bl := make(chan bool)
	<-bl
}
