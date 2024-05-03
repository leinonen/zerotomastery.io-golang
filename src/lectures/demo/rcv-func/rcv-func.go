package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNumber int) {
	lot.spaces[spaceNumber-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNumber int) {
	lot.spaces[spaceNumber-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNumber int) {
	lot.spaces[spaceNumber-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 10)}
	occupySpace(&lot, 1)
	lot.occupySpace(2)
	fmt.Println(lot)

	lot.vacateSpace(1)
	lot.vacateSpace(2)
	fmt.Println(lot)
}
