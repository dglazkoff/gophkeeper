package storageservice

import (
	"context"
	"fmt"
	"gophkeeper/internal/logger"
)

var ErrDataAlreadyExists = fmt.Errorf("data already exists")
var ErrDataNotFound = fmt.Errorf("data not found")

type storage interface {
	SaveStringData(ctx context.Context, userId, key, stringData, metadata string) error
	GetStringData(ctx context.Context, userId, key string) (string, string, error)
	DeleteStringData(ctx context.Context, userId, key string) error
}

type StorageService struct {
	storage storage
}

func NewStorageService(storage storage) *StorageService {
	return &StorageService{storage}
}

// надо как-то шифровать данные в базе, но пока не понятно как
// потому что будет требоваться чтобы пароли в явном виде не лежали в базе
func (s *StorageService) SavePassword(ctx context.Context, userId, key, login, password, metadata string) error {
	data, _, _ := s.storage.GetStringData(ctx, userId, key)

	if data != "" {
		return ErrDataAlreadyExists
	}

	stringData := fmt.Sprintf("%s %s", login, password)

	return s.storage.SaveStringData(ctx, userId, key, stringData, metadata)
}

func (s *StorageService) GetPassword(ctx context.Context, userId, key string) (login, password, metadata string, err error) {
	data, metadata, err := s.storage.GetStringData(ctx, userId, key)

	if err != nil {
		logger.Log.Error("Error while get password: ", err)
		return "", "", "", ErrDataNotFound
	}

	fmt.Sscanf(data, "%s %s", &login, &password)
	return login, password, metadata, nil
}

func (s *StorageService) DeletePassword(ctx context.Context, userId, key string) error {
	_, _, err := s.storage.GetStringData(ctx, userId, key)

	if err != nil {
		return ErrDataNotFound
	}

	return s.storage.DeleteStringData(ctx, userId, key)
}
