package mojito

// Context is the base request context applications should use.
// Custom applications can extend this context and add new fields.
type Context struct {
	Config *Config
}
