package userservice

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"gophkeeper/internal/logger"
	"gophkeeper/internal/models"
)

var ErrorLoginExists = fmt.Errorf("login exists")
var ErrorWrongCredentials = fmt.Errorf("wrong pair login/password")

type storage interface {
	CreateUser(ctx context.Context, tx *sql.Tx, login, password string) error
	BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)
	GetUserByLoginForUpdate(ctx context.Context, tx *sql.Tx, login string) (user models.User, err error)
	GetUserByLogin(ctx context.Context, login string) (user models.User, err error)
}

type UserService struct {
	storage storage
}

func NewUserService(storage storage) *UserService {
	return &UserService{storage}
}

func GetHashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	hp := h.Sum(nil)
	return hex.EncodeToString(hp)
}

func (s *UserService) Register(ctx context.Context, login, password string) error {
	tx, err := s.storage.BeginTx(ctx, nil)

	if err != nil {
		logger.Log.Error("Error while begin transaction: ", err)
		return err
	}

	defer tx.Rollback()

	_, err = s.storage.GetUserByLoginForUpdate(ctx, tx, login)

	if err == nil {
		return ErrorLoginExists
	}

	if !errors.Is(err, sql.ErrNoRows) {
		logger.Log.Error("Error while get user by login: ", err)
		return err
	}

	err = s.storage.CreateUser(ctx, tx, login, GetHashPassword(password))

	if err != nil {
		logger.Log.Error("Error while create user: ", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		logger.Log.Error("Error while commit transaction: ", err)
		return err
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, login, password string) error {
	user, err := s.storage.GetUserByLogin(ctx, login)

	if err != nil {
		logger.Log.Error("Error while get user by login: ", err)
		return ErrorWrongCredentials
	}

	if user.Password != GetHashPassword(password) {
		return ErrorWrongCredentials
	}

	return nil
}
