package mojito

import (
	"log"
	"os"
	"strings"
)

// GetLogger creates a logger from the config file.
// This is likely to be developed further a bit later, for now this will do.
func GetLogger(config *Config) *log.Logger {
	handler := strings.ToLower(config.GetString("logger.handler"))
	prefix := config.GetString("logger.prefix")
	flag := config.GetInt("logger.flag")

	// for now this is the only supported handler.
	if handler == "stdout" || handler == "os.stdout" {
		return log.New(os.Stdout, prefix, flag)
	}

	// fallback to stdout if config is wrong.
	return log.New(os.Stdout, "", 0)
}
