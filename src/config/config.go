package config

import (
	"os"
	"strconv"
)

const defaultPort = 1
const defaultDSN = "user@/location"

type Config struct {
	DefaultPort int
	DSN         string
}

func (c *Config) Load() {
	var err error

	port, err := strconv.Atoi(os.Getenv("DEFAULT_PORT"))
	if err != nil || port == 0 {
		port = defaultPort
	}

	c.DefaultPort = port
	if err != nil {
		return
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = defaultDSN
	}

	c.DSN = dsn
}

func (c Config) GetPort() int {
	return c.DefaultPort
}

func (c Config) GetDsn() string {
	return c.DSN
}
