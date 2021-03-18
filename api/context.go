package api

import (
	"github.com/keremavci/todo-api/log"
	"github.com/keremavci/todo-api/model"
	"github.com/keremavci/todo-api/service"
	"net/http"
)

type Context struct {
  	RequestId string
	IpAddress string
	Path      string
	Err       *model.AppError

}


type handler struct {
	handleFunc func(*Context, http.ResponseWriter, *http.Request, *service.Service)
	service *service.Service
}


func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Logger.Info("First")
	c := &Context{}
	c.RequestId = model.NewId()
	log.Logger.Infof("Starting Request.Request Id:%s", c.RequestId)


	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token,X-Request-Id, Authorization")


	service:= h.service
	writer.Header().Set("Content-Type","application/json")
	writer.Header().Set("X-Request-Id", c.RequestId)

	if c.Err == nil {
		h.handleFunc(c, writer, request, service)
	}

	if c.Err != nil {
		c.Err.RequestId = c.RequestId
		c.Err.Where = request.URL.Path
		log.Logger.Error(c.Err.Where)
		writer.WriteHeader(c.Err.StatusCode)
		writer.Write([]byte(c.Err.ToJson()))
	}

	log.Logger.Infof("Finished Request.Request Id:%s", c.RequestId)

}


func (c *Context) SetInvalidParam(parameter string) {
	c.Err =  model.NewAppError("Context", map[string]interface{}{"Name": parameter}, "", http.StatusBadRequest)
}
