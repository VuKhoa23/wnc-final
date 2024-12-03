package repositoryimplement

import (
	"context"
	"database/sql"
	"errors"
	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db database.Db) repository.UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) RegisterCommand(c context.Context, username string, password string) error {
	// Check if username already exists
	var existingUser entity.User
	query := "SELECT id FROM users WHERE username = ?"
	err := u.db.GetContext(c, &existingUser, query, username)
	if err != nil && err.Error() != httpcommon.ErrorMessage.SqlxNoRow {
		return err
	}
	if err == nil {
		return errors.New("username already exists")
	}

	// Insert the new user
	insertQuery := `INSERT INTO users (username, password, refresh_token) VALUES (:username, :password, :refresh_token)`
	_, err = u.db.NamedExecContext(c, insertQuery, map[string]interface{}{
		"username":      username,
		"password":      password,
		"refresh_token": nil,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) LoginCommand(c context.Context, username string, password string) (entity.User, error) {
	var user entity.User
	query := "SELECT id, username, password FROM users WHERE username = ?"
	err := u.db.GetContext(c, &user, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errors.New(httpcommon.ErrorMessage.BadCredential)
		}
		return entity.User{}, err
	}

	if user.Password != password {
		return entity.User{}, errors.New(httpcommon.ErrorMessage.BadCredential)
	}

	return user, nil
}

func (u UserRepository) UpdateRefreshToken(c context.Context, userId uint64, token string) error {
	query := `UPDATE users SET refresh_token = :refresh_token WHERE id = :id`

	_, err := u.db.NamedExecContext(c, query, map[string]interface{}{
		"refresh_token": token,
		"id":            userId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) ValidateRefreshToken(c context.Context, userId uint64, token string) bool {
	query := `SELECT refresh_token FROM users WHERE id = ?`

	var storedToken string
	err := u.db.GetContext(c, &storedToken, query, userId)
	if err != nil {
		return false
	}

	return storedToken == token
}
