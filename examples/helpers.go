package main

import (
	"github.com/go-martini/martini"
	"github.com/kristoiv/dirtymartini"
	"github.com/martini-contrib/render"
)

func main() {
	d := dirty.Classic()
	d.Use(render.Renderer())
	d.Group("/api", Routes)
	d.Run()
}

func Routes(router martini.Router) {
	router.Get("/", IndexAction).Name("/api")
	router.Get("/welcome/(?P<name>[a-zA-Z]+)", WelcomeAction).Name("/api/welcome")
}

func IndexAction(helper dirty.Helper, r render.Render) {
	r.JSON(200, map[string]interface{}{
		"Self":    helper.URLFor("/api"),
		"Welcome": helper.URLFor("/api/welcome", "<Test>"),
	})
}

func WelcomeAction(helper dirty.Helper, r render.Render, params martini.Params) {
	name := params["name"]
	r.JSON(200, map[string]interface{}{
		"Message": "Welcome to the API " + name,
		"Self":    helper.URLFor("/api/welcome", name),
	})
}
