package main

/***_________QUEUE_____________***/

/**___________First struct________**/
//создать др. структуру и сделать
/*
func main() {
	// queue := learning.QueueStructInt{}
	newqueue := learning.NewQueueStructInt(7)
	a := learning.QueueNodeInt{Id: 12, Status: 12}
	b := learning.QueueNodeInt{Id: 13, Status: 13}

	// queue.Push(a)
	newqueue.Push(b)
	fmt.Println(newqueue)

	// elem1, errpull := newqueue.Pull()
	// fmt.Println(elem1, errpull)
	_, err := newqueue.Push(a)
	fmt.Println(newqueue, err)
	_, newerr := newqueue.Push(b)
	fmt.Println(newqueue, newerr)
	_, newerr2 := newqueue.Push(learning.QueueNodeInt{Id: 2000})
	fmt.Println(newqueue, err, newerr, newerr2)

	// newpop, poperr := newqueue.Pull()
	// fmt.Println(newqueue, newpop, poperr)

	// newpop, poperr = newqueue.Pull()
	// fmt.Println(newqueue, newpop, poperr)
	// newpop, poperr = newqueue.Pull()
	// fmt.Println(newqueue, newpop, poperr)

	_, err = newqueue.Push(learning.QueueNodeInt{Id: 489})
	fmt.Println(newqueue, err)

	newpop2, poperr2 := newqueue.Pop(b)
	fmt.Println(newqueue, newpop2, poperr2)
	newpop2, poperr2 = newqueue.Pop(b)
	fmt.Println(newqueue, newpop2, poperr2)

	// lenghtqueue := newqueue.LenghtQueue()
	// fmt.Println(lenghtqueue)

	// index, delnode, errdel := newqueue.DelChange(1, a)
	// fmt.Println(index, delnode, errdel, "\n", newqueue)

	// node, err := queue.Pull()
	// newnode, newerr := newqueue.Pull()
	// node, err := newqueue.Pull()

	// fmt.Println(node, err, "\n", newnode, newerr)
	// InterfacePush(newqueue)
	// var inter learning.QueueInterface
	// inter = queue
	// inter.Push(a)
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
	// // queue := learning.NewQueueStructStr(2)
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

// func main() {
// 	// stack := &learning.StackStruct{}
// 	var stackelem3, stackelem4, stackelem5 learning.NodeStack
// 	// queue := learning.QueueStructInt{}
// 	newqueue := learning.NewQueueStructInt(7)

// 	newqueue.Push(learning.QueueNodeInt{Id: 12, Status: 12})
// 	newqueue.Push(learning.QueueNodeInt{Id: 13, Status: 13})

// 	newstack := learning.NewStackStructInt(4)
// 	// stackelem := learning.NodeStack{Id: 11, Status: 12, Parameters: 13}
// 	// stackelem2 := learning.NodeStack{Id: 21, Status: 22, Parameters: 23}
// 	// stackelem3 = learning.NodeStack{Id: 31, Status: 32, Parameters: 33}
// 	// stackelem4 = learning.NodeStack{Id: 41, Status: 42, Parameters: 43}
// 	// stackelem5 = learning.NodeStack{Id: 51, Status: 52, Parameters: 53}
// 	newstack.Push(stackelem)
// 	fmt.Println(newstack)
// 	newstack.Push(stackelem2)
// 	fmt.Println(newstack)
// 	newstack.Push(stackelem3)
// 	newstack.Push(stackelem4)
// 	fmt.Println(newstack)
// 	newstack.Push(stackelem5)
// 	fmt.Println(newstack)
// 	node, err := newstack.Pull()
// 	fmt.Println(node, err, "\n", newstack)

// 	node, err = newstack.Pop(stackelem3)
// 	fmt.Println(node, err, newstack)

// elem, err := newstack.Pop(stackelem)
// fmt.Println("\n", elem, " ", err, newstack)

// elem, err = newstack.Pop()
// fmt.Println("\n", elem, " ", err, newstack)
// Printer(newstack)

/*_______Interface__________*/

// 	var stackinter learning.StackInterface
// 	// stackinter = newstack
// 	StackInterf(stackinter)
// }

// func Printer(printerinter learning.PrinterIface) {
// 	printerinter.Print()
// }

// func StackInterf(stackinter learning.StackInterface) {
// 	stackinterElem := learning.NodeStack{Id: 110, Status: 111, Parameters: 112}
// 	stackinter.Push(stackinterElem)
// 	pullnode, pullerr := stackinter.Pull(stackinterElem)
// 	// node, err := stackinter.Pop(stackinterElem)
// 	fmt.Println("Stack Interface: ", stackinter, "\n", "Param Pull: ", pullnode, pullerr, " ")
// }

//**___________Printer_________**/

// func main() {
// 	queueStr := &learning.QueueStructStr{}
// 	qelemStr := learning.QueueNodeStr{Id: "12", QuantityElems: "123", Status: "143"}
// 	selem := &learning.NodeStack{Id: 12}
// 	queueStr.Push(qelemStr)
// 	node, err := queueStr.Pull(qelemStr)
// 	PrintInterf(qelemStr)
// 	PrintInterf(selem)
// 	fmt.Println("first struct: ", node, " ", err)
// 	queueStr.Pop(qelemStr)
// 	fmt.Println()

// 	// delem, _ := queue.Pop(elem)
// 	// fmt.Println("del struct: ", node, "\n", delem)
// }

// func PrintInterf(print learning.InterPrinter) {
// 	print.Printer()
// }

// __________________________________________________________________

/*
type Node struct {
	Value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		// fmt.Println(q.head, q.size, q.tail, q.nodes)
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
	// fmt.Println(q.head, q.tail, q.size, q.count, "\n", n)
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func main() {

	q := NewQueue(2)
	q.Push(&Node{2})
	q.Push(&Node{4})
	q.Push(&Node{6})
	q.Push(&Node{8})
	fmt.Println(q.Pop(), q.Pop(), q.Pop(), q.Pop())
}
*/

///__________Interfaces_____________///

/*
func main() {
	newqueue := learning.NewQueueStructInt(3)

	newqueue.Push(learning.QueueNodeInt{Id: 12, Status: 12})
	newqueue.Push(learning.QueueNodeInt{Id: 13, Status: 13})

	newstack := learning.NewStackStructInt(3)

	newstack.Push(learning.NodeStack{Id: 51, Status: 52, Parameters: 53})
	newstack.Push(learning.NodeStack{Id: 41, Status: 42, Parameters: 43})

	PrintAnyObject(newqueue)
	PrintAnyObject(newstack)

	dog := &Dog{Speed: 300}
	cat := &Cat{Speed: 2500}

	AngryBirds(dog, dog.Speed)
	AngryBirds(cat, cat.Speed)
}

type PrinterIface interface { // структура удовл. интерфейсу, если у нее есть такой же набор функций //
	Print()
}

type Runner interface {
	Run(Speed int)
}

type Dog struct {
	Speed int
}

type Cat struct {
	Speed int
}

func (blablabla *Dog) Run(Speed int) {
	fmt.Println("Dog running!", Speed)
}

func (c *Cat) Run(Speed int) {
	fmt.Println("Cat running!", Speed)
}

func AngryBirds(r Runner, speed int) {
	r.Run(speed)
}

func PrintAnyObject(AnyObject PrinterIface) {
	AnyObject.Print()

}
*/
/*
func main() {
	// air := learning.NewAir(0, 0, false, 0, 0)
	// air := learning.Air{A: learning.NewPassengersArray(2)}
	// car := learning.Car{PassengersArray: learning.NewPassengersArray(2)}
	// passengerParam := learning.Passenger{Key: 123, PassportNum: 124}

	// fmt.Println(air.A)
	// _, err := air.A.Push(passengerParam)
	// passengerParam = learning.Passenger{Key: 223, PassportNum: 224}
	// _, err = car.Push(passengerParam)
	// fmt.Println(air.A, "\n", car.PassengersArray, err)

	Boeing := learning.NewAir(1, 1, false, 1200, 4000)
	BMW := learning.NewCar(1, 1, false, 120, 2000)
	Honda := learning.NewBike(1, 1, false, 100, 1500)
	Moving(Boeing)
	Moving(BMW)
	Moving(Honda)

}

func Moving(t learning.Transport) {
	t.Move()
	t.Add(learning.Passenger{Key: 2312, PassportNum: 123123})
}
*/

//	func Add(x, y int) int {
//		return x + y
//	}
