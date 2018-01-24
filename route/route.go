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

// New new router
func New(c *controller.Controller) *Router {
	return &Router{
		r:      mux.NewRouter(),
		c:      c,
		isInit: false,
	}
}

// Init 注册路由
func (r *Router) Init() {
	if !r.isInit {
		r.isInit = true
		r.r.Handle("/{shroturl}", nil).Methods(http.MethodGet)
		r.r.Handle("/", nil).Methods(http.MethodPost)
		http.Handle("r", r.r)
	}
}

// Run listern http
func (r *Router) Run() error {
	if !r.isInit {
		r.Init()
	}
	if err := http.ListenAndServe("", nil); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
