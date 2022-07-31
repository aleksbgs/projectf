package main

import (
	"context"
	_ "github.com/aleksbgs/projectf/doc/statik"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"

	pb "github.com/aleksbgs/projectf/pb"
	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var addr string = "0.0.0.0:50051"
var collection *mongo.Collection

type ServerPb struct {
	pb.UserServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("faceit").Collection("users")

	go runGatewayServer()
	runGrpcServer()

}
func runGrpcServer() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening grpc at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &ServerPb{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
func runGatewayServer() {

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: false,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterUserServiceHandlerServer(ctx, grpcMux, &ServerPb{})
	if err != nil {
		log.Fatal("cannot register handler server:", err)
	}
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("cannot create statik fs:", err.Error())
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start HTTP gateway server:", err)
	}

}
