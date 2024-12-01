package db

import (
	"context"
)

func (s *dbStorage) GetStringData(ctx context.Context, userId, key string) (string, string, error) {
	row := s.db.QueryRowContext(ctx, "SELECT value, metadata from storage WHERE user_id = $1 AND key = $2", userId, key)

	var data string
	var metadata string
	err := row.Scan(&data, &metadata)

	return data, metadata, err
}

func (s *dbStorage) SaveStringData(ctx context.Context, userId, key, stringData, metadata string) error {
	_, err := s.db.ExecContext(
		ctx,
		"INSERT INTO storage (key, user_id, value, metadata) VALUES($1, $2, $3, $4)",
		key, userId, stringData, metadata,
	)

	return err
}

func (s *dbStorage) DeleteStringData(ctx context.Context, userId, key string) error {
	_, err := s.db.ExecContext(
		ctx,
		"DELETE FROM storage WHERE user_id = $1 AND key = $2",
		userId, key,
	)

	return err
}
