package mojito

import (
	"log"
	"net/http"
)

// Context is the base request context for all Mojito applications.
type Context struct {
	ResponseWriter *http.ResponseWriter
	Request        *http.Request
	Logger         *log.Logger
}
