package api

import (
	"context"
	pbServer "gophkeeper/internal/proto/server"

	"google.golang.org/grpc/status"
)

type serverService interface {
	Ping(ctx context.Context) error
}

type Server struct {
	pbServer.ServerServer
	service serverService
}

func NewServer(service serverService) *Server {
	return &Server{service: service}
}

func (s *Server) Ping(ctx context.Context, in *pbServer.PingRequest) (*pbServer.PingResponse, error) {
	err := s.service.Ping(ctx)

	if err != nil {
		return nil, status.Errorf(500, "failed to ping: %v", err)
	}

	return &pbServer.PingResponse{}, nil
}
