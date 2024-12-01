package api

import (
	"context"
	"errors"
	"fmt"
	"gophkeeper/internal/auth"
	"gophkeeper/internal/logger"
	pbStorage "gophkeeper/internal/proto/storage"
	storageservice "gophkeeper/internal/service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type storageService interface {
	SavePassword(ctx context.Context, userId, key, login, password, metadata string) error
	GetPassword(ctx context.Context, userId, key string) (login, password, metadata string, err error)
	DeletePassword(ctx context.Context, userId, key string) error
}

type StorageServer struct {
	pbStorage.UnimplementedStorageServer
	service storageService
}

func NewStorageServer(service storageService) *StorageServer {
	return &StorageServer{service: service}
}

func (s *StorageServer) SavePassword(ctx context.Context, in *pbStorage.SavePasswordRequest) (*pbStorage.SavePasswordResponse, error) {
	if in.Login == "" || in.Password == "" || in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "login, password and key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	fmt.Println(userID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	err = s.service.SavePassword(ctx, userID, in.Key, in.Login, in.Password, *in.Metadata)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "user with login %s already exists")
		}

		logger.Log.Error("Error while register user: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.SavePasswordResponse{}, nil
}

func (s *StorageServer) GetPassword(ctx context.Context, in *pbStorage.GetPasswordRequest) (*pbStorage.GetPasswordResponse, error) {
	if in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	login, password, metadata, err := s.service.GetPassword(ctx, userID, in.Key)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataNotFound) {
			return nil, status.Errorf(codes.NotFound, "data not found")
		}

		logger.Log.Error("Error while get password: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.GetPasswordResponse{Login: login, Password: password, Metadata: &metadata}, nil
}

func (s *StorageServer) DeletePassword(ctx context.Context, in *pbStorage.DeletePasswordRequest) (*pbStorage.DeletePasswordResponse, error) {
	if in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	err = s.service.DeletePassword(ctx, userID, in.Key)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataNotFound) {
			return nil, status.Errorf(codes.NotFound, "data not found")
		}

		logger.Log.Error("Error while delete password: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.DeletePasswordResponse{}, nil
}
