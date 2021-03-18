package api

import (
	"github.com/keremavci/todo-api/service"
	"net/http"
)

func (api *API) InitStatus() {

	api.BaseRoutes.Status.Handle("", api.ApiHandler(getStatus)).Methods("GET","OPTIONS")

}


func getStatus(c *Context, w http.ResponseWriter, r *http.Request, service *service.Service) {
	w.Write([]byte("OK"))
}