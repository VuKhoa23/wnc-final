package service

import (
	"github.com/VuKhoa23/advanced-web-be/internal/constants"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/env"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/jwt"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u UserService) Register(c *gin.Context, request model.AuthRequest) error {
	err := u.userRepository.RegisterCommand(c, request.Username, request.Password)
	return err
}

func (u UserService) Login(c *gin.Context, request model.AuthRequest) (entity.User, error) {
	user, err := u.userRepository.LoginCommand(c, request.Username, request.Password)

	jwtSecret, err := env.GetEnv("JWT_SECRET")
	if err != nil {
		return entity.User{}, err
	}
	accessToken, err := jwt.GenerateToken(constants.ACCESS_TOKEN_DURATION, jwtSecret, map[string]interface{}{
		"id": user.Id,
	})

	if err == nil {
		c.SetCookie(
			"access_token",
			accessToken,
			constants.COOKIE_DURATION,
			"/",
			"",
			false,
			true,
		)
	}

	refreshToken, err := jwt.GenerateToken(constants.REFRESH_TOKEN_DURATION, jwtSecret, map[string]interface{}{
		"id": user.Id,
	})
	if err == nil {
		c.SetCookie(
			"refresh_token",
			refreshToken,
			constants.COOKIE_DURATION,
			"/",
			"",
			false,
			true,
		)
	}
	err = u.userRepository.UpdateRefreshToken(c, user.Id, refreshToken)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
