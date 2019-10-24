package lab07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cities []City

func LoadCities(filepath string) Cities {
	file, err := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	var cities Cities
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		for scanner.Scan() {
			line := scanner.Text()
			city := getCity(line)
			cities = append(cities, city)
		}
	}
	if len(cities) > 0 {
		cities = cities[1:]
	}
	return cities
}

func getCity(line string) (city City) {
	words := strings.Split(line, ",")
	if len(words) == 3 {
		population, _ := strconv.Atoi(words[2])
		city = City{
			Name:       words[0],
			Country:    words[1],
			Population: population,
		}
	}
	return
}
