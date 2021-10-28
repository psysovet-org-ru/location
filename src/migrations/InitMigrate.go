package migrations

import (
	"database/sql"

	"github.com/iamsalnikov/mymigrate"
)

func init() {
	mymigrate.Add(
		"mig_init",
		func(db *sql.DB) error {

			_, err := db.Exec("create table country\n(\n\tid int auto_increment,\n\ttitle varchar(255) null,\n\tupdate_at datetime null,\n\tconstraint country_pk\n\t\tprimary key (id)\n);")
			if err != nil {
				return err
			}

			_, err = db.Exec("create table regions\n(\n\tid int auto_increment,\n\tcountry_id int null,\n\ttitle varchar(255) null,\n\tupdate_at datetime null,\n\tconstraint regions_pk\n\t\tprimary key (id),\n    constraint regions_country_id_fk\n        foreign key (country_id) references country (id)\n            on update cascade on delete cascade\n);")
			if err != nil {
				return err
			}

			_, err = db.Exec("create table cities\n(\n\tid int auto_increment,\n\tregion_id int null,\n\ttitle varchar(255) null,\n\tarea_title varchar(255) null,\n\tregion_title varchar(255) null,\n    update_at datetime null,\n\tconstraint cities_pk\n\t\tprimary key (id),\n    constraint cities_regions_id_fk\n        foreign key (region_id) references regions (id)\n            on update cascade on delete cascade\n);")
			if err != nil {
				return err
			}
			return nil
		},
		func(db *sql.DB) error {

			_, err := db.Exec("drop table cities")
			if err != nil {
				return err
			}
			_, err = db.Exec("drop table regions")
			if err != nil {
				return err
			}

			_, err = db.Exec("drop table country")
			if err != nil {
				return err
			}

			return nil
		},
	)

}
