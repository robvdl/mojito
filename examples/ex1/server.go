package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/robvdl/mojito"
)

// Context is your own custom request context, it must embed *mojito.Context
type Context struct {
	*mojito.Context
}

// Home is a simple route based on your own Context.
func (c *Context) Home() {
	c.Logger.Println("Home route")

	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"hello": "world",
	})
}

func main() {
	config, err := mojito.LoadConfig("config.toml")
	if err != nil {
		fmt.Println("Failed to load config file:", err)
		os.Exit(1)
	}

	m := mojito.New(config, Context{})
	m.Get("/", (*Context).Home)

	m.Run()
}
