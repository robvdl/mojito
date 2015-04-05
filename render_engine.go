package mojito

import (
	"encoding/json"
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

// JSON built-in renderer.
type JSON struct {
	Head
	Indent bool
	Prefix []byte
}

// Write outputs the header content.
func (h *Head) Write(w http.ResponseWriter) {
	w.Header().Set(ContentType, h.ContentType)
	w.WriteHeader(h.Status)
}

// Render an HTML response.
func (h *HTML) Render(w http.ResponseWriter, data map[string]interface{}) error {
	h.Head.Write(w)
	err := h.Template.ExecuteWriter(data, w)
	return err
}

// Render a JSON response.
func (j *JSON) Render(w http.ResponseWriter, data interface{}) error {
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
