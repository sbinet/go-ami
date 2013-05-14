package ami

import (
	"github.com/gonuts/config"
)

type Config struct {
	config.Config
}

func NewConfig() Config {
	return Config{}
}

// EOF
