package connect

import (
	"database/sql"
	"os"
)

type Connect struct {
	db *sql.DB
}

func (c *Connect) new(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	return db, err
}

func (c *Connect) Get() (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	if c.db == nil {
		db, err := c.new(dsn)
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
