package mojito

import (
	"net/http"
	"path/filepath"

	"github.com/flosch/pongo2"
)

const (
	// ContentType header constant.
	ContentType = "Content-Type"
	// ContentLength header constant.
	ContentLength = "Content-Length"
	// ContentBinary header value for binary data.
	ContentBinary = "application/octet-stream"
	// ContentJSON header value for JSON data.
	ContentJSON = "application/json"
	// ContentJSONP header value for JSONP data.
	ContentJSONP = "application/javascript"
	// ContentHTML header value for HTML data.
	ContentHTML = "text/html"
	// ContentXHTML header value for XHTML data.
	ContentXHTML = "application/xhtml+xml"
	// ContentXML header value for XML data.
	ContentXML = "text/xml"
	// Default character encoding.
	defaultCharset = "UTF-8"
)

// Context is a render context to view functions
type Context struct {
	Options *Options
	Request *http.Request
	Writer  http.ResponseWriter
}

// HTML renders a template and returns the output as an HTML response
func (c *Context) HTML(status int, template string, data map[string]interface{}) {
	// Load Pongo2 template
	tpl, err := pongo2.FromFile(filepath.Join(c.Options.TemplateDirectory, template))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	head := Head{
		ContentType: c.Options.HTMLContentType + "; charset=" + c.Options.Charset,
		Status:      status,
	}

	h := HTML{
		Head:     head,
		Name:     template,
		Template: tpl,
	}

	head.Write(c.Writer)
	h.Render(c.Writer, data)
}

// JSON renders a JSON response.
func (c *Context) JSON(status int, data interface{}) {

}
