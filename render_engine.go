package mojito

import (
	"net/http"

	"github.com/flosch/pongo2"
)

// Head defines the basic ContentType and Status fields.
type Head struct {
	ContentType string
	Status      int
}

// HTML built-in template renderer.
type HTML struct {
	Head
	Name     string
	Template *pongo2.Template
}

// Write outputs the header content.
func (h *Head) Write(w http.ResponseWriter) {
	w.Header().Set(ContentType, h.ContentType)
	w.WriteHeader(h.Status)
}

// Render renders the HTML template to the response.
func (h *HTML) Render(w http.ResponseWriter, data map[string]interface{}) error {
	return h.Template.ExecuteWriter(data, w)
}
