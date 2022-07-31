package main

import (
	"fmt"
	"log"

	pb "github.com/aleksbgs/projectf/faceit/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	id := createUser(c)
	fmt.Println("user id ", id)
	//readBlog(c, id)
	//readBlog(c, "aNonExistingID")
	//updateBlog(c, id)
	//listBlog(c)
	//deleteBlog(c, id)
}
