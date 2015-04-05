package mojito

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/russross/blackfriday"
)

// Engine is the generic interface for all responses.
type Engine interface {
	Render(http.ResponseWriter, interface{}) error
}

// Head defines the basic ContentType and Status fields.
type Head struct {
	ContentType string
	Status      int
}

// Data built-in renderer.
type Data struct {
	Head
}

// HTML built-in Pongo2 renderer.
type HTML struct {
	Head
	Name     string
	Template *pongo2.Template
}

// JSON built-in renderer.
type JSON struct {
	Head
	Indent bool
	Prefix []byte
}

// JSONP built-in renderer.
type JSONP struct {
	Head
	Indent   bool
	Callback string
}

// XML built-in renderer.
type XML struct {
	Head
	Indent bool
	Prefix []byte
}

// Markdown built-in renderer using Blackfriday.
type Markdown struct {
	Head
	Name string
}

// Write outputs the header content.
func (h Head) Write(w http.ResponseWriter) {
	w.Header().Set(ContentType, h.ContentType)
	w.WriteHeader(h.Status)
}

// Render a data response.
func (d Data) Render(w http.ResponseWriter, data interface{}) error {
	c := w.Header().Get(ContentType)
	if c != "" {
		d.Head.ContentType = c
	}

	d.Head.Write(w)
	w.Write(data.([]byte))
	return nil
}

// Render an HTML response using Pongo2.
func (h HTML) Render(w http.ResponseWriter, data interface{}) error {
	h.Head.Write(w)
	err := h.Template.ExecuteWriter(data.(map[string]interface{}), w)
	return err
}

// Render a JSON response.
func (j JSON) Render(w http.ResponseWriter, data interface{}) error {
	var result []byte
	var err error

	if j.Indent {
		result, err = json.MarshalIndent(data, "", "  ")
		result = append(result, '\n')
	} else {
		result, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}

	// JSON marshaled fine, write out the result.
	j.Head.Write(w)
	if len(j.Prefix) > 0 {
		w.Write(j.Prefix)
	}
	w.Write(result)
	return nil
}

// Render a JSONP response.
func (j JSONP) Render(w http.ResponseWriter, data interface{}) error {
	var result []byte
	var err error

	if j.Indent {
		result, err = json.MarshalIndent(data, "", "  ")
	} else {
		result, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}

	// JSON marshaled fine, write out the result.
	j.Head.Write(w)
	w.Write([]byte(j.Callback + "("))
	w.Write(result)
	w.Write([]byte(");"))

	// If indenting, append a new line.
	if j.Indent {
		w.Write([]byte("\n"))
	}
	return nil
}

// Render an XML response.
func (x XML) Render(w http.ResponseWriter, data interface{}) error {
	var result []byte
	var err error

	if x.Indent {
		result, err = xml.MarshalIndent(data, "", "  ")
		result = append(result, '\n')
	} else {
		result, err = xml.Marshal(data)
	}
	if err != nil {
		return err
	}

	// XML marshaled fine, write out the result.
	x.Head.Write(w)
	if len(x.Prefix) > 0 {
		w.Write(x.Prefix)
	}
	w.Write(result)
	return nil
}

// Render a Markdown response as HTML.
func (m Markdown) Render(w http.ResponseWriter, data interface{}) error {
	m.Head.Write(w)
	output := blackfriday.MarkdownCommon(data.([]byte))
	w.Write(output)
	return nil
}
