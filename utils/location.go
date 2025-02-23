package utils

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/roh4nyh/qube_challenge_2016/models"
)

var (
	CityMap     map[string]models.Location
	ProvinceMap map[string]models.Location
	CountryMap  map[string]models.Location
)

func LoadCities() ([]models.Location, error) {
	var locations []models.Location

	f, err := os.Open("data/cities.csv")
	if err != nil {
		err := fmt.Errorf("error opening cities.csv file: %v", err)
		return nil, err
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		err := fmt.Errorf("error reading cities.csv file: %v", err)
		return nil, err
	}

	CityMap = make(map[string]models.Location)
	ProvinceMap = make(map[string]models.Location)
	CountryMap = make(map[string]models.Location)

	for i, record := range records {
		if i == 0 {
			continue
		}

		location := models.Location{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			CityName:     record[3],
			ProvinceName: record[4],
			CountryName:  record[5],
		}

		CityMap[location.CityCode] = location
		ProvinceMap[location.ProvinceCode] = location
		CountryMap[location.CountryCode] = location
		locations = append(locations, location)
	}

	return locations, nil
}

func RemoveDuplicateLocations(locations []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, location := range locations {
		if _, value := keys[location]; !value {
			keys[location] = true
			list = append(list, location)
		}
	}

	return list
}
