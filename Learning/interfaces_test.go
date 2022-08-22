package learning

import (
	"queue"
	"testing"
)

func TestAir(t *testing.T) {
	// mockCtrl := gomock.Controller(t)
	// defer mockCtrl.Finish()

	queue.NewQueueStructInt(12)

	mockAir := queue.NewQueueInterfaceMock(t)

	// air := NewAir(queue., 12, 32, true, 12, 43)
	// t.Log(air)
}
