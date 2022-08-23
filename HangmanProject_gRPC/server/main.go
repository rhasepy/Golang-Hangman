package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hangman_grpc/api"
	"hangman_grpc/util"
	"math/rand"
	"net"
	"os"
)

type myWeatherService struct {
	api.UnimplementedWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context, req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {
	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{CityCode: "tr_ank",
				CityName: "Ankara"},
			&api.CityEntry{CityCode: "tr_ist",
				CityName: "Istanbul"},
		},
	}, nil
}

func (m *myWeatherService) QueryWeather(req *api.WeatherRequest, resp api.WeatherService_QueryWeatherServer) error {
	for {
		err := resp.Send(&api.WeatherResponse{Temperature: rand.Float32()*10 + 10})
		if err != nil {
			break
		}
	}
	return nil
}

func WeatherServiceStart() {

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterWeatherServiceServer(srv, &myWeatherService{})
	fmt.Println("Starting server...")
	panic(srv.Serve(lis))
}

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
