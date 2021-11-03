package repository

import (
	"database/sql"
	"strconv"
)

type Storage struct {
	db *sql.DB
}

type Country struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type Region struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type City struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	AreaTitle   string `json:"area_title,omitempty"`
	RegionTitle string `json:"region_title,omitempty"`
}

func (s Storage) GetCountry() ([]Country, error) {

	rows, err := s.db.Query("select id, title from country")
	if err != nil {
		return []Country{}, err
	}

	var countries []Country

	for rows.Next() {
		c := Country{}
		err := rows.Scan(&c.Id, &c.Title)
		if err != nil {
			return []Country{}, err
		}

		countries = append(countries, c)
	}

	return countries, nil
}

func (s Storage) GetRegions(countryID int) ([]Region, error) {

	rows, err := s.db.Query("select id, title from regions where country_id=?", countryID)
	if err != nil {
		return []Region{}, err
	}

	var regions []Region

	for rows.Next() {
		region := Region{}
		err := rows.Scan(&region.Id, &region.Title)
		if err != nil {
			return []Region{}, err
		}
		regions = append(regions, region)
	}

	return regions, nil
}

func (s Storage) GetCities(regionID int) ([]City, error) {
	reg := strconv.Itoa(regionID)

	rows, err := s.db.Query("select id, title, area_title, region_title from cities where region_id=?", reg)
	if err != nil {
		return []City{}, err
	}

	return s.listCities(rows)
}

func (s Storage) Search(regionID int, search string) ([]City, error) {
	reg := strconv.Itoa(regionID)

	rows, err := s.db.Query("select id, title, area_title, region_title from cities where region_id=? and title like ?", reg, search+"%")

	if err != nil {
		return []City{}, err
	}

	return s.listCities(rows)
}

func (s Storage) listCities(rows *sql.Rows) ([]City, error) {
	var cities []City

	for rows.Next() {
		city := City{}
		err := rows.Scan(&city.Id, &city.Title, &city.AreaTitle, &city.RegionTitle)
		if err != nil {
			return []City{}, err
		}
		cities = append(cities, city)
	}

	return cities, nil
}

func (s *Storage) SetDb(db *sql.DB) {
	s.db = db
}

func NewStorage(db *sql.DB) Storage {
	return Storage{db: db}
}
