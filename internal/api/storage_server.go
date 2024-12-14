package api

import (
	"context"
	"errors"
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

	SaveText(ctx context.Context, userId, key, text, metadata string) error
	GetText(ctx context.Context, userId, key string) (text, metadata string, err error)
	DeleteText(ctx context.Context, userId, key string) error

	SaveBinary(ctx context.Context, userId, key string, value []byte, metadata string) error
	GetBinary(ctx context.Context, userId, key string) (value []byte, metadata string, err error)
	DeleteBinary(ctx context.Context, userId, key string) error

	SaveBankCard(ctx context.Context, userId, key, num, holder, cvv, expirationDate, metadata string) error
	GetBankCard(ctx context.Context, userId, key string) (num, holder, cvv, expirationDate, metadata string, err error)
	DeleteBankCard(ctx context.Context, userId, key string) error
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

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	err = s.service.SavePassword(ctx, userID, in.Key, in.Login, in.Password, *in.Metadata)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "data with key %s already exists", in.Key)
		}

		logger.Log.Error("Error while save password: ", err)
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

func (s *StorageServer) SaveText(ctx context.Context, in *pbStorage.SaveTextRequest) (*pbStorage.SaveTextResponse, error) {
	if in.Text == "" || in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "text and key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	err = s.service.SaveText(ctx, userID, in.Key, in.Text, *in.Metadata)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "data with key %s already exists", in.Key)
		}

		logger.Log.Error("Error while save text: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.SaveTextResponse{}, nil
}

func (s *StorageServer) GetText(ctx context.Context, in *pbStorage.GetTextRequest) (*pbStorage.GetTextResponse, error) {
	if in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	text, metadata, err := s.service.GetText(ctx, userID, in.Key)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataNotFound) {
			return nil, status.Errorf(codes.NotFound, "data not found")
		}

		logger.Log.Error("Error while get text: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.GetTextResponse{Text: text, Metadata: &metadata}, nil
}

func (s *StorageServer) DeleteText(ctx context.Context, in *pbStorage.DeleteTextRequest) (*pbStorage.DeleteTextResponse, error) {
	if in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	err = s.service.DeleteText(ctx, userID, in.Key)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataNotFound) {
			return nil, status.Errorf(codes.NotFound, "data not found")
		}

		logger.Log.Error("Error while delete text: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.DeleteTextResponse{}, nil
}

func (s *StorageServer) SaveBinary(ctx context.Context, in *pbStorage.SaveBinaryRequest) (*pbStorage.SaveBinaryResponse, error) {
	if len(in.Value) == 0 || in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "value and key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	err = s.service.SaveBinary(ctx, userID, in.Key, in.Value, *in.Metadata)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "data with key %s already exists", in.Key)
		}

		logger.Log.Error("Error while save text: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.SaveBinaryResponse{}, nil
}

func (s *StorageServer) GetBinary(ctx context.Context, in *pbStorage.GetBinaryRequest) (*pbStorage.GetBinaryResponse, error) {
	if in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	value, metadata, err := s.service.GetBinary(ctx, userID, in.Key)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataNotFound) {
			return nil, status.Errorf(codes.NotFound, "data not found")
		}

		logger.Log.Error("Error while get binary data: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.GetBinaryResponse{Value: value, Metadata: &metadata}, nil
}

func (s *StorageServer) DeleteBinary(ctx context.Context, in *pbStorage.DeleteBinaryRequest) (*pbStorage.DeleteBinaryResponse, error) {
	if in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	err = s.service.DeleteBinary(ctx, userID, in.Key)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataNotFound) {
			return nil, status.Errorf(codes.NotFound, "data not found")
		}

		logger.Log.Error("Error while delete binary data: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.DeleteBinaryResponse{}, nil
}

func (s *StorageServer) SaveBankCard(ctx context.Context, in *pbStorage.SaveBankCardRequest) (*pbStorage.SaveBankCardResponse, error) {
	if in.Holder == "" || in.Number == "" || in.Cvv == "" || in.ExpirationDate == "" || in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "holeder, number, cvv, expiration date and key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	err = s.service.SaveBankCard(ctx, userID, in.Key, in.Number, in.Holder, in.Cvv, in.ExpirationDate, *in.Metadata)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "data with key %s already exists", in.Key)
		}

		logger.Log.Error("Error while save bank card: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.SaveBankCardResponse{}, nil
}

func (s *StorageServer) GetBankCard(ctx context.Context, in *pbStorage.GetBankCardRequest) (*pbStorage.GetBankCardResponse, error) {
	if in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	num, holder, cvv, expirationDate, metadata, err := s.service.GetBankCard(ctx, userID, in.Key)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataNotFound) {
			return nil, status.Errorf(codes.NotFound, "data not found")
		}

		logger.Log.Error("Error while get bank card: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.GetBankCardResponse{Number: num, Holder: holder, Cvv: cvv, ExpirationDate: expirationDate, Metadata: &metadata}, nil
}

func (s *StorageServer) DeleteBankCard(ctx context.Context, in *pbStorage.DeleteBankCardRequest) (*pbStorage.DeleteBankCardResponse, error) {
	if in.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "key must be not empty")
	}

	userID, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	err = s.service.DeleteBankCard(ctx, userID, in.Key)

	if err != nil {
		if errors.Is(err, storageservice.ErrDataNotFound) {
			return nil, status.Errorf(codes.NotFound, "data not found")
		}

		logger.Log.Error("Error while delete bank card: ", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pbStorage.DeleteBankCardResponse{}, nil
}
