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

type s3Storage interface {
	SaveBinaryData(ctx context.Context, userId, key string, data []byte, metadata string) error
	GetBinaryData(ctx context.Context, userId, key string) ([]byte, string, error)
	DeleteBinaryData(ctx context.Context, userId, key string) error
}

type StorageService struct {
	storage   storage
	s3Storage s3Storage
}

func NewStorageService(storage storage, s3Storage s3Storage) *StorageService {
	return &StorageService{storage: storage, s3Storage: s3Storage}
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

func (s *StorageService) SaveText(ctx context.Context, userId, key, text, metadata string) error {
	data, _, _ := s.storage.GetStringData(ctx, userId, key)

	if data != "" {
		return ErrDataAlreadyExists
	}

	stringData := fmt.Sprintf("%s", text)

	return s.storage.SaveStringData(ctx, userId, key, stringData, metadata)
}

func (s *StorageService) GetText(ctx context.Context, userId, key string) (text, metadata string, err error) {
	text, metadata, err = s.storage.GetStringData(ctx, userId, key)

	if err != nil {
		logger.Log.Error("Error while get password: ", err)
		return "", "", ErrDataNotFound
	}

	return text, metadata, nil
}

func (s *StorageService) DeleteText(ctx context.Context, userId, key string) error {
	_, _, err := s.storage.GetStringData(ctx, userId, key)

	if err != nil {
		return ErrDataNotFound
	}

	return s.storage.DeleteStringData(ctx, userId, key)
}

func (s *StorageService) SaveBankCard(ctx context.Context, userId, key, num, holder, cvv, expirationDate, metadata string) error {
	data, _, _ := s.storage.GetStringData(ctx, userId, key)

	if data != "" {
		return ErrDataAlreadyExists
	}

	stringData := fmt.Sprintf("%s %s %s %s", num, holder, cvv, expirationDate)

	return s.storage.SaveStringData(ctx, userId, key, stringData, metadata)
}

func (s *StorageService) GetBankCard(ctx context.Context, userId, key string) (num, holder, cvv, expirationDate, metadata string, err error) {
	data, metadata, err := s.storage.GetStringData(ctx, userId, key)

	if err != nil {
		logger.Log.Error("Error while get password: ", err)
		return "", "", "", "", "", ErrDataNotFound
	}

	_, err = fmt.Sscanf(data, "%s %s %s %s", &num, &holder, &cvv, &expirationDate)

	if err != nil {
		logger.Log.Error("Error while get password: ", err)
		return "", "", "", "", "", err
	}

	return num, holder, cvv, expirationDate, metadata, nil
}

func (s *StorageService) DeleteBankCard(ctx context.Context, userId, key string) error {
	_, _, err := s.storage.GetStringData(ctx, userId, key)

	if err != nil {
		return ErrDataNotFound
	}

	return s.storage.DeleteStringData(ctx, userId, key)
}

func (s *StorageService) SaveBinary(ctx context.Context, userId, key string, data []byte, metadata string) error {
	return s.s3Storage.SaveBinaryData(ctx, userId, key, data, metadata)
}

func (s *StorageService) GetBinary(ctx context.Context, userId, key string) ([]byte, string, error) {
	return s.s3Storage.GetBinaryData(ctx, userId, key)
}

func (s *StorageService) DeleteBinary(ctx context.Context, userId, key string) error {
	return s.s3Storage.DeleteBinaryData(ctx, userId, key)
}
