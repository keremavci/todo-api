package service

import (
	"database/sql"
	"github.com/keremavci/todo-api/model"
	"github.com/keremavci/todo-api/store"
	"net/http"
)

type TodoService struct {
	sqlStore *store.SqlStore
}

var ToDoService *TodoService
func NewTodoService(sqlStore *store.SqlStore) *TodoService{

	ToDoService := &TodoService{
		sqlStore: sqlStore,
	}

	return ToDoService
}

func(todoService *TodoService)Save(todo *model.Todo)(*model.Todo, *model.AppError){
	t,err := todoService.sqlStore.TodoStore().Save(todo)
	if err != nil {
		return nil, model.NewAppError("Save Todo", map[string]interface{}{"todo": todo},err.Error(),http.StatusInternalServerError)
	}
	return t,nil
}


func(todoService *TodoService)Update(todo *model.Todo)(*model.Todo, *model.AppError){
	_,err:=todoService.sqlStore.TodoStore().GetById(todo.Id)
	if err != nil {
		return nil, model.NewAppError("Update Todo",map[string]interface{}{"todo": todo}, err.Error(),http.StatusNotFound)
	}
	t,err := todoService.sqlStore.TodoStore().Update(todo)
	if err != nil {
		return nil, model.NewAppError("Update Todo",map[string]interface{}{"todo": todo}, err.Error(),http.StatusInternalServerError)
	}
	return t,nil
}


func(todoService *TodoService)Delete(todoId string)(*model.AppError){
	err := todoService.sqlStore.TodoStore().Delete(todoId)
	if err != nil {
		return model.NewAppError("Delete Todo",map[string]interface{}{"todoId": todoId}, err.Error(),http.StatusInternalServerError)
	}
	return nil
}

func(todoService *TodoService)GetAll()([]*model.Todo, *model.AppError){
	todoList, err := todoService.sqlStore.TodoStore().GetAll()
	if err != nil {
		return nil,model.NewAppError("Delete Todo",nil, err.Error(),http.StatusInternalServerError)
	}
	return todoList,nil
}


func(todoService *TodoService)GetById (todoId string)(*model.Todo, *model.AppError){
	todo, err := todoService.sqlStore.TodoStore().GetById(todoId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil,model.NewAppError("Get Todo",map[string]interface{}{"todo_id": todoId}, "Todo not found.",http.StatusNotFound)
		}
		return nil,model.NewAppError("Get Todo",map[string]interface{}{"todo_id": todoId}, err.Error(),http.StatusInternalServerError)
	}
	return todo, nil
}