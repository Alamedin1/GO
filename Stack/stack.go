package learning

import (
	"errors"
)

/***_________STACK_____________***/

type NodeStack struct {
	Id, Status, Parameters int
}

type StackStruct struct {
	Stack                    []NodeStack
	Cap, Head, Tail, Numelem int
}

func NewStackStructInt(lenght int) *StackStruct {
	return &StackStruct{
		Stack:   make([]NodeStack, lenght),
		Cap:     lenght,
		Head:    0,
		Tail:    0,
		Numelem: 0,
	}
}

type StackInterface interface {
	Push(node NodeStack) (NodeStack, error)
	Pull(node NodeStack) (NodeStack, error)
	Pop(node NodeStack) (NodeStack, error)
}

func (s *StackStruct) Push(node NodeStack) (NodeStack, error) {
	// s.Stack = append(s.Stack[:s.Head], node)
	if s.Head == s.Cap {
		s.Stack = append(s.Stack, node)
		s.Head++
		s.Cap++
		s.Numelem++
		return node, errors.New("stack is full")
	} else {
		s.Stack[s.Head] = node
		s.Head++
		s.Numelem++
		return node, nil
	}
}

func (s *StackStruct) Pull() (NodeStack, error) {
	if s.Head == s.Tail {
		return NodeStack{}, errors.New("empty stack!")
	} else {
		node := s.Stack[s.Head-1]
		s.Stack[s.Head-1] = NodeStack{}
		s.Head--
		s.Numelem--
		return node, nil
	}
}

func (s *StackStruct) Pop(node NodeStack) (NodeStack, error) {
	Index := -1
	for i := s.Tail; i < s.Head; i++ {
		if s.Stack[i] == node {
			Index = i
			break
		}
	}
	if Index != -1 {
		for i := Index; i < s.Head-1; i++ {
			s.Stack[i] = s.Stack[i+1]
		}
		s.Stack[s.Head-1] = NodeStack{}
		s.Head--
	} else {
		return node, errors.New("No element in stack")
	}
	return node, nil
}

///
