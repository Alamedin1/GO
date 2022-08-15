package learning

import (
	"errors"
	"fmt"
)

type Transport interface {
	Move()
	Add(node Passenger) (Passenger, error)
}

type Passenger struct {
	Key, PassportNum int
}

type PassengersArray struct {
	PassengerParam           []Passenger
	Tail, Head, Numelem, Cap int
}

func NewPassengersArray(Lenght int) *PassengersArray {
	return &PassengersArray{
		PassengerParam: make([]Passenger, Lenght),
		Tail:           0,
		Head:           0,
		Numelem:        0,
		Cap:            Lenght,
	}
}

func (this *PassengersArray) Add(node Passenger) (Passenger, error) {
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

type Air struct { //приватные поля с маленькой буквы не видны из внешних пакетов и функции
	Passengers      *PassengersArray
	Wings, Turbine  int
	Motor           bool
	Speed, Distance int
}

func (a *Air) Add(node Passenger) (Passenger, error) {
	return a.Passengers.Add(node)
}

type Car struct {
	*PassengersArray
	Hood, Turbine   int
	Motor           bool
	Speed, Distance int
}

type Bike struct {
	*PassengersArray
	Rack, TelescopicFork int
	Motor                bool
	Speed, Distance      int
}

func NewAir(PassengersAmount, Wings, Turbine int, Motor bool,
	Speed, Distance int) *Air {
	return &Air{
		Passengers: NewPassengersArray(PassengersAmount),
		Wings:      Wings,
		Turbine:    Turbine,
		Motor:      Motor,
		Speed:      Speed,
		Distance:   Distance,
	}
}
func NewCar(PassengersAmount, Hood, Turbine int, Motor bool,
	Speed, Distance int) *Car {
	return &Car{
		PassengersArray: NewPassengersArray(PassengersAmount),
		Hood:            Hood,
		Turbine:         Turbine,
		Motor:           Motor,
		Speed:           Speed,
		Distance:        Distance,
	}
}
func NewBike(PassengersAmount, Rack, TelescopicFork int, Motor bool,
	Speed, Distance int) *Bike {
	return &Bike{
		PassengersArray: NewPassengersArray(PassengersAmount),
		Rack:            Rack,
		TelescopicFork:  TelescopicFork,
		Motor:           Motor,
		Speed:           Speed,
		Distance:        Distance,
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
