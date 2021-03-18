package main

import (
	"github.com/keremavci/todo-api/api"
	"github.com/keremavci/todo-api/app"
	"github.com/keremavci/todo-api/config"
	_ "github.com/keremavci/todo-api/log"
	"github.com/keremavci/todo-api/server"
)



func main() {
		app :=&app.App{}
		app.Config = config.DefaultConfig()
		app.Server = server.NewServer(app.Config)
		api.Init(app.Config,app.Server.Router,app.Server.Service)

		app.Server.Start()

}
