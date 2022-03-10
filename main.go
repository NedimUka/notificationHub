package main

import (
	"log"
	"net"
	"os"

	pb "github.com/NedimUka/notification-hub/protoModels"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedNotificationServer
}

func main() {

	// Start listening on port
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port: %v. Error: %v", "50051", err)
	} else {
		log.Printf("Start Listening on port: %v\n", "50051")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServer(grpcServer, &server{})

	// Register reflaction service on gRPC server
	reflection.Register(grpcServer)

	// Starts the gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Printf("There is an error while starting a grpc server: %v\n", err)
		os.Exit(1)
	}
}
