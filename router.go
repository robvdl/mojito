package mojito

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ViewFunc is a view handler that takes a mojito.Context
type ViewFunc func(c *Context)

// Handle registers a generic request with the router
func (m *Mojito) Handle(path string, handler ViewFunc) *mux.Route {
	return m.router.HandleFunc(path, func(rw http.ResponseWriter, req *http.Request) {
		handler(&Context{
			Options: m.options,
			Request: req,
			Writer:  rw,
		})
	})
}

// Get registers a GET request with the router
func (m *Mojito) Get(path string, handler ViewFunc) *mux.Route {
	return m.Handle(path, handler).Methods("GET")
}

// Post registers a POST request with the router
func (m *Mojito) Post(path string, handler ViewFunc) *mux.Route {
	return m.Handle(path, handler).Methods("POST")
}

// Put registers a PUT request with the router
func (m *Mojito) Put(path string, handler ViewFunc) *mux.Route {
	return m.Handle(path, handler).Methods("PUT")
}

// Patch registers a PATCH request with the router
func (m *Mojito) Patch(path string, handler ViewFunc) *mux.Route {
	return m.Handle(path, handler).Methods("PATCH")
}

// Delete registers a DELETE request with the router
func (m *Mojito) Delete(path string, handler ViewFunc) *mux.Route {
	return m.Handle(path, handler).Methods("DELETE")
}
