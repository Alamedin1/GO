package main

import (
	"fmt"
	learning "golang-learning/basic/Learning"
)

/***_________QUEUE_____________***/

/**___________First struct________**/
//создать др. структуру и сделать
/*
func main() {
	queue := &learning.TaskQueue{}
	a := learning.QueueNode{Id: 1}
	queue.Push(a)
	InterfacePush(queue)
	node, err := queue.Pull()
	fmt.Println(node, err)

	var inter learning.QueueInterface
	inter = queue
	inter.Push(a)
}

func InterfacePush(inter learning.QueueInterface) {
	r := learning.QueueNode{Status: 3}
	inter.Push(r)
	node, err := inter.Pull()
	fmt.Println(node, err)

}
*/

/**___________Second struct________**/

func main() {
	/*
		queue := &learning.QueueStructStr{}
		elem := learning.QueueNodeStr{Id: "", QuantityElems: "", Status: "14"}
		queue.Push(elem)
		node, err := queue.Pull(elem)
		fmt.Println(elem, "\n", node, err)
		delem, delerr := queue.Pop()
		fmt.Println("Delet elem!", delem, delerr)
	*/
	/** interface  */

	queue := &learning.QueueStructStr{}
	elem := learning.QueueNodeStr{Id: "", QuantityElems: "", Status: "14"}
	queue.Push(elem)
	node, err := queue.Pull(elem)
	AddElem(queue)
	fmt.Println("first struct: ", node, " ", err)

	var inter learning.QueueStrInterface
	inter = queue
	AddElem(inter)
	inter.Push(node)
	newnode, err := inter.Pull(node)
	fmt.Println("therd struct: ", newnode, " ", err)

}

func AddElem(inter learning.QueueStrInterface) {
	add := learning.QueueNodeStr{Id: "123"}
	inter.Push(add)
	node, err := inter.Pull(add)
	fmt.Println("Second added struct: ", node, " ", err)
	delem, _ := inter.Pop()
	fmt.Println("Remove elem in second struct: ", node, err, "\n", "Delet elem: ", delem)
}
