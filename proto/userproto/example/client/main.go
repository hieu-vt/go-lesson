package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	user "lesson-5-goland/proto/userproto"
)

func main() {
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial("localhost:50051", opt)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	client := user.NewUserServiceClient(cc)

	response, err := client.GetUserByIds(context.Background(), &user.UserRequest{UserIds: []int32{1}})

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	log.Println(response.Users)
}
