package service

import (
	"archive/zip"
	"io"
	"location/config"
	"net/http"
	"os"
)

func download(URLFile, fileName string) (int64, error) {

	resp, err := http.Get(URLFile)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	f, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	size, err := io.Copy(f, resp.Body)
	return size, nil
}

func unzip(fileName, unzipPath string) ([]string, error) {
	// Распаковка содержимого архива
	zipR, err := zip.OpenReader(fileName)
	if err != nil {
		return nil, err
	}

	var result []string

	for _, file := range zipR.File {
		r, err := file.Open()
		if err != nil {
			return nil, err
		}
		unzipFile, err := os.Create(unzipPath + file.Name)

		_, err = io.Copy(unzipFile, r)
		if err != nil {
			return nil, err
		}
		err = r.Close()
		if err != nil {
			return nil, err
		}

		result = append(result, unzipFile.Name())
	}

	return result, nil
}

func GetData() []string {
	cfg := config.Config{}
	cfg.Load()

	fileName := cfg.GetDownloadPath() + "location_data.zip"
	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}

	size, err := download(cfg.GetDownloadData(), fileName)
	if err != nil || size == 0 {
		return nil
	}

	_, err = unzip(fileName, cfg.GetDownloadPath())
	if err != nil {
		return nil
	}

	os.Remove(fileName)

	return []string{cfg.DownloadPath + "location_country.sql", cfg.DownloadPath + "location_regions.sql", cfg.DownloadPath + "location_cities.sql"}
}
