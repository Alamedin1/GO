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
	queue := &learning.QueueStructInt{}
	a := learning.QueueNodeInt{Id: 1}
	queue.Push(a)
	InterfacePush(queue)
	node, err := queue.Pull()
	fmt.Println(node, err)

	var inter learning.QueueInterface
	inter = queue
	inter.Push(a)
}

func InterfacePush(inter learning.QueueInterface) {
	r := learning.QueueNodeInt{Status: 3}
	inter.Push(r)
	node, err := inter.Pull()
	fmt.Println(node, err)

}
*/
/**___________Second struct________**/
/*
func main() {
	// queue := &learning.QueueStructStr{}
	// queue := learning.NewQueueStructStr(2)
	// elem := learning.QueueNodeStr{Id: "", QuantityElems: "", Status: "14"}
	// queue.Push(elem)
	// fmt.Println(queue.Queue)
	// node, err := queue.Pull(elem)
	// fmt.Println(elem, "\n", node, err)
	// delem, delerr := queue.Pop()
	// fmt.Println("Delet elem!", delem, delerr)

	// a := []int{3, 5}

	q := []learning.QueueNodeStr{
		learning.QueueNodeStr{
			Id:            "12",
			QuantityElems: "23",
		},
		learning.QueueNodeStr{
			Id:            "123",
			QuantityElems: "2324",
		},
	}

	q2 := []learning.QueueStructStr{
		learning.QueueStructStr{
			Queue: []learning.QueueNodeStr{
				learning.QueueNodeStr{
					Id:            "12",
					QuantityElems: "23",
				},
				learning.QueueNodeStr{
					Id:            "12",
					QuantityElems: "23",
				},
			},
		},
		learning.QueueStructStr{
			Queue: []learning.QueueNodeStr{
				learning.QueueNodeStr{
					Id:            "12",
					QuantityElems: "23",
				},
			},
		},
	}

	fmt.Println(q, "\n", q2)

	// var b []int     // не выделяем память только инийиализируем переменную
	// b = []int{3, 6} // выделяем память для элементов

	// c := make([]int, 1) // инициализация только память без значений

	var inter learning.InterPrinter
	inter = learning.QueueNodeStr{}
	inter.Printer()
	inter = learning.NodeStack{}
	inter.Printer()

}
*/
/** interface  */
/*
	queue := &learning.QueueStructStr{}
	elem := learning.QueueNodeStr{}
	queue.Push(elem)
	node, err := queue.Pull(elem)
	AddElem(queue)
	fmt.Println("first struct: ", node, " ", err)
	delem, _ := queue.Pop(elem)
	fmt.Println("del struct: ", node, "\n", delem)
	// var inter learning.QueueStrInterface
	// inter = queue
	// AddElem(inter)
	// inter.Push(node)
	// newnode, err := inter.Pull(node)
	// fmt.Println("therd struct: ", newnode, " ", err)

}

func AddElem(inter learning.QueueStrInterface) {
	add := learning.QueueNodeStr{Id: "123"}
	inter.Push(add)
	node, err := inter.Pull(add)
	fmt.Println("Second added struct: ", node, " ", err)
	// delem, _ := inter.Pop()
	// fmt.Println("Remove elem in second struct: ", node, err, "\n", "Delet elem: ", delem)
}
*/

/***_________STACK_____________***/
/*
func main() {
	stack := &learning.StackStruct{}
	stackelem := learning.NodeStack{Id: 12, Status: 13, Parameters: 14}
	stack.Push(stackelem)
	elem, err := stack.Pop(stackelem)
	fmt.Println(stack.Pop(stackelem))
	fmt.Println("\n", elem, " ", err)
*/
/*_______Interface__________*/
/*
	var stackinter learning.StackInterface
	stackinter = stack
	StackInterf(stackinter)
}

func StackInterf(stackinter learning.StackInterface) {
	stackinterElem := learning.NodeStack{Id: 120, Status: 12}
	stackinter.Push(stackinterElem)
	node, err := stackinter.Pop(stackinterElem)
	fmt.Println("Stack Interface: ", node, err)
}
*/

//**___________Printer_________**/

func main() {
	queueStr := &learning.QueueStructStr{}
	qelemStr := learning.QueueNodeStr{Id: "12", QuantityElems: "123", Status: "143"}
	selem := &learning.NodeStack{Id: 12}
	queueStr.Push(qelemStr)
	node, err := queueStr.Pull(qelemStr)
	PrintInterf(qelemStr)
	PrintInterf(selem)
	fmt.Println("first struct: ", node, " ", err)
	queueStr.Pop(qelemStr)
	fmt.Println()

	// delem, _ := queue.Pop(elem)
	// fmt.Println("del struct: ", node, "\n", delem)
}

func PrintInterf(print learning.InterPrinter) {
	print.Printer()
}
