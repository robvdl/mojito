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

// HTML renders a template and returns the output as an HTML response
func (c *Context) HTML(status int, template string, data interface{}) {
	tpl, err := pongo2.FromFile(filepath.Join(c.Options.TemplateDirectory, template))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	// Pongo2 context must be of type map[string]interface{}, use type assertion
	// to check that the type is correct first and convert it.
	pongoContext, ok := data.(map[string]interface{})
	if ok {
		err = tpl.ExecuteWriter(pongoContext, c.Writer)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	// TODO: status code and headers are not handled yet
	// see implementation in unrolled/render
}

// JSON renders a JSON response.
func (c *Context) JSON(status int, data interface{}) {

}
