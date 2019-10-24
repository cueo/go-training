package main

import (
	"./lab08"
	"fmt"
)

func main() {
	lab08.LoadCitiesCsv()
	printCitiesByCountry()
	printCityWithSmallestAndLargestPopulation()
	printChinaPopulation()
}

func printChinaPopulation() {
	var population int64 = 0
	for _, city := range lab08.CountriesDB["China"] {
		population += city.Population
	}
	fmt.Println("China's Population", population)
}

func printCityWithSmallestAndLargestPopulation() {
	largest, smallest := lab08.CountriesDB["China"][0], lab08.CountriesDB["China"][0]

	for _, cities := range lab08.CountriesDB {
		for i := 0; i < len(cities); i++ {
			if cities[i].Population > largest.Population {
				largest = cities[i]
			}
			if cities[i].Population < smallest.Population {
				smallest = cities[i]
			}
		}
	}
	fmt.Println(largest, smallest)
}

func printCitiesByCountry() {
	for country, cities := range lab08.CountriesDB {
		fmt.Println("***** " + country + " *****")
		for _, city := range cities {
			fmt.Println(city)
		}
	}
}
