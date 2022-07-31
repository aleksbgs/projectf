package main

import (
	"context"
	"log"

	pb "github.com/aleksbgs/projectf/pb"
)

func deleteUser(c pb.UserServiceClient, id string) {
	log.Println("---deleteUser was invoked---")
	_, err := c.DeleteUser(context.Background(), &pb.UserId{Id: id})

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("User was deleted")
}
