package mojito

import (
	"fmt"
	"reflect"
)

type Application interface {
	Configure() *Config
}

// New creats a Mojito application by creating the root Router.
func New(app Application) *Router {
	config := app.Configure()

	r := &Router{
		pathPrefix:  "/",
		contextType: reflect.TypeOf(app),
		Config:      config,
		Logger:      GetLogger(config),
	}

	return r
}

// SubRouter adds a new subrouter to an existing Router.
func (r *Router) SubRouter(ctx interface{}) *Router {
	// TODO: attach properly to base Router
	subrouter := &Router{
		Config: r.Config,
	}

	return subrouter
}

// Run starts the application and begins serving.
func (r *Router) Run() {
	host := r.Config.GetString("server.host")
	port := r.Config.GetInt("server.port")
	useTLS := r.Config.GetBool("server.tls")
	address := fmt.Sprintf("%s:%d", host, port)
	r.Logger.Printf("Running on %s, tls=%t", address, useTLS)
}
