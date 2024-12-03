package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/controller/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine, userHandler *UserHandler, todoHandler *TodoHandler, authMiddleware *middleware.AuthMiddleware) {
	router.Use(middleware.CorsMiddleware())

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", userHandler.Login)
			auth.POST("/register", userHandler.Register)
		}

		todos := v1.Group("/todos")
		todos.Use(authMiddleware.VerifyToken)
		{
			todos.GET("", todoHandler.GetList)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
