# dirtymartini
Martini is a great Go(lang) library for creating web-related software. A part of what makes it great is that it is very clean and contains few features, making it easy to read and understand the source code. However, to increase productivity and promote code reuse in my projects I wanted to dirty my Martini up a bit. This is my Dirty Martini library which adds some useful features to the Martini library that fits my purposes.


## Examples

```go
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


```
