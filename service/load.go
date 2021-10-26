package service

import (
	"database/sql"
	"encoding/csv"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"time"
)

type country struct {
	id        int       `json:"id,omitempty"`
	title     string    `json:"title,omitempty"`
	sourceId  int       `json:"source_id,omitempty"`
	updatedAt time.Time `json:"updated_at,omitempty"`
}

type region struct {
	id        int       `json:"id,omitempty"`
	title     string    `json:"title,omitempty"`
	sourceId  int       `json:"source_id,omitempty"`
	countryId int       `json:"country_id,omitempty"`
	updatedAt time.Time `json:"updated_at,omitempty"`
}

type city struct {
	id        int       `json:"id,omitempty"`
	sourceId    int    `json:"source_id,omitempty"`
	regionId    int    `json:"region_id,omitempty"`
	title        string `json:"title,omitempty"`
	areaTitle   string `json:"area_title,omitempty"`
	regionTitle string `json:"region_title,omitempty"`
	updatedAt time.Time `json:"updated_at,omitempty"`
}

func LoadData() error {
	connStr := "user=psy password=psy dbname=psy sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return err
	}

	err = loadC(db)

	if err != nil {
		return err
	}

	err = loadR(db)

	if err != nil {
		return err
	}

	err = loadCI(db)

	if err != nil {
		return err
	}

	return nil
}

func loadC(db *sql.DB) error {
	rows, err := db.Query("select id, source_id, title, updated_at from countries")

	if err != nil {
		return err
	}

	defer rows.Close()

	f, err := os.Create("countries.csv")
	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for rows.Next() {
		c := country{}
		err := rows.Scan(&c.id, &c.sourceId, &c.title, &c.updatedAt)
		if err != nil {
			return err
		}
		//	fmt.Println(c.title, "\n")

		var record = []string{
			strconv.Itoa(c.sourceId),
			c.title,
			c.updatedAt.String(),
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func loadR(db *sql.DB) error {
	rows, err := db.Query("select id, source_id, source_country_id, title, updated_at from regions")

	if err != nil {
		return err
	}

	defer rows.Close()

	f, err := os.Create("regions.csv")
	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for rows.Next() {
		r := region{}
		err := rows.Scan(&r.id, &r.sourceId, &r.countryId, &r.title, &r.updatedAt)
		if err != nil {
			return err
		}
		//fmt.Println(r.title, "\n")

		var record = []string{
			strconv.Itoa(r.sourceId),
			strconv.Itoa(r.countryId),
			r.title,
			r.updatedAt.String(),
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func loadCI(db *sql.DB) error {
	rows, err := db.Query("select id, source_id, source_region_id, title,region_title,area_title, updated_at from cities")

	if err != nil {
		return err
	}

	defer rows.Close()

	f, err := os.Create("cities.csv")
	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for rows.Next() {
		c := city{}
		err := rows.Scan(&c.id, &c.sourceId, &c.regionId, &c.title, &c.regionTitle, &c.areaTitle, &c.updatedAt)
		if err != nil {
			return err
		}
		//fmt.Println(r.title, "\n")

		var record = []string{
			strconv.Itoa(c.sourceId),
			strconv.Itoa(c.regionId),
			c.title,
			c.regionTitle,
			c.areaTitle,
			c.updatedAt.String(),
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}
