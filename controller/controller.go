package controller

import (
	"log"
	"net/http"

	"github.com/kaiiak/shorturl/data"
	"github.com/labstack/echo"
)

// Controller c in mvc
type Controller struct {
	data *data.Data
}

// New new controller
func New(d *data.Data) *Controller {
	return &Controller{d}
}

// GetRawURL 获取原始url
func (c *Controller) GetRawURL(ctx echo.Context) error {
	surl := ctx.Param("shorturl")
	rurl, err := c.data.Get(surl)
	if err != nil {
		if err == data.ErrNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		log.Printf("get [%s], error [%s]", surl, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// http.Redirect(rw, r, rurl, http.StatusMovedPermanently)
	ctx.Redirect(http.StatusMovedPermanently, rurl)
	return nil
}

// func (c *Controller) SetRawURL() http.Handler {
// 	return http.hf
// }
