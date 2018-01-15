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

	parsed := parseVehicles(vehicles, vehicleGroups)

	println(fmt.Sprintf("%+v", parsed))
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

func parseVehicles(vehicles []Vehicle, groups []VehicleGroup) (todo interface{}) {
	// do work here

	return
}

