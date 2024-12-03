package serviceimplement

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
)

type TodoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) service.TodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (service *TodoService) AddNewTodo(ctx context.Context, todoRequest *model.TodoRequest) (*entity.Todo, error) {
	todo := &entity.Todo{
		Content:   todoRequest.Content,
		UserId:    todoRequest.UserId,
		Completed: false,
	}
	todoId, err := service.todoRepository.AddNewTodo(ctx, todo)
	if err != nil {
		return nil, err
	}
	todo.Id = todoId
	return todo, nil
}

func (service *TodoService) UpdateTodo(ctx context.Context, todoRequest *model.TodoRequest, todoId int64) (*entity.Todo, error) {
	todo := &entity.Todo{
		Content:   todoRequest.Content,
		UserId:    todoRequest.UserId,
		Completed: todoRequest.Completed,
	}
	updateTodo, err := service.todoRepository.UpdateTodo(ctx, todo, todoId, todoRequest.UserId)
	if err != nil {
		return nil, err
	}
	return updateTodo, nil
}

func (service *TodoService) GetListTodo(ctx context.Context, userId uint64, searchText string) ([]entity.Todo, error) {
	todoList, err := service.todoRepository.GetListTodo(ctx, userId, searchText)
	if err != nil {
		return nil, err
	}
	return todoList, nil
}
