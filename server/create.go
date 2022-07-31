package main

import (
	"context"
	"fmt"
	pb "github.com/aleksbgs/projectf/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (*Server) CreateUser(ctx context.Context, in *pb.User) (*pb.UserId, error) {
	log.Printf("Create User was invoked with %v\n", in)

	data := UserItem{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		NickName:  in.Nickname,
		Password:  in.Password,
		Email:     in.Email,
		Country:   in.Country,
		CreatedAt: in.CreatedAt,
		UpdateAt:  in.UpdatedAt,
		DeletedAt: in.DeletedAt,
	}
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.UserId{
		Id: oid.Hex(),
	}, nil
}
