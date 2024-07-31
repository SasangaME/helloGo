package basics

import "fmt"

type Vehicle struct {
	color      string
	noOfWheels int
	make       string
	model      string
}

func (v Vehicle) start() {
	fmt.Println("vehicle is starting")
}

type Car struct {
	Vehicle
	engineCapacity int
}

func Structs() {
	car := Car{engineCapacity: 2000, Vehicle: Vehicle{
		color:      "red",
		noOfWheels: 4,
		make:       "Honda",
		model:      "Civic",
	}}
	fmt.Println(car.make)
	car.start()
}
