package main

import (
	"fmt"

	"github.com/robvdl/mojito"
)

// Context is your custom request context, it extends *mojito.Context
type Context struct {
	*mojito.Context
}

// APIContext is a second context only used by the API subrouter
type APIContext struct {
	*Context
}

func (c *Context) appMiddleware(rw mojito.ResponseWriter, req *mojito.Request, next mojito.NextMiddlewareFunc) {
	c.Config.Logger.Println("App middleware")
	next(rw, req)
}

func (c *Context) testRoute(rw mojito.ResponseWriter, req *mojito.Request) {
	fmt.Fprint(rw, "Test Route")
}

func (c *APIContext) apiRoute(rw mojito.ResponseWriter, req *mojito.Request) {
	fmt.Fprint(rw, "API Route")
}

func (c *APIContext) apiMiddleware(rw mojito.ResponseWriter, req *mojito.Request, next mojito.NextMiddlewareFunc) {
	c.Config.Logger.Println("API middleware")
	next(rw, req)
}

func main() {
	m := mojito.Classic(Context{})
	m.Middleware((*Context).LoggerMiddleware)
	m.Middleware((*Context).ShowErrorsMiddleware)
	m.Middleware((*Context).appMiddleware)
	m.Get("/", (*Context).testRoute)

	api := m.Subrouter(APIContext{}, "/api")
	api.Middleware((*APIContext).apiMiddleware)
	api.Get("/tickets", (*APIContext).apiRoute)

	m.Run("localhost:8000")
}
