package service

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"location/repository"
	"log"
	"net/http"
	"strconv"
)

func Service(r *chi.Mux, s repository.Storage) *chi.Mux {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/countries/", func(writer http.ResponseWriter, request *http.Request) {

		countries, err := s.GetCountry()
		if err != nil {
			return
		}

		jsonData, err := json.Marshal(countries)

		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
		}

		writer.Write(jsonData)
	})

	r.Get("/regions/{country}/", func(writer http.ResponseWriter, request *http.Request) {
		c := chi.URLParam(request, "country")
		country, err := strconv.Atoi(c)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
		}

		regions, err := s.GetRegions(country)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
		}

		jsonData, err := json.Marshal(regions)

		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
		}

		writer.Write(jsonData)
	})

	r.Get("/cities/{region}/", func(writer http.ResponseWriter, request *http.Request) {
		reg := chi.URLParam(request, "region")
		region, err := strconv.Atoi(reg)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}

		cities, err := s.GetCities(region)

		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
		}

		jsonData, err := json.Marshal(cities)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}

		writer.Write(jsonData)
	})

	r.Get("/search/city/{region}/{search}", func(writer http.ResponseWriter, request *http.Request) {
		reg := chi.URLParam(request, "region")
		search := chi.URLParam(request, "search")
		region, err := strconv.Atoi(reg)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}

		searchData, err := s.Search(region, search)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}

		jsonData, err := json.Marshal(searchData)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}

		writer.Write(jsonData)
	})

	return r
}
