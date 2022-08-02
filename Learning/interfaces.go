package learning

import (
	"errors"
	"fmt"
)

type Transport interface {
	Move()
}

type Passenger struct {
	Key, PassportNum int
}

type NodePassenger struct {
	PassengerParam           []Passenger
	Tail, Head, Numelem, Cap int
}

func NewPassenger(Lenght int) *NodePassenger {
	return &NodePassenger{
		PassengerParam: make([]Passenger, Lenght),
		Tail:           0,
		Head:           0,
		Numelem:        0,
		Cap:            Lenght,
	}
}

func (this *NodePassenger) Push(node Passenger) (Passenger, error) {
	if this.Cap == 0 {
		return node, errors.New("capacity is nil")
	}
	if this.Tail == this.Cap {
		return node, errors.New("capacity limit exceeded")
	} else {
		this.PassengerParam[this.Tail] = node
		this.Tail++
		// q.Tail = (q.Tail + 1) % q.Cap
		this.Numelem++
		// fmt.Println(q.Iqueue[0], q.Tail, q.Numelem)
	}
	return node, nil
}

type Air struct {
	// NodePassenger
	Wings, Turbine  int
	Motor           bool
	Speed, Distance int
}

type Car struct {
	// NodePassenger
	Hood, Turbine   int
	Motor           bool
	Speed, Distance int
}

type Bike struct {
	// NodePassenger
	Rack, TelescopicFork int
	Motor                bool
	Speed, Distance      int
}

func NewAir(Wings, Turbine int, Motor bool,
	Speed, Distance int) *Air {
	return &Air{
		Wings:    Wings,
		Turbine:  Turbine,
		Motor:    Motor,
		Speed:    Speed,
		Distance: Distance,
	}
}
func NewCar(Hood, Turbine int, Motor bool,
	Speed, Distance int) *Car {
	return &Car{
		Hood:     Hood,
		Turbine:  Turbine,
		Motor:    Motor,
		Speed:    Speed,
		Distance: Distance,
	}
}
func NewBike(Rack, TelescopicFork int, Motor bool,
	Speed, Distance int) *Bike {
	return &Bike{
		Rack:           Rack,
		TelescopicFork: TelescopicFork,
		Motor:          Motor,
		Speed:          Speed,
		Distance:       Distance,
	}
}

func (a Air) Move() {
	if a.Speed > 0 && a.Distance > 0 {
		a.Wings = 1
		a.Turbine = 1
		a.Motor = true
		fmt.Println("the plane is flying!")
	} else {
		fmt.Println("Speed or Distance = 0")
	}
}

func (a Car) Move() {
	if a.Speed > 0 && a.Distance > 0 {
		a.Hood = 1
		a.Turbine = 1
		a.Motor = true
		fmt.Println("the car on the way!")
	} else {
		fmt.Println("Speed or Distance = 0")
	}
}

func (a Bike) Move() {
	if a.Speed > 0 && a.Distance > 0 {
		a.Rack = 1
		a.TelescopicFork = 1
		a.Motor = true
		fmt.Println("the bike on the way!")
	} else {
		fmt.Println("Speed or Distance = 0")
	}
}
