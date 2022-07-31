package main

import (
	pb "github.com/aleksbgs/projectf/faceit/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	NickName  string             `bson:"nick_name"`
	Password  string             `bson:"password"`
	Email     string             `bson:"email"`
	Country   string             `bson:"country"`
	CreatedAt string             `bson:"created_at"`
	UpdateAt  string             `bson:"update_at"`
	DeletedAt string             `bson:"deleted_at"`
}

func documentToUser(data *UserItem) *pb.User {
	return &pb.User{
		Id:        data.ID.Hex(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Nickname:  data.NickName,
		Password:  data.Password,
		Email:     data.Email,
		Country:   data.Country,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdateAt,
		DeletedAt: data.DeletedAt,
	}
}
