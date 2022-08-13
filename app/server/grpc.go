package server

import (
	"context"
	"net"

	sample "github.com/kk-no/proto-terminal/sample/v1"
	"google.golang.org/grpc"
)

type GRPCServer struct{}

var _ Server = (*GRPCServer)(nil)

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (s *GRPCServer) Serve(port string) error {
	server := grpc.NewServer()
	sample.RegisterSampleServiceServer(server, NewSampleServer())

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return server.Serve(listener)
}

type SampleServer struct {
	sample.UnimplementedSampleServiceServer
}

func NewSampleServer() *SampleServer {
	return &SampleServer{}
}

func (s *SampleServer) Ping(_ context.Context, _ *sample.PingRequest) (*sample.PingResponse, error) {
	return &sample.PingResponse{}, nil
}
