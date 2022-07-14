package main

import (
	"fmt"
	learning "golang-learning/basic/Learning"
)

// type Queue struct {
// 	Id     int
// 	Pos    int
// 	Speed  int
// 	Status int
// }

//создать др. структуру и сделать

func main() {
	queue := &learning.TaskQueue{}
	a := learning.QueueNode{Id: 1}
	queue.Push(a)
	InterfacePush(queue)
	node, err := queue.Pull()
	fmt.Println(node, err)

	// var inter learning.QueueInterface
	// inter = queue
	// inter.Push(a)
	// fmt.Println()
}

func InterfacePush(inter learning.QueueInterface) {
	r := learning.QueueNode{Status: 3}
	inter.Push(r)

}
