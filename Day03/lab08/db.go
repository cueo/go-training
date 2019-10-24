package lab08

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var CountriesDB map[string][]City = make(map[string][]City)

func LoadCitiesCsv() {
	fileName := "cities.csv"
	file, err := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "City,Country") {
				continue;
			}
			items := strings.Split(line, ",")
			populateCountriesDB(items)
		}
	}
}

func populateCountriesDB(items []string) {
	city, country := items[0], items[1]
	population, _ := strconv.ParseInt(items[2], 10, 64)
	cityInstance := City{Name: city, Population: population}
	if len(CountriesDB[country]) == 0 {
		CountriesDB[country] = make([]City, 0)
	}
	citiesInCountry := CountriesDB[country]
	citiesInCountry = append(citiesInCountry, cityInstance)
	CountriesDB[country] = citiesInCountry
}
