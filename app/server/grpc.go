package server

import (
	"context"
	"net"

	sample "github.com/kk-no/proto-terminal/sample/v1"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server *grpc.Server
}

var _ Server = (*GRPCServer)(nil)

func NewGRPCServer() (*GRPCServer, error) {
	s := &GRPCServer{
		server: grpc.NewServer(),
	}
	sample.RegisterSampleServiceServer(s.server, NewSampleServer())
	return s, nil
}

func (s *GRPCServer) Serve(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return s.server.Serve(listener)
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
