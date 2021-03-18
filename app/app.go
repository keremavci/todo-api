package app

import (
	"github.com/keremavci/todo-api/config"
	"github.com/keremavci/todo-api/server"
)

type App struct {
	Server *server.Server
	Config *config.Config
}