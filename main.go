package main

import (
	"fmt"
	learning "golang-learning/basic/Learning"
)

func main() {
	square := learning.Square{5}
	printShapeArea(square)

}

func printShapeArea(shape Shape) { // shape - объект Shape - интерфейс Area - объект
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

/*
________________приведение типов___________________

func main() {
	printInterface("Prikol")
	printInterface(20)
}

func printShapeArea(shape Shape) { // shape - объект Shape - интерфейс Area - объект
	fmt.Println(shape.Area())
}

func printInterface(i interface{}) {
	str, exists := i.(string)
	if !exists {
		fmt.Println("interface is not string") // неверно указан тип
		return
	}
	fmt.Println(len(str)) // выводит длину текста
}


*/
/*
____________________Интерфейсы_____________________

_________________Объединение интерфейсов (огрегирование)(Композиция интерфейсов)_____________

type Shape interface { // должен содержать все методы первого и второго интерфейса
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct { // чтоб объект соответствовал интерфейсу shape он должен имплементировать все методы первого и второго интерфейса
	sideLenght float32
}

func (s Square) Perimeter() float32 {
	return s.sideLenght * 4
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

func main() {
	square := Square{5}
	printShapeArea(square)

}

func printShapeArea(shape Shape) { // shape - объект Shape - интерфейс Area - объект
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

_______________Type Switch_________________________

func printShapeArea(shape Shape) { // shape - объект Shape - интерфейс Area - объект
	fmt.Println(shape.Area())
}

func printInterface(i interface{}) { // любой тип любая перемнная имплементирует пустой интерфейс
	switch value := i.(type) { // приведение интерфейса к определенному типу)
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("boot", value)
	default:
		fmt.Println("unknown type")
	}
	fmt.Printf("%v\n", i)
}
_______________Нулевой интерфейс_____________________

func main() {
	printInterface("Prikol")
	printInterface(true)
}

func printInterface(i interface{}) { // любой тип любая перемнная имплементирует пустой интерфейс
	fmt.Printf("%v\n", i)
}

________________11111111111111________________
type Shape interface {
	Area() float32
}

type Square struct {
	sideLenght float32
}

func (s Square) Area() float32 { // структура которая имплиментируют метод Area() (т.е. сигнатура аналогична интерфейсу - тип один и тот же возвращаемые знаяения и т.д.)
	return s.sideLenght * s.sideLenght
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{5}
	circle := Circle{8}

	printShapeArea(square)
	printShapeArea(circle)
}

func printShapeArea(shape Shape) { // shape - объект Shape - интерфейс Area - объект
	fmt.Println(shape.Area())
}

*/

/*
____________________Методы структуры_______________

___________________Приведение типа____________________
type Index int

func (i Index) isNorm() bool {
	return i >= 10
}

type Parameters struct {
	designator   string
	index        Index
	mode         bool
	position     int
	speed        int
	acceleration int
}

func (p *Parameters) printParametersInfo() { // (p Parameters) - метод ресивер по ссылке (pointer reciver) (работаем с ссылочным объектом)
	fmt.Println(p.designator, p.index, p.mode, p.position, p.position, p.speed, p.acceleration) //
}

func NewParameters(designator string, mode bool, index, position, speed, acceleration int) Parameters {
	return Parameters{
		designator:   designator,
		index:        Index(index), //приведение типа (Название_типа(тип переменных))
		mode:         mode,
		position:     position,
		speed:        speed,
		acceleration: acceleration,
	}
}
func main() {
	Axis_X := NewParameters("Axis_X", true, 1, 1000, 120, 10)
	Axis_Z := NewParameters("Axis_Z", true, 0, 1000, 120, 10)

	Axis_X.printParametersInfo()
	Axis_Z.printParametersInfo()

}

______________________111111111111111111______________

type Parameters struct {
	designator   string
	index        int
	mode         bool
	position     int
	speed        int
	acceleration int
}

func (p Parameters) printParametersInfo() { // (p Parameters) - метод ресивер по значению (value reciver)
	fmt.Println(p.designator, p.index, p.mode, p.position, p.position, p.speed, p.acceleration) //
}

func (p *Parameters) printParametersInfo() { // (p Parameters) - метод ресивер по ссылке (pointer reciver) (работаем с ссылочным объектом)
	fmt.Println(p.designator, p.index, p.mode, p.position, p.position, p.speed, p.acceleration) //
}

func NewParameters(designator string, mode bool, index, position, speed, acceleration int) Parameters {
	return Parameters{
		designator:   designator,
		index:        index,
		mode:         mode,
		position:     position,
		speed:        speed,
		acceleration: acceleration,
	}
}
func main() {
	Axis_X := NewParameters("Axis_X", true, 1, 1000, 120, 10)
	Axis_Z := NewParameters("Axis_Z", true, 0, 1000, 120, 10)

	Axis_X.printParametersInfo()
	Axis_Z.printParametersInfo()

}

*/

/*
________________________Структуры_____________________________

// Структура кастомный тип которая в себе хранит набор полей разных типов и эта структура может обладать различными методами.
	// На основании этой структуры можно создавать объекты, объекты этой структуры


___________________Конструктор_________________________


type parameters struct {
	designator   string
	index        int
	mode         bool
	position     int
	speed        int
	acceleration int
}


// конструктор - функция инициализирует объект определеннойго типа

func NewParameters(designator string, mode bool, index, position, speed, acceleration int) parameters {
	return parameters{
		designator:   designator,
		index:        index,
		mode:         mode,
		position:     position,
		speed:        speed,
		acceleration: acceleration,
	}
}

type DataBase struct {
	m map[string]string
}

func NewDataBase() *DataBase { //
	return &DataBase{
		m: make(map[string]string), // инициализация мапы для полноценной работы в дальнейшем
	}
}

func main() {
	Axis_X := NewParameters("Axis_X", true, 1, 1000, 120, 10)
	Axis_Z := NewParameters("Axis_Z", true, 0, 1000, 120, 10)

	fmt.Printf("%+v\n", Axis_X)
	fmt.Printf("%+v\n", Axis_Z)

	fmt.Println(Axis_Z.designator, Axis_Z.index) // обращение к объекту структуры ( Axis_Z.index)

}

_________________Переиспользованная структура______________


type parameters struct {
	designator   string
	index        int
	mode         bool
	position     int
	speed        int
	acceleration int
}

func main() {
	Axis_X := parameters{"Axis_X", 1, true, 1000, 120, 10}
	Axis_Z := parameters{"Axis_Z", 0, true, 1000, 120, 10}

	fmt.Printf("%+v\n", Axis_X)
	fmt.Printf("%+v\n", Axis_Z)

	fmt.Println(Axis_Z.designator, Axis_Z.index) // обращение к объекту структуры ( Axis_Z.index)

}



______________Непереиспользованная структура__________________
func main() {
	parameters := struct { // объявление структуры
		designator   string
		index        int
		mode         bool
		position     int
		speed        int
		acceleration int
	}{"Axis_X", 1, true, 1000, 150, 5} // инициализация структуры

	fmt.Printf("%+v\n", parameters) // подробная структура вида : {designator:Axis_X index:1 mode:true position:1000 speed:150 acceleration:5}
}


*/

/*
______________________________________Мапы______________________

__________________Удаление элемента мапы______________

func main() {
	users := map[string]int{ // ключ может быть любым типом (в том числе структурой или интерфейсом)
		"Mode":         1,
		"Position":     1000,
		"Acceleration": 20,
		"Speed":        120,
	}

	users["Abs_Pos"] = 56

	for key, value := range users {
		fmt.Println(key, value)
	}
	delete(users, "Mode") // удаляет элемент мапы
	fmt.Println(users)
}

_____________________Присвоение новых значений вне инициализации мапы_______

func main() {
	users := map[string]int{ // ключ может быть любым типом (в том числе структурой или интерфейсом)
		"Mode":         1,
		"Position":     1000,
		"Acceleration": 20,
		"Speed":        120,
	}
	users["Abs_Pos"] = 56 // добавление ключа Abs_Pos
	for key, value := range users {
		fmt.Println(key, value)
	}
}

_______________Проверка элемента на существование в мапе (exists)___________

func main() {
	users := map[string]int{ // ключ может быть любым типом (в том числе структурой или интерфейсом)
		"Mode":         1,
		"Position":     1000,
		"Acceleration": 20,
		"Speed":        120,
	}

	parameters, exists := users["No message"] // exists проверяет существует ли текущий элемент в мапе
	// if exists {
	// 	fmt.Println("Mode", parameters)
	// } else {
	// 	fmt.Println("Нет в списке!")
	// }

	fmt.Println(parameters, exists)
}

________________111111111111111_______________
func main() {
	users := map[string]int{ // ключ может быть любым типом (в том числе структурой или интерфейсом). Ключ всегда д.б. уникальный
		"Mode":         1,
		"Position":     1000,
		"Acceleration": 20,
		"Speed":        120,
	}
	parameters := users["Mode"] // мапу присваиваем переменной
	fmt.Println(parameters) // выводит значение ключа "Mode"
	fmt.Println(users) // выводит мапу с инициализированными значениями
	fmt.Println(users["Abs_Pos"]) // такого ключа нет в мапе, следовательно возвратиться значение 0
}
*/

/*

______________Panic (восстановление паники)_________

func main() {
	defer handlerPanic()
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	messages[4] = "message 5"
	fmt.Println(messages) // не вывведет message так как ошибка не произошла
}

func handlerPanic() {
	if r := recover(); r != nil { // recover -обработка паники. Возвращаемое значение из recover сообщает, паникует ли программа goroutine.
		fmt.Println(r)
	}

	fmt.Println("handlerPanic() ready")
}

func main() {
	defer handlerPanic()
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}
	fmt.Println(messages) // выводит messages так как не произошла ошибка
	messages[4] = "message 5"
}

func handlerPanic() {
	if r := recover(); r != nil { // recover -обработка паники. Возвращаемое значение из recover сообщает, паникует ли программа goroutine.
		fmt.Println(r)
	}

	fmt.Println("handlerPanic() ready")
}


*/

/*
_________Defer_____________

// Сначала выведет "main()", потом "printMessage()"
func main() {
	defer printMessage() // откладывание выполнение функции в самый конец перед выхода из преложения (добавляет 50мс к выполнению программы)
	fmt.Println("main()")
}

func printMessage() {
	fmt.Println("printMessage()")
}

// Сначала выведет "printMessage()", потом "main()"
func main() {
	printMessage()
	fmt.Println("main()")
}

func printMessage() {
	fmt.Println("printMessage()")
}

*/
/*
____________Range________________

func main() {

	messege := []string{
		"messeg 1", "messeg 2", "messeg 3", "messeg 4",
	}
	for i := range messege { //range - итерируется по слайсу
		fmt.Println(messege[i]) // обращаемся по индексу, выводим значение
	}

	for _, i := range messege { //range - итерируется по слайсу
		fmt.Println(i) // обращаемся по индексу, выводим значение
	}
}

*/

/*
______________Matrix_______________

func main() {

	matrix := make([][]int, 10)

	fmt.Println(matrix)

	counter := 0

	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			counter++
			matrix[x][y] = counter
		}
		fmt.Println(matrix[x])
	}

}

*/

/*
______________Массивы. Слайсы__________________

__________________ограниечение append_____

func main() {
	message := make([]string, 10000) // иннициализировался слайс и аллоцировался массив из 5-ти нулевых значений типа string. Функция(слайс, длина, вместимость (capacity))

	message = append(message, "10001")

	fmt.Println(message)

	fmt.Println(len(message)) // выведет 10001
	fmt.Println(cap(message)) // выведет 12800
}

___________________5555_Append_5555_____________

func main() {
	message := make([]string, 5) // иннициализировался слайс и аллоцировался массив из 5-ти нулевых значений типа string. Функция(слайс, длина, вместимость (capacity))

	message = append(message, "6") // append создал новый массив в два раза больше (т.е. размер message = 10), и записал шестой элемент. Произошла переалокаця массива

	fmt.Println(message) // выведет [      6]

	fmt.Println(len(message))
	fmt.Println(cap(message))
}

___________________444444444444_______________

func main() {
	message := make([]string, 5, 15) // иннициализировался слайс и аллоцировался массив из 5-ти нулевых значений типа string. Функция(слайс, длина, вместимость (capacity))
	fmt.Println(len(message))
	fmt.Println(cap(message))
}

func printMessages(messages []string, messages2 []int) error { // функция работает с слайами messages и messages2
	if len(messages) == 0 && len(messages2) == 0 {
		return errors.New(("empty array"))
	}

	messages[1] = "6" // происходит перезапись значений второго элемента слайса messages
	messages2[1] = 6  // происходит перезапись значений второго элемента слайса messages2

	fmt.Println(messages)  // вызовится массив [1, 6, 3]
	fmt.Println(messages2) // вызовится массив [4, 6, 6]
	return nil

}

______________33333333333333_Слайсы_Используя_Make_33333333333333____________

func main() {
	message := make([]string, 5, 15) // иннициализировался слайс и аллоцировался массив из 5-ти нулевых значений типа string. Функция(слайс, длина, вместимость (capacity))
	fmt.Println(len(message)) // длина
	fmt.Println(cap(message)) // вместимость
}

_________________Ошибка при компиляции______________________
func main() {
	var message []string
	message[0] = "1"
	fmt.Println(message) // Компилятор ругнется!!! Т.к. Указатель на массив = <nil>, и длина слайса нуль!
}


______________Нулевой слайс____________________________
func main() {
	var message []string
	fmt.Println(message)  // выведится пустой слайс
}




______________22222222222222_Слайсы_22222222222222______________



	// Слайс - это обертка над массивом. По факту слайс хранит в себе ссылку на массив.
		// 	Когда передаем слайс в функцию -> передается ссылка, а НЕ копия!!!
			// 	Следовательно, если записывать значение по индексу слайса, происходит перезапись значение в массиве на который ссылается данный слайс!!!



func main() {

	messages := []string{"1", "2", "3"}
	messages2 := []int{4, 5, 6}

	printMessages(messages, messages2)

	fmt.Println(messages)  // вызовится массив [1, 6, 3]. Поменялся оригинальный слайс
	fmt.Println(messages2) // вызовится массив [4, 6, 6]. Поменялся оригинальный слайс

}

func printMessages(messages []string, messages2 []int) error { // функция работает с слайами messages и messages2
	if len(messages) == 0 && len(messages2) == 0 {
		return errors.New(("empty array"))
	}

	messages[1] = "6" // происходит перезапись значений второго элемента слайса messages
	messages2[1] = 6  // происходит перезапись значений второго элемента слайса messages2

	fmt.Println(messages)  // вызовится массив [1, 6, 3]
	fmt.Println(messages2) // вызовится массив [4, 6, 6]
	return nil

}


_____________1111111111111_Массивы_1111111111111_____________
func main() {

	messages := [3]string{"1", "2", "3"} // массив
	messages2 := [5]int{}                // массив

	printMessage(messages, messages2)

	fmt.Println(messages) // вызовится массив [1, 2, 3]
	fmt.Println(messages2) // вызовится массив [0, 0, 0, 0, 0]

}

func printMessage(messages [3]string, messages2 [5]int) error { // функция работает с копиями массивов messages и messages2
	if len(messages) == 0 && len(messages2) == 0 {
		return errors.New(("empty array"))
	}

	messages[1] = "5"
	messages2[1] = 5

	fmt.Println(messages) // вызовится массив [1, 5, 3]
	fmt.Println(messages2) // вызовится массив [0, 5, 0, 0, 0]
	return nil

}

*/

/*

____________________Пустой указатель_____________

___________________3333333333333333_____________

func main() {

	number := 5
	var p *int

	p = &number

	fmt.Println(p)       // выведит адрес области памяти number
	fmt.Println(&number) // выведит адрес области памяти number

	*p = 10

	fmt.Println(number) // выведит 10. Т.к. ссылается на адрес памяти number и перезаписывает значение этой переменной
}

____________________222222222222222_____________

func main() {

	number := 5
	var p *int

	p = &number

	fmt.Println(p)       // выведит адрес области памяти number
	fmt.Println(&number) // выведит адрес области памяти number
}

____________________11111111111111_______________

func main() {

	var p *int

	fmt.Println(p) // выведит нулевой тип <nil> . Т.к. нет перменной на кторый он бы мог ссылаться
}


*/

/*
________________Область памяти, указатели_________

__________________________________5555555555555__________________________________



__________________________________4444444444444__________________________________

func main() {
	message := "Скоро я стану каштаном! "

	fmt.Println(&message) // выводит область памяти переменной message 0xc0000761e0
	changeMessage(message)

}

func changeMessage(message string) { // копированный объект он не ссылается на переменную message. Новая переменная которая скопировала значение переменной message
	// *message += "А может и не стану!?" // указатель на область памяти переменной messege и изменение переменной!
	fmt.Println(&message) // выведет адрес области памяти 0xc0000761f0 другая область отличающаяся от message
}


__________________________________3333333333333___________________________________

func main() {
	message := "Скоро я стану каштаном! "

	fmt.Println(&message) // выводит область памяти переменной message
	changeMessage(&message)
	fmt.Println(&message) // выводит область памяти переменной message

}

func changeMessage(message *string) {
	*message += "А может и не стану!?" // указатель на область памяти переменной messege и изменение переменной!
}

__________________________________2222222222222____________________________________

func main() {
	message := "Скоро я стану каштаном! "

	fmt.Println(message) // выводит перменную message. Т.е. "Скоро я стану каштаном! "
	changeMessage(&message) // вызов функции / обращение к области памяти переменной message и изменение значения message
	fmt.Println(message) // выводит измененную строку. Т.е. Скоро я стану каштаном! А может и не стану!?

}

func changeMessage(message *string) {
	*message += "А может и не стану!?" // указатель на область памяти переменной messege и изменение переменной!
}

__________________________________111111111111_____________________________________
func main() {
	message = "Скоро я стану каштаном! "
	printMessage(message) // вызовиться message из функции printMessage(). Т.е.  "Скоро я стану каштаном! А может и не стану!?"
	fmt.Println(message)  // вызовиться значение переменной message. Т.е.  "Скоро я стану каштаном! "

}

func printMessage(message string) {
	message += "А может и не стану!?" // аллоцируется место в стеке
	fmt.Println(message)
}

*/

/*
______Функция init()___________
func init() { // зарегестрированная привелигирированная функция - всегда срабатывает при инициализации пакетов (срабатывает перед ф-ей main())
	message = "Kashtan!"
}
*/

/*
___________________Анонмные функции_________________

//Замыкание - функция которая ссылается на "независимые" свободные переменные (при замыкании функция запоминает состояние в котором она была создана)//

func main() {

	inc := increment()

	fmt.Println(inc(), increment2())
	fmt.Println(inc(), increment2())
	fmt.Println(inc(), increment2())
	fmt.Println(inc(), increment2())

	action(5, 5, func(x int, y int) int { return x + y })
	action(5, 5, func(x int, y int) int { return x * y })

	SF := SelectFunction(1)
	fmt.Println(SF(5, 5))

	func() {
		fmt.Println("анонимная функция")
	}()
}

func SelectFunction(n int) func(int, int) int {
	if n == 1 {
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }
	}
}

// анонимная функция как аргумент функции
func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

func increment() func() int {
	count := 0 //при каждом вызове функции increment() состояние count - сохраняется
	return func() int {
		count++
		return count
	}
}

func increment2() int {
	count := 0
	count++
	return count
}
*/

/*
__________________знакомство с Slices______________

func findMin(numbers ...int) int { // numbers - it's type of "Slice!!"
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	// min := numbers[0]

	for _, i := range numbers {
		if i > max {
			max = i
			// min = i
		}
	}
	return max
}
*/

/*

_______________________Строки________________

func sayHelllo(name string, age uint8) string {
	return fmt.Sprintf("-\tSalam, %s!\n-\tHow old are Your?\n-\tMy Age: %d", name, age) // генерирует(возвращает строку)
}

func FaceControl(age uint8) (string, error) {
	if age >= 18 && age <= 60 {
		return "Like, Kashtan! Welcom to the Army", nil
	} else if age >= 60 && age <= 70 {
		return "It's time for you to retire!", errors.New("you are too old")
	}
	return "Dalbep!", errors.New("uncorrect age!")
}
*/

// func prediction(dayOfWeek string) (string, error) {
// 	switch dayOfWeek {
// 	case "monday":
// 		return "Hard Day", nil
// 	case "tuesday":
// 		return "second working day", nil
// 	case "wednesday":
// 		return "midweek", nil
// 	case "thursday":
// 		return "Yee! Namaz tomorrow", nil
// 	case "friday":
// 		return "Namaz day!", errors.New("Weekend!")
// 	case "saturday":
// 		return "Shabbat!", errors.New("Weekend!")
// 	case "sunday":
// 		return "Sanday!", errors.New("Weekend!")
// 	default:
// 		return "Invalid day!", errors.New("invalid day of the week!")
// 	}
// }
