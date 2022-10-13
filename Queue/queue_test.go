package queue

import (
	"errors"
	"fmt"
	"testing"
)

func TestQueue_1(t *testing.T) {

	queue := NewQueueInterfaceMock(t)
	queue.PushMock.Set(func(p QueueNodeInt) (QueueNodeInt, error) {
		return p, errors.New(fmt.Sprintf("error %d", p))
	})

	// t.Run("push_error_lenght&cap", func(t *testing.T) {
	// 	lenght := 0
	// 	q := NewQueueStructInt(lenght)
	// 	elemq := &QueueNodeInt{
	// 		Id:         0,
	// 		Status:     0,
	// 		Parameters: 0,
	// 	}
	// 	_, err := q.Push(*elemq)
	// 	require.Error(t, err)
	// 	require.EqualError(t, ErrorLenNil, err.Error())
	// 	require.EqualError(t, ErrorNoElem, err.Error())
	// })

	//	t.Run("push_error_lenght", func(t *testing.T) {
	//		lenght := 0
	//		q := NewQueueStructInt(lenght)
	//		elemq := &QueueNodeInt{
	//			Id:         1,
	//			Status:     1,
	//			Parameters: 1,
	//		}
	//		_, err := q.Push(*elemq)
	//		fmt.Println(q.Cap)
	//		require.Error(t, err)
	//		require.EqualError(t, ErrorLenNil, err.Error())
	//		require.EqualError(t, ErrorNoElem, err.Error())
	//	})
}

/*
func TestQueueErr(t *testing.T) {
	testTable := []struct {
		name   string
		lenght int
		elemq  *QueueNodeInt
		queue  *QueueStructInt
		err    error
	}{
		{
			name:   "push_error_lenght",
			lenght: 0,
			elemq: &QueueNodeInt{
				Id:         0,
				Status:     0,
				Parameters: 0,
			},
			queue: NewQueueStructInt(1),
			err:   ErrorNoElem,
		},
		{
			name:   "push_error_cap",
			lenght: 0,
			elemq: &QueueNodeInt{
				Id:         1,
				Status:     1,
				Parameters: 1,
			},
			queue: NewQueueStructInt(0),
			err:   ErrorLenNil,
		},
	}
	q := testTable[0]
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := q.queue.Push(*q.elemq)
			require.Error(t, err)
			require.EqualError(t, testCase.err, err.Error())
		})
		// t.Run(testCase.name, func(t *testing.T) {

		// })
	}
}
*/
