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

	envs, err := godotenv.Read()
	if err != nil {
		return
	}

	port := os.Getenv("DEFAULT_PORT")

	var ok bool

	if port == "" {
		port, ok = envs["DEFAULT_PORT"]
		if !ok {
			return
		}
	}
	c.DefaultPort, err = strconv.Atoi(os.Getenv("DEFAULT_PORT"))
	if err != nil {
		return
	}

	dsn := os.Getenv("DSN")

	if dsn == "" {
		dsn, ok = envs["DSN"]
		if !ok {
			return
		}
	}

	c.DSN = dsn
}

func (c Config) GetPort() int {
	return c.DefaultPort
}

func (c Config) GetDsn() string {
	return c.DSN
}
