package api

import (
	"github.com/gorilla/mux"
	"github.com/keremavci/todo-api/config"
	"github.com/keremavci/todo-api/service"
	"net/http"
	"net/http/pprof"
)

type Routes struct {
	ApiRoot *mux.Router // 'api'
	Todo   *mux.Router // 'api/todo'
	Status *mux.Router // 'api/status'
}

type API struct {
	BaseRoutes          *Routes
	Service *service.Service
}

func (api *API) ApiHandler(h func(*Context, http.ResponseWriter, *http.Request, *service.Service)) http.Handler {

	return &handler{
		handleFunc: h,
		service: api.Service,
	}
}

func Init(config *config.Config,root *mux.Router,service *service.Service) *API {
	api := &API{
		BaseRoutes: &Routes{},
		Service: service,
	}


	api.BaseRoutes.ApiRoot = root.PathPrefix(config.APIBaseUri).Subrouter()
	api.BaseRoutes.Status = api.BaseRoutes.ApiRoot.PathPrefix("/status").Subrouter()
	api.BaseRoutes.Todo = api.BaseRoutes.ApiRoot.PathPrefix("/todo").Subrouter()

	api.InitStatus()
	api.InitTodo()

	AttachProfiler(root)


	root.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	return api
}

func AttachProfiler(router *mux.Router) {
	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}



