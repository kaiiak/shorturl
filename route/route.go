package route

import (
	"fmt"

	"github.com/kaiiak/shorturl/config"
	"github.com/kaiiak/shorturl/controller"
	"github.com/labstack/echo"
)

// Router 路由
type Router struct {
	port   int
	r      *echo.Echo
	c      *controller.Controller
	isInit bool
}

// New new router
func New(c *controller.Controller, cnf *config.Config) *Router {
	return &Router{
		port:   cnf.Port,
		r:      echo.New(),
		c:      c,
		isInit: false,
	}
}

// Init 注册路由
func (r *Router) Init() {
	if !r.isInit {
		r.isInit = true
		r.r.POST("/", r.c.SetRawURL)
		r.r.GET("/:shroturl", r.c.GetRawURL)
	}
}

// Run listern http
func (r *Router) Run() error {
	if !r.isInit {
		r.Init()
	}
	return r.r.Start(fmt.Sprintf(":%d", r.port))
}
