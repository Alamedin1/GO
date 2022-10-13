package learning

import (
	"fmt"
)

//go:generate minimock -i PassengersQueueIface -o ./passengersqueue_mock.go

type Transport interface {
	Move()
	Add(node Passenger) (Passenger, error)
}

type PassengersQueueIface interface {
	Push(node Passenger) (Passenger, error)
	Pop(node Passenger) (Passenger, error)
	Pull() (Passenger, error)
	List() []Passenger
}

type Passenger struct {
	Name             string
	Key, PassportNum int
}

// type PassengersArray struct {
// 	PassengerParam           []Passenger
// 	Tail, Head, Numelem, Cap int
// }

// func NewPassengersArray(lenght int) *PassengersArray {
// 	return &PassengersArray{
// 		PassengerParam: make([]Passenger, lenght),
// 		Tail:           0,
// 		Head:           0,
// 		Numelem:        0,
// 		Cap:            lenght,
// 	}
// }

// func (this *PassengersArray) Add(node Passenger) (Passenger, error) {
// 	if this.Cap == 0 {
// 		return node, errors.New("capacity is nil")
// 	}
// 	if this.Tail == this.Cap {
// 		return node, errors.New("capacity limit exceeded")
// 	} else {
// 		this.PassengerParam[this.Tail] = node
// 		this.Tail++
// 		// q.Tail = (q.Tail + 1) % q.Cap
// 		this.Numelem++
// 		// fmt.Println(q.Iqueue[0], q.Tail, q.Numelem)
// 	}
// 	return node, nil
// }

// сделать вывод самого популярного имени пассажира в самолете
// у air будет метод самых популярного имени пассажира( само много раз встречающегося имени в самолете)
// добавить имя пассажира

type Air struct { //приватные поля с маленькой буквы не видны из внешних пакетов и функции
	Passengers      PassengersQueueIface
	Wings, Turbine  int
	Motor           bool
	Speed, Distance int
}

func (a *Air) GetAllPassengers() ([]Passenger, error) {
	return nil, nil
}

func (a *Air) Add(node Passenger) (Passenger, error) {
	return Passenger{}, nil
}

type Car struct {
	Passengers      PassengersQueueIface
	Hood, Turbine   int
	Motor           bool
	Speed, Distance int
}

type Bike struct {
	Passengers           PassengersQueueIface
	Rack, TelescopicFork int
	Motor                bool
	Speed, Distance      int
}

func NewAir(passengers PassengersQueueIface, wings, turbine int, motor bool,
	speed, distance int) *Air {
	return &Air{
		Passengers: passengers,
		Wings:      wings,
		Turbine:    turbine,
		Motor:      motor,
		Speed:      speed,
		Distance:   distance,
	}
}
func NewCar(PassengersAmount, Hood, Turbine int, Motor bool,
	Speed, Distance int) *Car {
	return &Car{
		// PassengersArray: NewPassengersArray(PassengersAmount),
		Hood:     Hood,
		Turbine:  Turbine,
		Motor:    Motor,
		Speed:    Speed,
		Distance: Distance,
	}
}
func NewBike(PassengersAmount, Rack, TelescopicFork int, Motor bool,
	Speed, Distance int) *Bike {
	return &Bike{
		// PassengersArray: NewPassengersArray(PassengersAmount),
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
