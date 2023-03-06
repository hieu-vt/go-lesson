package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	demo2 "lesson-5-goland/proto/demo"
)

func main() {
	otps := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", otps)

	if err != nil {
		log.Error("Error", err)
	}

	defer cc.Close()

	client := demo2.NewRestaurantLikeServiceClient(cc)
	request := &demo2.RestaurantLikeStarRequest{ReIds: []int32{1, 2, 3}}

	for i := 0; i <= 5; i++ {
		resp, _ := client.GetRestaurantLikeStar(context.Background(), request)

		fmt.Println("Receiver response: ", resp.Result)
	}

}
