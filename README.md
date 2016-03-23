Mojito web framework for Go
===========================

Why the name?
-------------

Originally Mojito was based on a forked version of the Negroni "framework",
with some additions like a built-in renderer. I had also used the Gin framework
around that time, and it seemed it was custom in Go to name your framework
after a cocktail or other alcoholic drink so the name Mojito had stuck.

Some time later Mojito was rewritten from ground up, with new inspiration from
other frameworks, particularly gocraft/web with it's custom request contexts.
But I wanted something that felt more like a "cohesive" framework where
every component works well together, with authentication and session
middlewares already built in, and a base context provided by the framework
that contains things like a User struct for the current logged in user. This
eventually got me to rewrite Mojito from ground up as a brand new framework.

Sample application
------------------

```go
package main

import "github.com/robvdl/mojito"

// Context is your custom request context, it must embed *mojito.Context
type Context struct {
  *mojito.Context
}

// Configure gets called automatically by mojito.New().
func (c Context) Configure() *mojito.Config {
	config, err := mojito.LoadConfig("config.toml")
	if err != nil {
		fmt.Println("Failed to load config file:", err)
		os.Exit(1)
	}
	return config
}

// Home is a simple route based on your own Context.
func (c *Context) Home() {
  c.Logger.Println("Home route")
  c.HTML(http.StatusOK, "index.html", map[string]interface{}{"hello": "world"})
}

func main() {
  m := mojito.New(Context{})
  m.Get("/", (*Context).Home)
  m.Run()
}
```

Built in support for config files
---------------------------------

Support for loading config files is built-in. At the moment only TOML
config files are supported, but support for other config file formats
could be added in the future by creating an interface.

The built-in config module was modelled somewhat around the Viper library,
however Viper can only load one set of config settings at once, as Viper
uses a global configuration, while we store each configuration in a struct
so you can load multiple configurations, if this is what your app requires.

```go
config := mojito.LoadConfig("config.toml")
host := config.GetString("server.host")
port := config.GetString("server.port")
```

Your base context must implement the mojito.Application interface,
this means implementing the Configure() method (see example application).

Some other configuration libraries might map your configuration file
into a struct as it loads it, this seems nice at first, however it makes
it difficult to define a base config struct in Mojito, then get the user
to optionally extend it with user defined settings.  In the end, the dotted
config key approach seems to actually work best.

Custom request contexts
-----------------------

Router and sub-routers
----------------------

Middleware
----------

Sessions and authentication
---------------------------
