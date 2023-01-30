package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	demo "lesson-5-goland/proto"
)

func main() {
	otps := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", otps)

	if err != nil {
		log.Error("Error", err)
	}

	defer cc.Close()

	client := demo.NewRestaurantLikeServiceClient(cc)
	request := &demo.RestaurantLikeStarRequest{ReIds: []int32{1, 2, 3}}

	for i := 0; i <= 5; i++ {
		resp, _ := client.GetRestaurantLikeStar(context.Background(), request)

		fmt.Println("Receiver response: ", resp.Result)
	}

}
