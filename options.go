package mojito

// Options is a struct for specifying configuration options for the
// mojito.Mojito application object. It mostly contains render options.
// Note that defautlt options are only applied with a mojito.Classic().
type Options struct {
	// Directory to load templates from. Default is "templates".
	TemplateDirectory string
	// Appends the character set to the Content-Type header. Default is "UTF-8".
	Charset string
	// Outputs human readable JSON if true.
	IndentJSON bool
	// Outputs human readable XML if true.
	IndentXML bool
	// Prefixes the JSON output with the given bytes.
	PrefixJSON []byte
	// Prefixes the XML output with the given bytes.
	PrefixXML []byte
	// Allows changing of output to XHTML instead of HTML. Default is "text/html"
	HTMLContentType string
	// If true, recompile the templates on every request. Default is false.
	IsDevelopment bool
}
