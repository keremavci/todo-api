package api

import (
	"github.com/gorilla/mux"
	. "github.com/keremavci/todo-api/log"
	"github.com/keremavci/todo-api/model"
	"github.com/keremavci/todo-api/service"
	"net/http"
)

func (api *API) InitTodo() {
	api.BaseRoutes.Todo.Handle("", api.ApiHandler(getTodo)).Queries("id","{id}").Methods("GET")
	api.BaseRoutes.Todo.Handle("", api.ApiHandler(crateTodo)).Methods("POST")
	api.BaseRoutes.Todo.Handle("", api.ApiHandler(updateTodo)).Methods("PUT")
	api.BaseRoutes.Todo.Handle("/{id}", api.ApiHandler(deleteTodo)).Methods("DELETE")
	api.BaseRoutes.Todo.Handle("", api.ApiHandler(getAllTodo)).Methods("GET")
}


func getTodo(c *Context, w http.ResponseWriter, r *http.Request, service *service.Service) {

	var todoId string
	props := mux.Vars(r)
	if val, ok := props["id"]; ok {
		todoId = val
	}

	todo, err := service.Services.TodoService.GetById(todoId)
	if err != nil {

		c.Err = err
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(todo.ToJson()))
}

func getAllTodo(c *Context, w http.ResponseWriter, r *http.Request, service *service.Service) {

	todoList, err := service.Services.TodoService.GetAll();
	if err != nil {
		c.Err = err
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(model.TodoListToJson(todoList)))

}


func deleteTodo(c *Context, w http.ResponseWriter, r *http.Request, service *service.Service) {


	var todoId string
	params := mux.Vars(r)
	if val, ok := params["id"]; ok {
		todoId = val
	}
	err := service.Services.TodoService.Delete(todoId)
	if err != nil {
		c.Err = err
		return
	}

	w.WriteHeader(http.StatusOK)

}

func crateTodo(c *Context, w http.ResponseWriter, r *http.Request, service *service.Service) {

	todo := model.TodoFromJson(r.Body)
	if ! todo.IsValid() {
		c.SetInvalidParam("todo")
		return
	}
	Logger.Info(todo.Text)
	todo, err := service.Services.TodoService.Save(todo)
	if err != nil {
		c.Err = err
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(todo.ToJson()))

}

func updateTodo(c *Context, w http.ResponseWriter, r *http.Request, service *service.Service){
	todo := model.TodoFromJson(r.Body)
	if ! todo.IsValid() {
		c.SetInvalidParam("todo")
		return
	}
	todo, err := service.Services.TodoService.Update(todo)
	if err != nil {
		c.Err = err
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(todo.ToJson()))
}