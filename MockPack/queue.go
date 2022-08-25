package mock

import "errors"

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
}
