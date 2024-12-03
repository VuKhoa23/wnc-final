package startup

import (
	"github.com/VuKhoa23/advanced-web-be/internal"
	"github.com/VuKhoa23/advanced-web-be/internal/controller"
	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/database_todo"
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.Open()
	db_todo := database_todo.Open()
	return internal.InitializeContainer(db, db_todo)
}

func Execute() {
	container := registerDependencies()
	container.HttpServer.Run()
}
