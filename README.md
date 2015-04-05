# Mojito [![GoDoc](https://godoc.org/github.com/robvdl/mojito?status.svg)](http://godoc.org/github.com/robvdl/mojito) [![Build Status](https://travis-ci.org/robvdl/mojito.svg?branch=master)](https://travis-ci.org/robvdl/mojito) [![Coverage Status](https://img.shields.io/coveralls/robvdl/mojito.svg)](https://coveralls.io/r/robvdl/mojito)

Mojito is yet another micro web framework for Go, similar to Martini and
Negroni. It was originally based on the Negroni codebase, making the
middleware stack compatible with Negroni.

It is slightly more opinionated however, by choosing a particular router
(Gorilla Mux) and template library (Pongo2), I wanted to make it even easier
to get started and write apps with Mojito.  You don't need to bring in other
libraries like unrolled/render in order to render template responses for example
and then add another library on top of that to handle Pongo templates, this is
something I felt should be built into the Mojito framework itself.

Unlike Negroni, that DOES make Mojito a framework, Negroni is mostly just
a middleware stack but still requires other components to make it a framework.

WARNING: at the moment, Mojito is still in experimental stages, don't expect
things to be working just yet.  I am not entirely sure whether forking
Negroni is the best idea, but we'll see where it takes us.

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH),
create your first `.go` file. We'll call it `server.go`.

The following example has two view functions, one that renders a Pongo2
template and the other returns some JSON.

```go
package main

import (
    "encoding/xml"
    "io/ioutil"
    "net/http"

    "github.com/robvdl/mojito"
)

type exampleXML struct {
    XMLName xml.Name `xml:"example"`
    One     string   `xml:"one,attr"`
    Two     string   `xml:"two,attr"`
}

func main() {
    m := mojito.Classic()

    m.Get("/", func(c *mojito.Context) {
        c.HTML(http.StatusOK, "index.html", map[string]interface{}{"hello": "html"})
    })

    m.Get("/data", func(c *mojito.Context) {
        c.Data(http.StatusOK, []byte("Some binary data here."))
    })

    m.Get("/json", func(c *mojito.Context) {
        c.JSON(http.StatusOK, map[string]string{"hello": "json"})
    })

    m.Get("/jsonp", func(c *mojito.Context) {
        c.JSONP(http.StatusOK, "callbackName", map[string]string{"hello": "jsonp"})
    })

    m.Get("/xml", func(c *mojito.Context) {
        c.XML(http.StatusOK, exampleXML{One: "hello", Two: "xml"})
    })

    m.Get("/markdown", func(c *mojito.Context) {
        if markdown, err := ioutil.ReadFile("README.md"); err == nil {
            c.Markdown(http.StatusOK, markdown)
        } else {
            http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
        }
    })

    m.Run(":3000")
}
```

Because we are using mojito.Classic() to construct our application, which has
default settings for the templates and static files directories, you should
create the two directories "./templates" and "./public" now.

The above example expects a Pongo2 template "templates/index.html" to exist.
