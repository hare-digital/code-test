package main

import (
	"fmt"
	"io/ioutil"
)

type Vehicle struct {
}

type VehicleGroup struct {
}

func main() {
	vehicles := readVehicles()
	vehicleGroups := readVehicleGroups()

	println(fmt.Sprintf("%+v", vehicles))
	println(fmt.Sprintf("%+v", vehicleGroups))
}

func readVehicles() (vehicle []Vehicle) {
	raw, err := ioutil.ReadFile("./vehicles.json")
	if err != nil {
		panic(err)
	}

	// do work here

	return
}

func readVehicleGroups() (vehicleGroup []VehicleGroup) {
	raw, err := ioutil.ReadFile("./vehicle_groups.json")
	if err != nil {
		panic(err)
	}

	// do work here

	return
}

