package main

import (
	"context"
	"log"

	pb "github.com/aleksbgs/projectf/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*ServerPb) UpdateUser(ctx context.Context, in *pb.User) (*emptypb.Empty, error) {
	log.Printf("UpdateUser was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

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
	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find user with ID",
		)
	}

	return &emptypb.Empty{}, nil
}
