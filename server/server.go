package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/keremavci/todo-api/config"
	. "github.com/keremavci/todo-api/log"
	"github.com/keremavci/todo-api/service"
	"github.com/keremavci/todo-api/store"

	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	Router     *mux.Router
	Server     *http.Server
	Store *store.SqlStore
	Service *service.Service

}

func NewServer(config *config.Config) *Server {
	srv := &Server{
		Router: mux.NewRouter(),
	}

	srv.Store = store.NewSqlStore(config)
	srv.Service = service.NewService(srv.Store)
	srvAddr := config.ServerConfig.Host + ":" + config.ServerConfig.Port
	srv.Server = &http.Server{Addr: srvAddr, Handler: srv.Router}

	return srv
}
func (srv *Server) Start() {
	go func() {
		if err := srv.Server.ListenAndServe(); err != http.ErrServerClosed {
			Logger.Panic(err.Error())
		}
	}()
	Logger.Info("Starting Server")
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	_ = <-quit
	Logger.Info("Shutting down server... Reason:")
	if err := srv.Server.Shutdown(context.Background()); err != nil {
		Logger.Fatalf(err.Error())
	}
	Logger.Info("Server gracefully stopped")
}