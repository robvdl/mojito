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

// Render is the generic function called by XML, JSON, Data, HTML, and can be
// called by custom implementations.
func (c *Context) Render(w http.ResponseWriter, e Engine, data interface{}) {
	err := e.Render(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Data writes out the raw bytes as binary data.
func (c *Context) Data(status int, data []byte) {
	head := Head{
		ContentType: ContentBinary,
		Status:      status,
	}

	d := Data{
		Head: head,
	}

	c.Render(c.Writer, d, data)
}

// HTML renders a template and returns the output as an HTML response
func (c *Context) HTML(status int, name string, data map[string]interface{}) {
	// Load Pongo2 template
	if tpl, err := pongo2.FromFile(filepath.Join(c.Options.TemplateDirectory, name)); err == nil {
		head := Head{
			ContentType: c.Options.HTMLContentType + "; charset=" + c.Options.Charset,
			Status:      status,
		}

		h := HTML{
			Head:     head,
			Name:     name,
			Template: tpl,
		}

		c.Render(c.Writer, h, data)
	} else {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
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

	c.Render(c.Writer, j, data)
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

	c.Render(c.Writer, j, data)
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

	c.Render(c.Writer, x, data)
}

// Markdown renders Markdown content and writes the HTML response.
func (c *Context) Markdown(status int, data []byte) {
	head := Head{
		ContentType: ContentHTML + "; charset=" + c.Options.Charset,
		Status:      status,
	}

	m := Markdown{
		Head: head,
	}

	c.Render(c.Writer, m, data)
}
