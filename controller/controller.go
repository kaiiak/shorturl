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

// GetRawURL return raw url where shorturl
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
	ctx.Redirect(http.StatusMovedPermanently, rurl)
	return nil
}

// SetRawURL rend short and save to database
func (c *Controller) SetRawURL(ctx echo.Context) error {
	var raw struct {
		RawURL string `json:"raw_url"`
	}
	ctx.Bind(&raw)
	if raw.RawURL == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "")
	}
	value, err := c.data.Set(raw.RawURL)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"short_url": value,
	})
}
