package main

import "fmt"

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
