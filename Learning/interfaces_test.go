package learning

import (
	"errors"
	"testing"
)

// заполнить самолет и получить всех пассажиров самолета

func TestAir(t *testing.T) {

	// mockCtrl := gomock.Controller(t)
	// defer mockCtrl.Finish()

	// mockPassengerQueue := queue.NewQueueInterfaceMock(t) //queue.NewQueueInterfaceMock(t)

	// air := NewAir(queue., 12, 32, true, 12, 43)
	// t.Log(air)

	// air := NewAir(
	// 	&PassengersQueueIfaceMock{},
	// 	12, 13, true, 800, 1200,
	// )
	testTable := []struct {
		name string
		Air  Air
		err  error
	}{
		{
			name: "",
			Air: Air{
				Passengers: &PassengersQueueIfaceMock{},
				Wings:      10,
				Turbine:    11,
				Motor:      true,
				Speed:      1000,
				Distance:   2000,
			},
			err: errors.New(""),
		},
	}
	q := testTable[0]
	t.Log(q)
}
