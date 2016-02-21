package mojito

import "net/http"

// Run simply runs http.ListenAndServe.
func (r *Router) Run(addr string) {
	http.ListenAndServe(addr, r)
}
