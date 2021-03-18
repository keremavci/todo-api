package service

import (
	"github.com/keremavci/todo-api/store"
)

type Service struct {
	Services     Services
}

type Services struct {
	TodoService *TodoService
}

func NewService(sqlStore *store.SqlStore) *Service {
	service := &Service{}
	service.Services.TodoService = NewTodoService(sqlStore);

	return service

}