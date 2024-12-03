package repository

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type UserRepository interface {
	RegisterCommand(c context.Context, username string, password string) error
	LoginCommand(c context.Context, username string, password string) (entity.User, error)
	UpdateRefreshToken(c context.Context, userId uint64, token string) error
	ValidateRefreshToken(c context.Context, userId uint64, token string) bool
}
