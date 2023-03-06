package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	user "lesson-5-goland/proto/userproto"
	"net"
	"net/http"
)

type server struct{}

func (s *server) GetUserByIds(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	return &user.UserResponse{Users: []*user.User{
		{
			Id:        12,
			FirstName: "vu",
			LastName:  "hieu",
			Role:      "user",
		},
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

	user.RegisterUserServiceServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Error %v", err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = user.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":3000",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:3000")
	log.Fatalln(gwServer.ListenAndServe())
}
