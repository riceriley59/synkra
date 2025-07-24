package engine

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pingpongpb "github.com/riceriley59/synkra/proto"
)

type server struct {
	pingpongpb.UnimplementedPingPongServer
}

func (s *server) Ping(ctx context.Context, req *pingpongpb.PingRequest) (*pingpongpb.PongReply, error) {
	fmt.Printf("Received ping: %s\n", req.GetMessage())
	return &pingpongpb.PongReply{Message: "pong: " + req.GetMessage()}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pingpongpb.RegisterPingPongServer(grpcServer, &server{})
	fmt.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

