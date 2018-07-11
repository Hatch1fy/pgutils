package pgutils

import (
	"github.com/BurntSushi/toml"
)

// NewConfig will return a new parsed postgres configuration from a toml source
func NewConfig(src string) (c Config, err error) {
	_, err = toml.DecodeFile(src, &c)
	return
}

// Config is the configuration for a postgres database
type Config struct {
	Host string `toml:"host"`
	Port uint16 `toml:"port"`

	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`

	SSL bool `toml:"ssl"`
}
