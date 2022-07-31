package main

import (
	"context"
	pb "github.com/aleksbgs/projectf/faceit/proto"
	"log"
)

func createUser(c pb.UserServiceClient) string {
	log.Println("---createUser was invoked---")

	user := &pb.User{
		FirstName: "first faceit",
		LastName:  "last faceit",
		Nickname:  "nickname faceit",
		Password:  "password faceit",
		Email:     "email faceit@gmail.com",
		Country:   "uk faceit",
		CreatedAt: "",
		UpdatedAt: "",
		DeletedAt: "",
	}

	res, err := c.CreateUser(context.Background(), user)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("User has been created: %v\n", res)
	return res.Id
}
