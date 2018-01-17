package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc_transcoder/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println(time.Now().Format(time.RFC3339), "Incoming request", in)
	if in.Name == "timeout" {
                // sleep for some time because this is long process
		time.Sleep(90 * time.Second)
		fmt.Println(time.Now().Format(time.RFC3339), "Send response", codes.DeadlineExceeded)
		return nil, status.Error(codes.DeadlineExceeded, in.Name)
	}

	if in.Name == "internal" {
		fmt.Println(time.Now().Format(time.RFC3339), "Send response", codes.Internal)
		return nil, status.Error(codes.Internal, in.Name)
	}

	fmt.Println(time.Now().Format(time.RFC3339), "Send response", "Hello "+in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
  log.Printf("grpc server listens on port %s...\n", port)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
