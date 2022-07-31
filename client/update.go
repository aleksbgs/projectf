package main

import (
	"context"
	"log"

	pb "github.com/aleksbgs/projectf/pb"
)

func updateUser(c pb.UserServiceClient, id string) {
	log.Println("---updateUser was invoked---")

	user := &pb.User{
		Id:        id,
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

	_, err := c.UpdateUser(context.Background(), user)

	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Println("User was updated")
}
