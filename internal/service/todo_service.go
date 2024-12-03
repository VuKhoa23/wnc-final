package service

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
)

type TodoService interface {
	AddNewTodo(ctx context.Context, todoRequest *model.TodoRequest) (*entity.Todo, error)
	UpdateTodo(ctx context.Context, todoRequest *model.TodoRequest, todoId int64) (*entity.Todo, error)
	GetListTodo(ctx context.Context, userId uint64, searchText string) ([]entity.Todo, error)
}
