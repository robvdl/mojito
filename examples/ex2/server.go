package main

import (
	"fmt"
	"os"

	"github.com/robvdl/mojito"
)

// Context is your own custom request context, it must embed *mojito.Context
type Context struct {
	*mojito.Context
}

// Configure gets called automatically by mojito.New().
func (c Context) Configure() *mojito.Config {
	config, err := mojito.LoadConfig("config.toml")
	if err != nil {
		fmt.Println("Failed to load config file:", err)
		os.Exit(1)
	}
	return config
}

// Home is a simple route based on your own Context.
func (c *Context) Home() {
	c.Logger.Println("Home route")
	// c.ResponseWriter
	// c.Request
	// c.HTML(http.StatusOK, "index.html", map[string]interface{}{"hello": "html"})
	// c.JSON(http.StatusOK, map[string]string{"hello": "json"})
	// c.Config
	// c.User
}

func main() {
	m := mojito.New(Context{})
	// m.use(mojito.SessionMiddleware)
	// m.Use(mojito.AuthMiddleware)
	// m.Get("/", (*Context).Home)
	// admin := m.SubRouter("/admin", AdminContext{})
	// admin.Get("/", (*AdminContext).Home)

	m.Run()
}
