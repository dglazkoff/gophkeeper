package db

import (
	"context"
	"database/sql"
	"gophkeeper/internal/logger"
	"gophkeeper/internal/models"
)

func (s *dbStorage) CreateUser(ctx context.Context, tx *sql.Tx, login, password string) error {
	_, err := tx.ExecContext(
		ctx,
		"INSERT INTO users (login, password) VALUES($1, $2) ON CONFLICT (login) DO NOTHING",
		login, password,
	)

	if err != nil {
		logger.Log.Debug("error while creating user ", err)
		return err
	}

	return nil
}

func (s *dbStorage) GetUserByLoginForUpdate(ctx context.Context, tx *sql.Tx, login string) (user models.User, err error) {
	row := tx.QueryRowContext(ctx, "SELECT login, password from users WHERE login = $1 FOR UPDATE", login)
	err = row.Scan(&user.Login, &user.Password)

	return
}

func (s *dbStorage) GetUserByLogin(ctx context.Context, login string) (user models.User, err error) {
	row := s.db.QueryRowContext(ctx, "SELECT login, password from users WHERE login = $1", login)
	err = row.Scan(&user.Login, &user.Password)

	return
}
