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

	h.Render(c.Writer, data)
}

// JSON marshals the given interface object and writes a JSON response.
func (c *Context) JSON(status int, data interface{}) {
	head := Head{
		ContentType: ContentJSON + "; charset=" + c.Options.Charset,
		Status:      status,
	}

	j := JSON{
		Head:   head,
		Indent: c.Options.IndentJSON,
		Prefix: c.Options.PrefixJSON,
	}

	j.Render(c.Writer, data)
}

// JSONP marshals the given interface object and writes the JSON response.
func (c *Context) JSONP(status int, callback string, data interface{}) {
	head := Head{
		ContentType: ContentJSONP + "; charset=" + c.Options.Charset,
		Status:      status,
	}

	j := JSONP{
		Head:     head,
		Indent:   c.Options.IndentJSON,
		Callback: callback,
	}

	j.Render(c.Writer, data)
}

// XML marshals the given interface object and writes the XML response.
func (c *Context) XML(status int, data interface{}) {
	head := Head{
		ContentType: ContentXML + "; charset=" + c.Options.Charset,
		Status:      status,
	}

	x := XML{
		Head:   head,
		Indent: c.Options.IndentXML,
		Prefix: c.Options.PrefixXML,
	}

	x.Render(c.Writer, data)
}
