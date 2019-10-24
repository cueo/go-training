* In this exercise you'll use __cities.csv__ file to perform some operations
* Create a package __lab08__
* Create a file __types.go__ where you define the __City struct__
* Create a file __db.go__ where you have functions to load the cities.csv, convert every item to a struct and maintain in a collection
* Create __lab08_main.go___ in main package that does the following

* Print all the cities grouped by countries
* Print the city name with largest and smallest population
* Print China's population

* Code to load the file is given below. You can use it and modify it appropriately 

```go
import (
	"bufio"
	"os"
)
	fileName := "cities.csv"
	file, err := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
``` 