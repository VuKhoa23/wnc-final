package http

import (
	"fmt"
	"github.com/VuKhoa23/advanced-web-be/internal/controller/http/middleware"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	v1 "github.com/VuKhoa23/advanced-web-be/internal/controller/http/v1"
)

type Server struct {
	userHandler    *v1.UserHandler
	authMiddleware *middleware.AuthMiddleware
}

func NewServer(userHandler *v1.UserHandler, authMiddleware *middleware.AuthMiddleware) *Server {
	return &Server{userHandler: userHandler, authMiddleware: authMiddleware}
}

func (s *Server) Run() {
	router := gin.New()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	httpServerInstance := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	v1.MapRoutes(router, s.userHandler, s.authMiddleware)
	err := httpServerInstance.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Println("Server running at " + httpServerInstance.Addr)
}
