package repositoryimplement

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/database_todo"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db database_todo.Db) repository.TodoRepository {
	return &TodoRepository{db: db}
}

func (repo *TodoRepository) AddNewTodo(c context.Context, todo *entity.Todo) (int64, error) {
	err := repo.db.WithContext(c).Create(todo).Error
	if err != nil {
		return 0, err
	}
	return todo.Id, nil
}

func (repo *TodoRepository) UpdateTodo(c context.Context, todo *entity.Todo, todoId int64, userId int64) (*entity.Todo, error) {
	// Cập nhật chỉ những bản ghi có todoId và userId được gửi vào
	if err := repo.db.WithContext(c).
		Model(&entity.Todo{}).
		Where("id = ? AND user_id = ?", todoId, userId).
		Updates(todo).Error; err != nil {
		return nil, err
	}

	// Lấy bản ghi vừa cập nhật
	var updatedTodo entity.Todo
	if err := repo.db.WithContext(c).
		Where("id = ? AND user_id = ?", todoId, userId).
		First(&updatedTodo).Error; err != nil {
		return nil, err
	}

	return &updatedTodo, nil
}

func (repo *TodoRepository) GetListTodo(c context.Context, userId uint64, searchText string) ([]entity.Todo, error) {
	var todoList []entity.Todo

	query := repo.db.WithContext(c).Where("user_id = ?", userId)
	if searchText != "" {
		query = query.Where("content LIKE ?", "%"+searchText+"%")
	}

	if err := query.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return todoList, nil
}
