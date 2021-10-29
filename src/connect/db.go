package connect

import (
	"database/sql"
	"location/config"
)

type Connect struct {
	db     *sql.DB
	config config.Config
}

func (c *Connect) SetConfig(cfg config.Config) {
	c.config = cfg
}

func (c *Connect) new(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	return db, err
}

func (c *Connect) Get() (*sql.DB, error) {
	if c.db == nil {
		db, err := c.new(c.config.GetDsn())

		if err != nil {
			return nil, err
		}
		c.Set(db)
	}

	return c.db, nil
}

func (c *Connect) Set(db *sql.DB) {
	c.db = db
}
