package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaiiak/shorturl/controller"
)

// Router 路由
type Router struct {
	r      *mux.Router
	c      *controller.Controller
	isInit bool
}

// Init 注册路由
func (r *Router) Init() {
	r.isInit = true
	r.r.Handle("/{shroturl}", nil).Methods(http.MethodGet)
	r.r.Handle("/", nil).Methods(http.MethodPost)
	http.Handle("r",r.r)
}

// Run listern http
func (r *Router) Run() error {
	if !r.isInit {
		return ErrRouterUninitialized
	}
	return nil
}
