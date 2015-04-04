package mojito

import "net/http"

// ViewFunc is a view handler that takes a mojito.Context
type ViewFunc func(c *Context)

// Get registers a GET request with the router
func (m *Mojito) Get(path string, handler ViewFunc) {
	m.router.HandleFunc(path, func(rw http.ResponseWriter, req *http.Request) {
		handler(&Context{
			Options: m.options,
			Request: req,
			Writer:  rw,
		})
	})
}
