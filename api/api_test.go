package api

import (
	"github.com/keremavci/todo-api/app"
	"github.com/keremavci/todo-api/config"
	"github.com/keremavci/todo-api/helper"
	. "github.com/keremavci/todo-api/log"
	"github.com/keremavci/todo-api/server"

	"net/http/httptest"
)

func SetupTestApp() *app.App {
	app := &app.App{}
	app.Config = config.DefaultConfig()
	mySqlConnStr, err := helper.CreateMySqlContainerForTest()
	if err != nil {
		Logger.Fatalf("Cannot give valid mysql connection string: Error: %s", err.Error())
	}
	app.Config.DBConfig.MySqlConnectionString=mySqlConnStr
	app.Server = server.NewServer(app.Config)
	return app
}

func SetupTest() *httptest.Server{
	app:=SetupTestApp()
	Init(app.Config,app.Server.Router,app.Server.Service)
	server:=httptest.NewServer(app.Server.Router)
	return server
}
