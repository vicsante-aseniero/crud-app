package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "colinahealth_emr-api/cmd/app/rpc/pb"
	"google.golang.org/grpc"
)

// HelloService ...
type server struct {
	pb.UnimplementedMyServiceServer
}

var addr_gRPC = flag.String("addr_gRPC", ":40051", "the address to connect to")

// MyMethod this will be implementation method hello world
func (s *server) MyMethod(ctx context.Context, req *pb.MyRequest) (*pb.MyResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.MyResponse{Message: "Hello, " + req.GetName()}, nil
}

func main() {

	//serverAddress := "40051"
	//srv, err := net.Listen("tcp", fmt.Sprintf(`:%s`, serverAddress))

	srv, err := net.Listen("tcp", *addr_gRPC)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMyServiceServer(grpcServer, new(server))

	fmt.Printf(`Server running in address : %s`, *addr_gRPC)

	if err = grpcServer.Serve(srv); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}