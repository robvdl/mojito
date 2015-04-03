package mojito

import (
	"fmt"
	"net/http"
)

// Context is the application context
type Context struct {
	Options *Options
	Request *http.Request
	Writer  *http.ResponseWriter
}

func (c *Context) HTML(status int, template string, binding map[string]interface{}) {
	fmt.Println("Render template here")
}
