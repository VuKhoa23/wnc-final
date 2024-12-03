//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/VuKhoa23/advanced-web-be/internal/controller"
	"github.com/VuKhoa23/advanced-web-be/internal/controller/http"
	"github.com/VuKhoa23/advanced-web-be/internal/controller/http/middleware"
	v1 "github.com/VuKhoa23/advanced-web-be/internal/controller/http/v1"
	"github.com/VuKhoa23/advanced-web-be/internal/database"
	repositoryimplement "github.com/VuKhoa23/advanced-web-be/internal/repository/implement"
	serviceimplement "github.com/VuKhoa23/advanced-web-be/internal/service/implement"
	"github.com/google/wire"
)

var container = wire.NewSet(
	controller.NewApiContainer,
)

// may have grpc server in the future
var serverSet = wire.NewSet(
	http.NewServer,
)

// handler === controller | with service and repository layers to form 3 layers architecture
var handlerSet = wire.NewSet(
	v1.NewUserHandler,
	v1.NewTodoHandler,
)

var serviceSet = wire.NewSet(
	serviceimplement.NewUserService,
)

var repositorySet = wire.NewSet(
	repositoryimplement.NewUserRepository,
)

var middlewareSet = wire.NewSet(
	middleware.NewAuthMiddleware,
)

func InitializeContainer(
	db database.Db,
) *controller.ApiContainer {
	wire.Build(serverSet, handlerSet, serviceSet, repositorySet, middlewareSet, container)
	return &controller.ApiContainer{}
}
