package config

import (
	"os"
	"strconv"
)

const defaultPort = 1
const defaultDSN = "user@/location"
const defaultPath = "./"

const KeyDbDsn = "DB_DSN"
const KeyPort = "DEFAULT_PORT"
const KeyDownloadData = "DOWNLOAD_DATA"
const KeyDownloadPath = "DOWNLOAD_PATH"

type Config struct {
	DefaultPort  int
	DbDSN        string
	DownloadData string
	DownloadPath string
}

func (c *Config) Load() {
	var err error
	port, err := strconv.Atoi(os.Getenv(KeyPort))
	if err != nil || port == 0 {
		port = defaultPort
	}

	c.DefaultPort = port
	if err != nil {
		return
	}

	dsn := os.Getenv(KeyDbDsn)
	if dsn == "" {
		dsn = defaultDSN
	}

	c.DbDSN = dsn

	downloadData := os.Getenv(KeyDownloadData)
	if downloadData == "" {
		downloadData = ""
	}

	c.DownloadData = downloadData

	downloadPath := os.Getenv(KeyDownloadPath)
	if downloadPath == "" {
		downloadPath = defaultPath
	}

	c.DownloadPath = downloadPath
}

func (c Config) GetPort() int {
	return c.DefaultPort
}

func (c Config) GetDsn() string {
	return c.DbDSN
}

func (c Config) GetDownloadData() string {
	return c.DownloadData
}

func (c *Config) GetDownloadPath() string {
	return c.DownloadPath
}
