package mojito

import "net/http"

// Handler handler is an interface that objects can implement to be registered
// to serve as middleware in the Mojito middleware stack.
//
// ServeHTTP should yield to the next middleware in the chain by invoking the
// next http.HandlerFunc passed in.
//
// If the Handler writes to the ResponseWriter, the next http.HandlerFunc should
// not be invoked.
type Handler interface {
	ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc)
}

// HandlerFunc is an adapter to allow the use of ordinary functions as Mojito
// handlers. If f is a function with the appropriate signature, HandlerFunc(f)
// is a Handler object that calls f.
type HandlerFunc func(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc)

type middleware struct {
	handler Handler
	next    *middleware
}

func (h HandlerFunc) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	h(rw, req, next)
}

func (m middleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	m.handler.ServeHTTP(rw, req, m.next.ServeHTTP)
}

// Wrap converts a http.Handler into a mojito.Handler so it can be used as a
// Mojito middleware. The next http.HandlerFunc is automatically called after
// the Handler is executed.
func Wrap(handler http.Handler) Handler {
	return HandlerFunc(func(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		handler.ServeHTTP(rw, req)
		next(rw, req)
	})
}

func build(handlers []Handler) middleware {
	var next middleware

	if len(handlers) == 0 {
		return voidMiddleware()
	} else if len(handlers) > 1 {
		next = build(handlers[1:])
	} else {
		next = voidMiddleware()
	}

	return middleware{handlers[0], &next}
}

func voidMiddleware() middleware {
	return middleware{
		HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}),
		&middleware{},
	}
}
