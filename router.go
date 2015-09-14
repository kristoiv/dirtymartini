package dirty

import (
	"net/http"

	"github.com/go-martini/martini"
)

type Router interface {
	martini.Router
}

type router struct {
	martini.Router
}

func NewRouter() Router {
	return &router{martini.NewRouter()}
}

type Helper interface {
	URLFor(name string, params ...interface{}) string
}

type helper struct {
	Routes  martini.Routes
	BaseUrl string
}

func (h *helper) URLFor(name string, params ...interface{}) string {
	return h.BaseUrl + h.Routes.URLFor(name, params...)
}

func HelperMiddleware() func(c martini.Context, routes martini.Routes, req *http.Request) {
	return func(c martini.Context, routes martini.Routes, req *http.Request) {

		baseUrl := "http://" + req.Host
		if req.URL.IsAbs() {
			baseUrl = req.URL.Scheme + req.URL.Host
		}

		helper := &helper{
			routes,
			baseUrl,
		}

		c.Map(helper)

	}
}
