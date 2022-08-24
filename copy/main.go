package main

import (
	"fmt"
	"log"
)

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

func main1() {
	d1 := &Driver{}
	trips := []*Trip{{"A"}}
	d1.SetTrips(trips)

	trips[0].Name = "...."

	fmt.Println(d1.trips[0])
}

func main() {
	t1 := &Trip{
		Name: "t1",
	}

	t2 := new(Trip)

	*t2 = *t1
	log.Println(t1.Name)
	log.Println(t2.Name)

	t2.Name = "t2"

	log.Println(t1.Name)
	log.Println(t2.Name)

}
