package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/controller/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine, userHandler *UserHandler, authMiddleware *middleware.AuthMiddleware) {
	router.Use(middleware.CorsMiddleware())

	v1 := router.Group("/api/v1")
	{
		v1.POST("/users/login", userHandler.Login)
		v1.POST("/users/register", userHandler.Register)

		v1.GET("/users/whoami", authMiddleware.VerifyToken, userHandler.WhoAmI)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
