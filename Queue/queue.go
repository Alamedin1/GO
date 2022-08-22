package queue

import "errors"

/***_________QUEUE_____________***/

/**___________First struct________**/

//go:generate minimock -i QueueInterface

type QueueInterface interface {
	Push(node QueueNodeInt) (QueueNodeInt, error)
	Pop(node QueueNodeInt) (QueueNodeInt, error)
	Pull() (QueueNodeInt, error)
}

type QueueNodeInt struct {
	Id, Status, Parameters int
}

type QueueStructInt struct {
	Iqueue                   []QueueNodeInt // нода (node)
	Cap, Head, Tail, Numelem int
}

func NewQueueStructInt(lenght int) *QueueStructInt {
	return &QueueStructInt{
		Iqueue:  make([]QueueNodeInt, lenght),
		Cap:     lenght,
		Head:    0,
		Tail:    0,
		Numelem: 0,
	}
}

var ErrorQueueEmp = errors.New("Empty Queue!")
var ErrorNoElem = errors.New("No element in queue")
var ErrorLenNil = errors.New("capacity is nil")
var ErrorCapLimit = errors.New("capacity limit exceeded")

func (q *QueueStructInt) Push(node QueueNodeInt) (QueueNodeInt, error) {
	if q.Cap == 0 {
		return node, ErrorLenNil
	}
	if q.Tail == q.Cap && q.Cap != 0 {
		return node, ErrorCapLimit
	} else {
		q.Iqueue[q.Tail] = node
		q.Tail++
		// q.Tail = (q.Tail + 1) % q.Cap
		q.Numelem++
		// fmt.Println(q.Iqueue[0], q.Tail, q.Numelem)
	}
	return node, nil
}

// func (q *QueueStructInt) Push(node QueueNodeInt) QueueNodeInt {

// 	q.Iqueue = append(q.Iqueue, node)
// 	return node
// }

// func (q *QueueStructInt) Pull() (QueueNodeInt, error) { // вытащить элемент с головы
// 	if len(q.Iqueue) == 0 {
// 		return QueueNodeInt{}, errors.New("Empty Queue!")
// 	} else if q.Numelem == 0 {
// 		return QueueNodeInt{}, errors.New("No elems")
// 	} else {
// 		return q.Iqueue[q.Head], nil
// 	}
// }

func (q *QueueStructInt) Pull() (QueueNodeInt, error) {
	if len(q.Iqueue) == 0 {
		return QueueNodeInt{}, ErrorQueueEmp
	} else if q.Numelem == 0 {
		return QueueNodeInt{}, ErrorNoElem
	} else {
		node := q.Iqueue[q.Head]
		for i := q.Head; i < q.Tail-1; i++ {
			q.Iqueue[i] = q.Iqueue[i+1]
		}
		q.Iqueue[q.Tail-1] = QueueNodeInt{}
		q.Tail--
		q.Numelem--
		return node, nil
	}
}

func (q *QueueStructInt) LenghtQueue() int {
	return (q.Tail - q.Head + q.Cap) % q.Cap
}

func (q *QueueStructInt) Pop(node QueueNodeInt) (QueueNodeInt, error) {
	Index := -1
	for i := q.Head; i < q.Tail; i++ {
		if q.Iqueue[i] == node {
			Index = i
			break
		}
	}
	if Index != -1 {
		for i := Index; i < q.Tail-1; i++ {
			q.Iqueue[i] = q.Iqueue[i+1]
		}
		q.Iqueue[q.Tail-1] = QueueNodeInt{}
	} else {
		return node, ErrorNoElem
	}
	return node, nil

	// if index == q.Head {
	// 	q.Pull()
	// 	return index, node, errors.New("delet Head = index")
	// } else if index > q.Head {
	// 	node = q.Iqueue[index]
	// 	q.Iqueue[index] = QueueNodeInt{}
	// 	q.Head = index
	// 	q.Numelem--
	// 	return index, node, nil
	// } else if index > q.Head && (q.Iqueue[index] == QueueNodeInt{}) ||
	// 	index < q.Head && (q.Iqueue[index] == QueueNodeInt{}) {
	// 	return index, node, errors.New("this elem has been removed ")
	// }
	// return index, node, nil
}

// func (q *TaskQueue) Pop(node QueueNode) (QueueNode, error) {

// }

/**___________Second struct________**/
/*
type QueueNodeStr struct { // ddefinition
	Id            string
	QuantityElems string
	Status        string
}

type QueueStructStr struct {
	Queue []QueueNodeStr
}

func NewQueueStructStr(Lenght int) *QueueStructStr { // initialization
	return &QueueStructStr{
		Queue: make([]QueueNodeStr, Lenght),
	}
}

type QueueStrInterface interface {
	Pull(node QueueNodeStr) (QueueNodeStr, error)
	Push(node QueueNodeStr) (QueueNodeStr, error)
	// Pop() (QueueNodeStr, error)
}

func (q *QueueStructStr) Push(node QueueNodeStr) (QueueNodeStr, error) { // написать через head & tail
	// q.Queue = append(q.Queue, node)
	q.Queue = []
	Tail := q.Queue[len(q.Queue)-1]
	Head := q.Queue[0]
	Lenght := len(q.Queue)

	// q.Queue = node
	// q.Queue[] = node
	// Tail := q.Queue[len(q.Queue)-1]
	fmt.Println(Head, Tail, Lenght)
	return node, nil
}

// func (q *QueueStructStr) Pull(node QueueNodeStr) (QueueNodeStr, error) { // вытащить элемент с головы
// 	// if len(q.Queue) == 0 {
// 	// 	return QueueNodeStr{}, errors.New("Empty Queue!")
// 	// } else {

// 	if node.QuantityElems == "" {
// 		return QueueNodeStr{}, errors.New("Empty QuantityElems in Queue!")
// 	} else if node.Id == "" {
// 		return QueueNodeStr{}, errors.New("Empty Id in Queue!")
// 	} else if node.Status == "" {
// 		return QueueNodeStr{}, errors.New("Empty Status in Queue!")
// 	} else {
// 		head := q.Queue[0]
// 		q.Queue = q.Queue[1:len(q.Queue)]
// 		return head, nil
// 	}
// }

func (q *QueueStructStr) Pull(node QueueNodeStr) (QueueNodeStr, error) {
	Err := ErrorString(node)
	head := q.Queue[0]
	q.Queue = q.Queue[1:len(q.Queue)]
	return head, Err
}

func (q *QueueStructStr) Pop(node QueueNodeStr) (QueueNodeStr, error) {
	if len(q.Queue) == 0 {
		return QueueNodeStr{}, errors.New("Empty Queue!")
	} else {
		q.Queue = q.Queue[:len(q.Queue)-1]
		return node, nil
	}
}

func ErrorString(node QueueNodeStr) error {
	a := ""
	switch a {
	case node.Id:
		return errors.New("Empty Id in Queue!")
	case node.QuantityElems:
		return errors.New("Empty QuantityElems in Queue!")
	case node.Status:
		return errors.New("Empty Status in Queue!")
	}
	return nil
}

func (q *QueueStructStr) Empty() error {
	if len(q.Queue) == 0 {
		return errors.New("Empty Queue!")
	}
	return nil
}
*/
