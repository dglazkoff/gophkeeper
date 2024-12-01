package db

import (
	"context"
	"database/sql"
	"gophkeeper/internal/logger"
)

type dbStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *dbStorage {
	return &dbStorage{db: db}
}

func (s *dbStorage) Bootstrap() error {
	_, err := s.db.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"login VARCHAR(250) PRIMARY KEY, " +
		"password VARCHAR(250) NOT NULL" +
		")")

	if err != nil {
		logger.Log.Error("Error while create table users: ", err)
		return err
	}

	_, err = s.db.Exec("CREATE TABLE IF NOT EXISTS storage (" +
		"key VARCHAR(250) NOT NULL, " +
		"user_id VARCHAR(250) NOT NULL, " +
		"value VARCHAR(250) NOT NULL, " +
		"metadata VARCHAR(250), " +
		"PRIMARY KEY(key, user_id)" +
		")")

	if err != nil {
		logger.Log.Error("Error while create table storage: ", err)
		return err
	}

	return nil
}

func (s *dbStorage) BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return s.db.BeginTx(ctx, options)
}
