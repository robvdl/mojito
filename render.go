package mojito

import (
	"net/http"
	"path/filepath"

	"github.com/flosch/pongo2"
)

// Context is the application context
type Context struct {
	Options *Options
	Request *http.Request
	Writer  http.ResponseWriter
}

// Vars is a map for template variables, it is the same as pongo2.Context
type Vars map[string]interface{}

// HTML renders a template and returns the output as an HTML response
func (c *Context) HTML(status int, template string, vars map[string]interface{}) {
	tpl, err := pongo2.FromFile(filepath.Join(c.Options.TemplateDirectory, template))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	err = tpl.ExecuteWriter(vars, c.Writer)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	// TODO: status code and headers are not handled yet
	// see implementation in unrolled/render
}
