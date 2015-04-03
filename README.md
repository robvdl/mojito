# Mojito Web Framework [![GoDoc](https://godoc.org/github.com/robvdl/mojito?status.svg)](http://godoc.org/github.com/robvdl/mojito)

Mojito is yet another micro web framework for Go, similar to Martini and
Negroni. It was originally based on the Negroni codebase, making the
middleware stack compatible with Negroni.

It is slightly more opinionated however, by choosing a particular router
(Gorilla Mux) and template library (Pongo2), I wanted to make it even easier
to get started and write apps with Mojito.  You don't need to bring in other
libraries like unrolled/render in order to render template responses for example
and then add another library on top of that to handle Pongo templates, this is
something I felt should be built into the Mojito framework itself to make
things easier and reduce the number of dependencies required.

Unlike Negroni, that DOES make Mojito a framework, Negroni is mostly just
a middleware stack but still requires other components to make it a framework.

WARNING: at the moment, Mojito is still in experimental stages, don't expect
things to be working just yet.  I am not entirely sure whether forking
Negroni is the best idea, but we'll see where it takes us.

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH),
create your first `.go` file. We'll call it `server.go`.

```go
package main

import (
    "net/http"

    "github.com/robvdl/mojito"
)

func main() {
    m := mojito.Classic()

    m.Get("/", func(c *mojito.Context) {
        c.HTML(http.StatusOK, "index", map[string]string{"hello": "mojito"})
    })

    m.Get("/json", func(c *mojito.Context) {
        c.JSON(http.StatusOK, map[string]string{"hello": "mojito"})
    })

    m.Run(":3000")
}
```
