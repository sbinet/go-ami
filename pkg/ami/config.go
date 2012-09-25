package ami

import (
	"github.com/sbinet/go-config/config"
)

type Config struct {
	config.Config
}

func NewConfig() Config {
	return Config{}
}

// EOF
