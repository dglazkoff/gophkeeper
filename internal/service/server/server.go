package serverservice

import (
	"context"
)

type storage interface {
	Ping(ctx context.Context) error
}

type ServerService struct {
	storage storage
}

func NewServerService(storage storage) *ServerService {
	return &ServerService{storage: storage}
}

func (s *ServerService) Ping(ctx context.Context) error {
	return s.storage.Ping(ctx)
}
