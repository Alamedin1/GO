package learning

import (
	"errors"
)

/***_________QUEUE_____________***/

/**___________First struct________**/
/*
type QueueNode struct {
	Id, Status, Parameters int
}

type TaskQueue struct {
	Vqueue []QueueNode // нода (node)
}

type QueueInterface interface {
	Push(node QueueNode) QueueNode
	Pull() (QueueNode, error)
}

func (q *TaskQueue) Push(node QueueNode) QueueNode {
	q.Vqueue = append(q.Vqueue, node)
	return node
}

func (q *TaskQueue) Pull() (QueueNode, error) { // вытащить элемент с головы
	if len(q.Vqueue) == 0 {
		return QueueNode{}, errors.New("Empty Queue!")
	} else {
		head := q.Vqueue[0]
		q.Vqueue = q.Vqueue[1:len(q.Vqueue)]
		return head, nil
	}
}
*/
// func (q *TaskQueue) Pop(node QueueNode) (QueueNode, error) {

// }

/**___________Second struct________**/

type QueueNodeStr struct {
	Id, QuantityElems, Status string
}

type QueueStructStr struct {
	Queue []QueueNodeStr
}

type QueueStrInterface interface {
	Pull(node QueueNodeStr) (QueueNodeStr, error)
	Push(node QueueNodeStr) (QueueNodeStr, error)
	Pop() (QueueNodeStr, error)
}

func (q *QueueStructStr) Push(node QueueNodeStr) (QueueNodeStr, error) {
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

func (q *QueueStructStr) Pop() (QueueNodeStr, error) {
	if len(q.Queue) == 0 {
		return QueueNodeStr{}, errors.New("Empty Queue!")
	}
	FirstElem := q.Queue[0]
	q.Queue = q.Queue[1:]
	// fmt.Println("Delet Elem", FirstElem)
	return FirstElem, nil

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
