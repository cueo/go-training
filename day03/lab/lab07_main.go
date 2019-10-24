package main

import (
	"./lab07"
	"fmt"
)

func main() {
	sep := "\n===========\n"

	filename := "/Users/mmayank/Documents/codes/go/Oct-22-24-2019/day03/lab/lab07/cities.csv"
	cities := lab07.LoadCities(filename)

	fmt.Println(sep)
	countryMap := groupByCountries(cities)
	fmt.Println("Cities by country:\n", countryMap)
	fmt.Println(sep)

	m, l := findPopulationExtreme(cities)
	fmt.Println("Most populated", m.Name, m.Population)
	fmt.Println("Least populated", l.Name, l.Population)
	fmt.Println(sep)

	cp := getChinaPopulation(cities)
	fmt.Println("Population of China", cp)
	fmt.Println(sep)
}

func getChinaPopulation(cities lab07.Cities) (p int) {
	for _, city := range cities {
		if city.Country == "China" {
			p += city.Population
		}
	}
	return
}

func findPopulationExtreme(cities lab07.Cities) (m lab07.City, l lab07.City) {
	max, min := cities[0].Population, cities[0].Population
	for _, city := range cities {
		if max < city.Population {
			max = city.Population
			m = city
		} else if min > city.Population {
			min = city.Population
			l = city
		}
	}
	return
}

func groupByCountries(cities lab07.Cities) (m map[string]lab07.Cities) {
	m = make(map[string]lab07.Cities)
	for _, city := range cities {
		m[city.Country] = append(m[city.Country], city)
	}
	return
}
