package mojito

import "github.com/pelletier/go-toml"

// Config stores the application configuration.
type Config struct {
	data *toml.TomlTree
}

// LoadConfig loads a configuration file from disk.
func LoadConfig(filename string) (*Config, error) {
	config, err := toml.LoadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Config{data: config}, nil
}

// Get retrieves the config value and returns an untyped interface{}.
func (c *Config) Get(key string) interface{} {
	return c.data.Get(key)
}

// GetString retrieves the config value and casts it to a string.
func (c *Config) GetString(key string) string {
	return c.data.Get(key).(string)
}

// GetInt retrieves the config value and casts it to an int.
func (c *Config) GetInt(key string) int {
	return int(c.data.Get(key).(int64))
}

// GetInt64 retrieves the config value and casts it to an int64.
func (c *Config) GetInt64(key string) int64 {
	return c.data.Get(key).(int64)
}

// GetBool retrieves the config value and casts it to a bool.
func (c *Config) GetBool(key string) bool {
	return c.data.Get(key).(bool)
}

// GetFloat retrieves the config value and casts it to a float64.
func (c *Config) GetFloat(key string) float64 {
	return c.data.Get(key).(float64)
}
