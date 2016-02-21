package main

import (
	"fmt"
	"log"
	"os"

	"github.com/robvdl/mojito"
)

// Context is your custom request context, it extends *mojito.Context
type Context struct {
	*mojito.Context
	Test       int
	Another    bool
	AndAnother int
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
	m := mojito.New(Context{}, &mojito.Config{
		Logger: log.New(os.Stdout, "", 0),
	})

	setupMiddleware(m)
	setupRoutes(m)

	m.Run("localhost:8000")
}
