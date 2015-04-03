package mojito

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Options is a struct for specifying configuration options for the
// mojito.Mojito application object.
type Options struct {
	TemplateDirectory string
	IsDevelopment     bool
}

// Mojito is a stack of Middleware Handlers that can be invoked as an
// http.Handler. Mojito middleware are evaluated in the order that they are
// added to the stack using the Use and UseHandler methods.
type Mojito struct {
	middleware middleware
	handlers   []Handler
	options    *Options
	router     *mux.Router
}

// New returns a new Mojito instance with no middleware preconfigured.
func New(opt *Options, handlers ...Handler) *Mojito {
	m := Mojito{
		handlers:   handlers,
		middleware: build(handlers),
		options:    opt,
		router:     mux.NewRouter(),
	}
	return &m
}

// Classic returns a new Mojito instance with the default middleware already
// in the stack and default Options.
//
// Recovery - Panic Recovery Middleware
// Logger - Request/Response Logging
// Static - Static File Serving
func Classic() *Mojito {
	options := Options{
		TemplateDirectory: "templates",
		IsDevelopment:     false,
	}
	return New(&options, NewRecovery(), NewLogger(), NewStatic(http.Dir("public")))
}

func (m *Mojito) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	m.middleware.ServeHTTP(NewResponseWriter(rw), req)
}

// Use adds a Handler onto the middleware stack. Handlers are invoked in the
// order they are added to a Mojito.
func (m *Mojito) Use(handler Handler) {
	m.handlers = append(m.handlers, handler)
	m.middleware = build(m.handlers)
}

// UseFunc adds a Mojito-style handler function onto the middleware stack.
func (m *Mojito) UseFunc(handlerFunc func(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc)) {
	m.Use(HandlerFunc(handlerFunc))
}

// UseHandler adds a http.Handler onto the middleware stack. Handlers are
// invoked in the order they are added to a Mojito.
func (m *Mojito) UseHandler(handler http.Handler) {
	m.Use(Wrap(handler))
}

// UseHandlerFunc adds a http.HandlerFunc-style handler function onto the
// middleware stack.
func (m *Mojito) UseHandlerFunc(handlerFunc func(rw http.ResponseWriter, req *http.Request)) {
	m.UseHandler(http.HandlerFunc(handlerFunc))
}

// Run is a convenience function that runs the mojito stack as an HTTP
// server. The addr string takes the same format as http.ListenAndServe.
func (m *Mojito) Run(addr string) {
	// The router must be connected last, so connect it here
	m.UseHandler(m.router)
	l := log.New(os.Stdout, "[mojito] ", 0)
	l.Printf("listening on %s", addr)
	l.Fatal(http.ListenAndServe(addr, m))
}

// Handlers returns a list of all the handlers in the current Mojito
// middleware chain.
func (m *Mojito) Handlers() []Handler {
	return m.handlers
}
