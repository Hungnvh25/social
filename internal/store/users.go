package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int16  `json: "id"`
	UserName  string `json: "user_name"`
	Email     string `json: "email"`
	Password  string `json: "-"`
	CreatedAt string `json: "created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users(username,password,email) VALUES($1,$2,$3) returning id,
	created_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		user.UserName,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
