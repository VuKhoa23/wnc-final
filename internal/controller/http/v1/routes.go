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
		v1.GET("/todos", authMiddleware.VerifyToken, todoHandler.GetList)
		v1.POST("/todos", authMiddleware.VerifyToken, todoHandler.Add)
		v1.PUT("/todos", authMiddleware.VerifyToken, todoHandler.Update)
		v1.POST("/users/login", userHandler.Login)
		v1.POST("/users/register", userHandler.Register)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
