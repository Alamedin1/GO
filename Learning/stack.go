package learning

/***_________STACK_____________***/

type StackStruct struct {
	Id, Status, Parameters [3]int
}

type NodeStack struct {
	Stack []StackStruct
}

type StackInterface interface {
	Push(o StackStruct) NodeStack
	Pop() *StackStruct
}

// func (q *Queue) Initialize() *Queue {
// 	q.Vqueue = []Queue{}
// 	return q
// }

func (q *NodeStack) Push(Max StackStruct) *NodeStack {
	q.Stack = append(q.Stack, Max)
	return q
}

func (q *NodeStack) Pop() *StackStruct {
	Min := q.Stack[0]
	q.Stack = q.Stack[1:len(q.Stack)]
	return &Min
}
