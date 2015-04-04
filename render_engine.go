package mojito

import "net/http"

// Head defines the basic ContentType and Status fields.
type Head struct {
	ContentType string
	Status      int
}

// Write outputs the header content.
func (h Head) Write(w http.ResponseWriter) {
	w.Header().Set(ContentType, h.ContentType)
	w.WriteHeader(h.Status)
}
