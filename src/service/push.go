package service

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"
)

func PushData(db *sql.DB) error {
	err := pushC(db)
	if err != nil {
		return err
	}

	err = pushR(db)
	if err != nil {
		return err
	}

	err = pushCI(db)
	if err != nil {

		return err
	}

	return nil
}

func pushC(db *sql.DB) error {
	records, err := readData("./data/countries.csv")

	if err != nil {
		return err
	}

	for _, record := range records {
		t, _ := time.Parse("2006-01-02 15:04:05", record[2])
		_, err := db.Exec("insert into country (id,title,update_at) values (?,?,?);", record[0], record[1], t.Format("2006-01-02 15:04:05"))
		if err != nil {
			return err
		}
	}

	return nil
}

func pushR(db *sql.DB) error {
	records, err := readData("./data/regions.csv")

	if err != nil {
		return err
	}

	for _, record := range records {
		t, _ := time.Parse("2006-01-02 15:04:05", record[3])
		_, err := db.Exec("insert into regions (id,country_id,title,update_at) values (?,?,?,?);", record[0], record[1], record[2], t.Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Fatalln(record)
			return err
		}
	}

	return nil
}

func pushCI(db *sql.DB) error {
	records, err := readData("./data/cities.csv")

	if err != nil {
		return err
	}

	for _, record := range records {
		t, _ := time.Parse("2006-01-02 15:04:05", record[5])
		_, err := db.Exec("insert into cities (id,region_id,title,region_title,area_title,update_at) values (?,?,?,?,?,?);", record[0], record[1], record[2], record[3], record[4], t.Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Fatalln(implode(record))
			return err
		}
	}

	return nil
}

func implode(data []string) string {
	r := ""
	for _, item := range data {
		r += item
	}
	return r
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	var records [][]string
	numErr := 0

	for {
		record, e := r.Read()

		if err == io.EOF {
			break
		}
		if numErr > 3 {
			break
		}

		if e != nil {
			numErr++
			continue
		}

		records = append(records, record)
	}

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
