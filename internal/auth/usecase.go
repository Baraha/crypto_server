package auth

import (
	"context"

	"github.com/Baraha/crypto_server.git/models"
)

type UseCase interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, user *models.User) (string, error)
}
