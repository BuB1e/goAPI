package repositories

import (
	"context"

	"MODULE_NAME/domain/models"
	"MODULE_NAME/domain/requests"
)

type UserRepository interface {
	Create(ctx context.Context, req *requests.UserRegisterRequest) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}