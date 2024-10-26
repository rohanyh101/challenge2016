package utils

import "strings"

func ValidateIncludeExclude(locations []string) bool {
	for _, l := range locations {
		location := strings.Split(l, ",")

		if len(location) == 3 {
			_, foundCity := CityMap[location[0]]
			_, foundProvince := ProvinceMap[location[1]]
			_, foundCountry := CountryMap[location[2]]

			if !foundCity || !foundProvince || !foundCountry {
				return false
			}
		} else if len(location) == 2 {
			_, foundProvince := ProvinceMap[location[0]]
			_, foundCountry := CountryMap[location[1]]

			if !foundProvince || !foundCountry {
				return false
			}
		} else if len(location) == 1 {
			_, foundCountry := CountryMap[location[0]]

			if !foundCountry {
				return false
			}
		}
	}

	return true
}
