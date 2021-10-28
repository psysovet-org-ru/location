package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	DefaultPort int
	DSN         string
}

func (c *Config) Load() {
	var err error
	godotenv.Load()
	c.DefaultPort, err = strconv.Atoi(os.Getenv("DEFAULT_PORT"))
	if err != nil {
		c.DefaultPort = 3000
	}

	c.DSN = os.Getenv("DSN")
}

func (c Config) GetPort() int {
	return c.DefaultPort
}

func (c Config) GetDsn() string {
	return c.DSN
}
