package learning

import (
	"errors"
)

/***_________STACK_____________***/

type NodeStack struct {
	Id, Status, Parameters int
}

type StackStruct struct {
	Stack []NodeStack
}

type StackInterface interface {
	Push(node NodeStack) NodeStack
	Pop(node NodeStack) (NodeStack, error)
}

func (s *StackStruct) Push(node NodeStack) NodeStack {
	s.Stack = append(s.Stack[:len(s.Stack)], node)
	return node
}

func (s *StackStruct) Pop(node NodeStack) (NodeStack, error) {
	if len(s.Stack) == 0 {
		return node, errors.New("empty stack!")
	} else {
		node := s.Stack[len(s.Stack)-1]
		s.Stack = s.Stack[:len(s.Stack)-1]
		return node, nil
		// L := len(s.Stack)
		// return s.Stack[:L-1], nil
	}
}
