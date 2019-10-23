package main

import "fmt"

type Training struct {
	name string
	venue string
}

func main() {
	goTraining := Training{}
	initTraining(goTraining)
	fmt.Println(goTraining)

	var advancedGoTraining *Training = &Training{}
	initTrainingByRef(advancedGoTraining)
	fmt.Println(*advancedGoTraining)

	javaTraining := Training{ "Java", "6A"}
	javaTraining.printDetails()

}

func (training Training) printDetails() {
	fmt.Println("DETAILS: ", training.name, training.venue)
}



func initTrainingByRef(training *Training) {
	(*training).name = "Advanced Go"
	training.venue = "8"
}

//By default pass by value
func initTraining(training Training) {
	training.name = "Go"
	training.venue = "4A"
}
