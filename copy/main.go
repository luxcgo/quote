package main

import "fmt"

func (d *Driver) SetTrips(trips []*Trip) {
	d.trips = make([]*Trip, len(trips))
	copy(d.trips, trips)

	d.trips = trips

}

type Trip struct {
	Name string
}

type Driver struct {
	trips []*Trip
}

func main() {
	d1 := &Driver{}
	trips := []*Trip{{"A"}}
	d1.SetTrips(trips)

	trips[0].Name = "...."

	fmt.Println(d1.trips[0])
}
