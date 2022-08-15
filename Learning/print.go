package learning

import "fmt"

type NodePrinter struct { //это тип полями которого являтся другие типы

}

type PrinterIface interface { // это тип полями которого являются методы (описывает поведение)
	Print()
}

// func (q *QueueStructInt) Print() {
// 	fmt.Println(q.Iqueue)
// }

// принтер для очереди который в цикле вызывает принтер для ноды (q QueueNodeStr)printer

func (s *StackStruct) Print() {
	fmt.Println(s.Stack)
}
