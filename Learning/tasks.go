package learning

import "errors"

/***_________STACK_____________***/

// type Stack struct {
// 	Id, Status, Parameters [3]int
// }

// type TaskStack struct {
// 	VStack []Stack
// }

// type QueueInterface interface {
// 	Push(o Stack) TaskStack
// 	Pop() *Stack
// }

// // func (q *Queue) Initialize() *Queue {
// // 	q.Vqueue = []Queue{}
// // 	return q
// // }

// func (q *TaskStack) Push(Max Stack) *TaskStack {
// 	q.VStack = append(q.VStack, Max)
// 	return q
// }

// func (q *TaskStack) Pop() *Stack {
// 	Min := q.VStack[0]
// 	q.VStack = q.VStack[1:len(q.VStack)]
// 	return &Min
// }

/***_________QUEUE_____________***/

/**___________First struct________**/

// добавление получение удаление

type QueueNode struct {
	Id, Status, Parameters int
}

type TaskQueue struct {
	Vqueue []QueueNode // нода (node)
}

// type QueueInterface interface {
// 	Push(o Queue) TaskQueue
// 	Pop() *Queue
// }

// func (q *Queue) Initialize() *Queue {
// 	q.Vqueue = []Queue{}
// 	return q
// }

type QueueInterface interface {
	Push(node QueueNode) QueueNode
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

// func (q *TaskQueue) Pop(node QueueNode) error {
// 	// удалить
// }

/**___________Second_Struct_______**/

/*



func (q *TaskQueue) Initialize() *TaskQueue {
	q.Vqueue = []Queue{}
	return q
}

func (q *TaskQueue) Push(o Queue) *TaskQueue {
	q.Vqueue = append(q.Vqueue, o)
	return q
}

func (q *TaskQueue) Pop() *Queue {
	min := q.Vqueue[0]
	q.Vqueue = q.Vqueue[1:len(q.Vqueue)]
	return &min
}

*/
