package learning

import "fmt"

// type NodePrinter struct { //это тип полями которого являтся другие типы

// }

type InterPrinter interface { // это тип полями которого являются методы (описывает поведение)
	Printer()
}

func (q QueueNodeStr) Printer() {
	fmt.Printf("Id:  %s, QuantityElems: %s, Status: %s \n", q.Id, q.QuantityElems, q.Status)
}

// принтер для очереди который в цикле вызывает принтер для ноды (q QueueNodeStr)printer

func (s *NodeStack) Printer() {
	fmt.Printf("Id:  %s, Parameters: %s, Status: %s \n", s.Id, s.Parameters, s.Status)
}
