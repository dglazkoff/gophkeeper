package api

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

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

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.SavePasswordRequest{
		Key:      "test-key",
		Login:    "test-login",
		Password: "test-password",
		Metadata: stringPtr("test-metadata"),
	}

	mockStorage.On("GetStringData", ctx, "test-user", "test-key").
		Return("", "", errors.New("not found"))
	mockStorage.On("SaveStringData", ctx, "test-user", "test-key", "test-login test-password", "test-metadata").
		Return(nil)

	resp, err := server.SavePassword(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockStorage.AssertExpectations(t)
}

func TestStorageServer_GetPassword(t *testing.T) {
	mockStorage, _, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.GetPasswordRequest{
		Key: "test-key",
	}

	mockStorage.On("GetStringData", ctx, "test-user", "test-key").
		Return("test-login test-password", "test-metadata", nil)

	resp, err := server.GetPassword(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "test-login", resp.Login)
	assert.Equal(t, "test-password", resp.Password)
	assert.Equal(t, "test-metadata", *resp.Metadata)
	mockStorage.AssertExpectations(t)
}

func TestStorageServer_DeletePassword(t *testing.T) {
	mockStorage, _, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.DeletePasswordRequest{
		Key: "test-key",
	}

	mockStorage.On("GetStringData", ctx, "test-user", "test-key").
		Return("test-login test-password", "test-metadata", nil)
	mockStorage.On("DeleteStringData", ctx, "test-user", "test-key").
		Return(nil)

	resp, err := server.DeletePassword(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockStorage.AssertExpectations(t)
}

func TestStorageServer_SaveBinary(t *testing.T) {
	_, mockS3Storage, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.SaveBinaryRequest{
		Key:      "test-key",
		Value:    []byte("test-data"),
		Metadata: stringPtr("test-metadata"),
	}

	mockS3Storage.On("SaveBinaryData", ctx, "test-user", "test-key", []byte("test-data"), "test-metadata").
		Return(nil)

	resp, err := server.SaveBinary(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockS3Storage.AssertExpectations(t)
}

func TestStorageServer_GetBinary(t *testing.T) {
	_, mockS3Storage, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.GetBinaryRequest{
		Key: "test-key",
	}

	mockS3Storage.On("GetBinaryData", ctx, "test-user", "test-key").
		Return([]byte("test-data"), "test-metadata", nil)

	resp, err := server.GetBinary(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockS3Storage.AssertExpectations(t)
}

func TestStorageServer_DeleteBinary(t *testing.T) {
	_, mockS3Storage, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.DeleteBinaryRequest{
		Key: "test-key",
	}

	mockS3Storage.On("DeleteBinaryData", ctx, "test-user", "test-key").
		Return(nil)

	resp, err := server.DeleteBinary(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockS3Storage.AssertExpectations(t)
}

func TestStorageServer_SaveText(t *testing.T) {
	mockStorage, _, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.SaveTextRequest{
		Key:      "test-key",
		Text:     "test-text",
		Metadata: stringPtr("test-metadata"),
	}

	mockStorage.On("GetStringData", ctx, "test-user", "test-key").
		Return("", "", errors.New("not found"))
	mockStorage.On("SaveStringData", ctx, "test-user", "test-key", "test-text", "test-metadata").
		Return(nil)

	resp, err := server.SaveText(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockStorage.AssertExpectations(t)
}

func TestStorageServer_GetText(t *testing.T) {
	mockStorage, _, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.GetTextRequest{
		Key: "test-key",
	}

	mockStorage.On("GetStringData", ctx, "test-user", "test-key").
		Return("text", "test-metadata", nil)

	resp, err := server.GetText(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "text", resp.Text)
	assert.Equal(t, "test-metadata", *resp.Metadata)
	mockStorage.AssertExpectations(t)
}

func TestStorageServer_DeleteText(t *testing.T) {
	mockStorage, _, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.DeleteTextRequest{
		Key: "test-key",
	}

	mockStorage.On("GetStringData", ctx, "test-user", "test-key").
		Return("test-text", "test-metadata", nil)
	mockStorage.On("DeleteStringData", ctx, "test-user", "test-key").
		Return(nil)

	resp, err := server.DeleteText(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockStorage.AssertExpectations(t)
}

func TestStorageServer_SaveBankCard(t *testing.T) {
	mockStorage, _, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.SaveBankCardRequest{
		Key:            "test-key",
		Number:         "test-number",
		Holder:         "test-holder",
		Cvv:            "test-cvv",
		ExpirationDate: "test-expiration-date",
		Metadata:       stringPtr("test-metadata"),
	}

	mockStorage.On("GetStringData", ctx, "test-user", "test-key").
		Return("", "", errors.New("not found"))
	mockStorage.On("SaveStringData", ctx, "test-user", "test-key", "test-number test-holder test-cvv test-expiration-date", "test-metadata").
		Return(nil)

	resp, err := server.SaveBankCard(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockStorage.AssertExpectations(t)
}

func TestStorageServer_GetBankCard(t *testing.T) {
	mockStorage, _, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.GetBankCardRequest{
		Key: "test-key",
	}

	mockStorage.On("GetStringData", ctx, "test-user", "test-key").
		Return("test-number test-holder test-cvv test-expiration-date", "test-metadata", nil)

	resp, err := server.GetBankCard(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "test-number", resp.Number)
	assert.Equal(t, "test-holder", resp.Holder)
	assert.Equal(t, "test-cvv", resp.Cvv)
	assert.Equal(t, "test-expiration-date", resp.ExpirationDate)
	assert.Equal(t, "test-metadata", *resp.Metadata)
	mockStorage.AssertExpectations(t)
}

func TestStorageServer_DeleteBankCard(t *testing.T) {
	mockStorage, _, server := setupStorageTest(t)

	ctx := context.WithValue(context.Background(), "userID", "test-user")
	req := &pbStorage.DeleteBankCardRequest{
		Key: "test-key",
	}

	mockStorage.On("GetStringData", ctx, "test-user", "test-key").
		Return("test-number test-holder test-cvv test-expiration-date", "test-metadata", nil)
	mockStorage.On("DeleteStringData", ctx, "test-user", "test-key").
		Return(nil)

	resp, err := server.DeleteBankCard(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockStorage.AssertExpectations(t)
}

func stringPtr(s string) *string {
	return &s
}
