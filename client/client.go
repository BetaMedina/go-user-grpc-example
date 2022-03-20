package main

import (
	"context"
	"grpc-user/pkg/user"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := user.NewUserServiceClient(conn)
	response, err := c.SaveUser(context.Background(), &user.CreateUserInput{
		Name:     "dev",
		Email:    "any_dev@test.com",
		Document: "anyDocument",
		Age:      23,
	})
	if err != nil {
		log.Fatalf("error when calling saveUser: %v", err)
	}
	log.Printf("Response from the server is equal to: %s", response)

	userResponse, err := c.ReadUser(context.Background(), &user.ReadUserInput{Id: 1})
	if err != nil {
		log.Fatalf("error when calling ReadUser: %v", err)
	}
	log.Printf("Response from the server is equal to: %s", userResponse)
}
