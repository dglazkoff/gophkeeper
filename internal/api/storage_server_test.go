package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gophkeeper/internal/auth"
	pbStorage "gophkeeper/internal/proto/storage"
	storageservice "gophkeeper/internal/service/storage"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) SaveStringData(ctx context.Context, userId, key, stringData, metadata string) error {
	args := m.Called(ctx, userId, key, stringData, metadata)
	return args.Error(0)
}

func (m *MockStorage) GetStringData(ctx context.Context, userId, key string) (string, string, error) {
	args := m.Called(ctx, userId, key)
	return args.String(0), args.String(1), args.Error(2)
}

func (m *MockStorage) DeleteStringData(ctx context.Context, userId, key string) error {
	args := m.Called(ctx, userId, key)
	return args.Error(0)
}

type MockS3Storage struct {
	mock.Mock
}

func (m *MockS3Storage) SaveBinaryData(ctx context.Context, userId, key string, data []byte, metadata string) error {
	args := m.Called(ctx, userId, key, data, metadata)
	return args.Error(0)
}

func (m *MockS3Storage) GetBinaryData(ctx context.Context, userId, key string) ([]byte, string, error) {
	args := m.Called(ctx, userId, key)
	return args.Get(0).([]byte), args.String(1), args.Error(2)
}

func (m *MockS3Storage) DeleteBinaryData(ctx context.Context, userId, key string) error {
	args := m.Called(ctx, userId, key)
	return args.Error(0)
}

func setupStorageTest(t *testing.T) (*MockStorage, *MockS3Storage, *StorageServer) {
	mockStorage := new(MockStorage)
	mockS3Storage := new(MockS3Storage)
	storageService := storageservice.NewStorageService(mockStorage, mockS3Storage)
	server := NewStorageServer(storageService)
	return mockStorage, mockS3Storage, server
}

func TestStorageServer_SavePassword(t *testing.T) {
	mockStorage, _, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), auth.Claims{}, auth.Claims{UserID: "test-user"})
	req := &pbStorage.SavePasswordRequest{
		Key:      "test-key",
		Login:    "test-login",
		Password: "test-password",
		Metadata: stringPtr("test-metadata"),
	}

	mockStorage.On("SaveStringData", ctx, "test-user", "test-key", "test-password", "test-metadata").Return(nil)

	err := server.SavePassword(ctx, req)
	assert.NoError(t, err)

	mockStorage.AssertExpectations(t)
}
