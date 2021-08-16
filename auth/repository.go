package auth

import (
	"context"

	"./models"
)

type Repository interface {
	Insert(ctx context.Context, user *models.User) error
	Get(ctx context.Context, username, password string) (*models.User, error)
}
