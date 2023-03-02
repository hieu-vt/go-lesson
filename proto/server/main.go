package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	demo2 "lesson-5-goland/proto/demo"
	"net"
)

type server struct {
	demo2.UnimplementedRestaurantLikeServiceServer
}

func (s *server) GetRestaurantLikeStar(ctx context.Context, req *demo2.RestaurantLikeStarRequest) (*demo2.RestaurantLikeStarResponse, error) {
	log.Println("Server receiver: ", req.ReIds)
	return &demo2.RestaurantLikeStarResponse{Result: map[int32]int32{
		1: 1,
		2: 4,
	}}, nil
}

func main() {
	address := "0.0.0.0:50051"

	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error", err)
		log.Fatalf("Error %v", err)
	}

	fmt.Println("Server listening with address", address)

	s := grpc.NewServer()

	demo2.RegisterRestaurantLikeServiceServer(s, &server{})

	s.Serve(lis)
}
