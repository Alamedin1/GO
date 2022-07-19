package learning

import "errors"

/***_________QUEUE_____________***/

/**___________First struct________**/

type QueueNodeInt struct {
	Id, Status, Parameters int
}

type QueueStructInt struct {
	Iqueue []QueueNodeInt // нода (node)
}

type QueueInterface interface {
	Push(node QueueNodeInt) QueueNodeInt
	Pull() (QueueNodeInt, error)
}

func (q *QueueStructInt) Push(node QueueNodeInt) QueueNodeInt {
	q.Iqueue = append(q.Iqueue, node)
	return node
}

func (q *QueueStructInt) Pull() (QueueNodeInt, error) { // вытащить элемент с головы
	if len(q.Iqueue) == 0 {
		return QueueNodeInt{}, errors.New("Empty Queue!")
	} else {
		head := q.Iqueue[0]
		q.Iqueue = q.Iqueue[1:len(q.Iqueue)]
		return head, nil
	}
}

func (q *QueueStructInt) Pop(node QueueNodeInt) (QueueNodeInt, error) {
	if len(q.Iqueue) == 0 {
		return QueueNodeInt{}, errors.New("Empty Queue!")
	} else {
		q.Iqueue = q.Iqueue[:len(q.Iqueue)-1]
		return node, nil
	}
}

// func (q *TaskQueue) Pop(node QueueNode) (QueueNode, error) {

// }

/**___________Second struct________**/

type QueueNodeStr struct { // ddefinition
	Id            string
	QuantityElems string
	Status        string
}

type que QueueNodeStr

type QueueStructStr struct {
	Queue []QueueNodeStr
}

func NewQueueStructStr(lenght int) *QueueStructStr { // initialization
	return &QueueStructStr{
		Queue: make([]QueueNodeStr, lenght),
	}

}

type QueueStrInterface interface {
	Pull(node QueueNodeStr) (QueueNodeStr, error)
	Push(node QueueNodeStr) (QueueNodeStr, error)
	// Pop() (QueueNodeStr, error)
}

func (q *QueueStructStr) Push(node QueueNodeStr) (QueueNodeStr, error) { // написать через head & tail
	q.Queue = append(q.Queue, node)
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
