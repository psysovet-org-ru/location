package migrations

import (
	"bufio"
	"database/sql"
	"github.com/iamsalnikov/mymigrate"
	"location/service"
	"log"
	"os"
)

func init() {
	mymigrate.Add(
		"mig_load_data",
		func(db *sql.DB) error {
			dataSQL := service.GetData()

			for _, fileName := range dataSQL {
				f, err := os.Open(fileName)
				if err != nil {
					return err
				}

				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					q := scanner.Text()
					_, err := db.Exec(q)
					if err != nil {
						log.Println(q)
						return err
					}
				}

				if err := scanner.Err(); err != nil {
					return err
				}
			}

			return nil
		},
		func(db *sql.DB) error {
			_, err := db.Exec("delete from cities")
			if err != nil {
				return err
			}
			_, err = db.Exec("delete from regions")
			if err != nil {
				return err
			}

			_, err = db.Exec("delete from country")
			if err != nil {
				return err
			}

			return nil
		},
	)

}
