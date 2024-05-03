package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	peter := Passenger{"Peter", 1, true}
	fmt.Println(peter)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		emma = Passenger{Name: "Emma", TicketNumber: 3}
	)
	fmt.Println(bill, emma)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	peter.Boarded = true
	bill.Boarded = true

	if peter.Boarded {
		fmt.Println("Peter has boarded the bus")
	}
	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{FrontSeat: heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
