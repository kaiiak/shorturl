package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaiiak/shorturl/data"
)

// Controller c in mvc
type Controller struct {
	data *data.Data
}

// GetRawURL 获取原始url
func (c *Controller) GetRawURL() http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			surl := mux.Vars(r)["shroturl"]
			rurl, err := c.data.Get(surl)
			if err != nil {
				if err == data.ErrNotFound {
					rw.WriteHeader(http.StatusNotFound)
					return
				}
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.Redirect(rw,r, rurl, http.StatusMovedPermanently)
			return
		})
}
