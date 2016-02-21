package main

import (
	"fmt"

	"github.com/robvdl/mojito"
)

// Context is your custom request context, it extends *mojito.Context
type Context struct {
	*mojito.Context
}

func (c *Context) testRoute(rw mojito.ResponseWriter, req *mojito.Request) {
	fmt.Fprint(rw, "Test Route")
}

func setupMiddleware(r *mojito.Router) {
	r.Middleware((*Context).LoggerMiddleware)
	r.Middleware((*Context).ShowErrorsMiddleware)
}

func setupRoutes(r *mojito.Router) {
	r.Get("/", (*Context).testRoute)
}

func main() {
	m := mojito.Classic(Context{})

	setupMiddleware(m)
	setupRoutes(m)

	m.Run("localhost:8000")
}
