package dirty

import (
	"log"
	"os"

	"github.com/go-martini/martini"
)

type Dirty struct {
	*martini.Martini
	logger *log.Logger
}

func New() *Dirty {
	d := &Dirty{martini.New(), log.New(os.Stdout, "[dirty] ", 0)}
	d.Map(d.logger)
	return d
}

func (d *Dirty) SetLoggerIdentity(name string) {
	if name == "" {
		d.logger.SetPrefix("")
		return
	}
	d.logger.SetPrefix("[" + name + "] ")
}

type ClassicDirty struct {
	*Dirty
	Router
}

func Classic() *ClassicDirty {

	r := NewRouter()
	d := New()

	d.Use(martini.Logger())
	d.Use(martini.Recovery())
	d.Use(martini.Static("public"))

	d.Use(HelperMiddleware()) // Setup helper variables and methods for this request

	d.MapTo(r, (*martini.Routes)(nil))
	d.Action(r.Handle)

	return &ClassicDirty{d, r}

}
