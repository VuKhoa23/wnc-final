package service

import (
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(c *gin.Context, request model.AuthRequest) error
	Login(c *gin.Context, request model.AuthRequest) error
}
