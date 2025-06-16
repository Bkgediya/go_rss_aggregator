package db

import (
	"context"

	"github.com/Bkgediya/go_rss_aggregator/internal/model"
)

func CreateUser(ctx context.Context, user model.User) (int, error) {
	var id int
	query := "INSERT INTO users(name,email) VALUES ($1, $2) RETURNING id"

	err := DB.QueryRow(ctx, query, user.Name, user.Email).Scan(&id)

	return id, err
}

func GetUserByID(ctx context.Context, id int) (model.User, error) {
	var user model.User
	query := "SELECT id, name, email FROM users WHERE id = $1"

	err := DB.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)

	return user, err
}
