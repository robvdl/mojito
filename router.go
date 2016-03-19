package mojito

import (
	"log"
	"reflect"
)

// Router is used for the root router and sub routers.
type Router struct {
	parent     *Router // parent Router, nil if this is the root Router.
	children   []*Router
	pathPrefix string

	// stores a pointer to the logger to use for this application.
	Logger *log.Logger

	// stores a pointer to the application config struct.
	Config *Config

	// one of these custom Context structs will be created per request.
	contextType reflect.Type
}
