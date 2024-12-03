package api

import (
	"context"
	"database/sql"
	"gophkeeper/internal/logger"
	userservice "gophkeeper/internal/service/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"gophkeeper/internal/db"
	pbUser "gophkeeper/internal/proto/user"
)

func setupUserTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *UserServer) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	dbStorage := db.New(mockDB)
	userService := userservice.NewUserService(dbStorage)
	server := NewUserServer(userService)

	return mockDB, mock, server
}

func TestUserServer_RegisterUser(t *testing.T) {
	err := logger.Initialize()
	assert.NoError(t, err)

	db, mock, server := setupUserTest(t)
	defer db.Close()

	ctx := context.Background()
	login := "test-login"
	password := "test-password"

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT login, password from users").
		WithArgs(login).
		WillReturnRows(sqlmock.NewRows([]string{"login", "password"}))

	mock.ExpectExec("INSERT INTO users").
		WithArgs(login, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	resp, err := server.RegisterUser(ctx, &pbUser.RegisterUserRequest{
		Login:    login,
		Password: password,
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GetAccessToken())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserServer_LoginUser(t *testing.T) {
	err := logger.Initialize()
	assert.NoError(t, err)

	db, mock, server := setupUserTest(t)
	defer db.Close()

	ctx := context.Background()
	login := "test-login"
	password := "test-password"
	hashedPassword := userservice.GetHashPassword(password)

	mock.ExpectQuery("SELECT login, password from users").
		WithArgs(login).
		WillReturnRows(sqlmock.NewRows([]string{"login", "password"}).
			AddRow(login, hashedPassword))

	resp, err := server.LoginUser(ctx, &pbUser.LoginUserRequest{
		Login:    login,
		Password: password,
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GetAccessToken())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserServer_RegisterUser_EmptyCredentials(t *testing.T) {
	_, _, server := setupUserTest(t)

	ctx := context.Background()
	resp, err := server.RegisterUser(ctx, &pbUser.RegisterUserRequest{
		Login:    "",
		Password: "",
	})

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUserServer_LoginUser_EmptyCredentials(t *testing.T) {
	_, _, server := setupUserTest(t)

	ctx := context.Background()
	resp, err := server.LoginUser(ctx, &pbUser.LoginUserRequest{
		Login:    "",
		Password: "",
	})

	assert.Error(t, err)
	assert.Nil(t, resp)
}
