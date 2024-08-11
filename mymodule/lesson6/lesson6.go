package lesson6

import "fmt"

func Lesson6() {
	// 関数
	fmt.Println(Sum(1, 2)) // => 3
	fmt.Println(Diff(5, 3)) // => 2
	q, r := Div(10, 3)
	fmt.Println(q, r) // => 3 3
	fmt.Println(Double(100)) // => 200
	Print() // => lesson6

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(1, 2)) // => 3

	f1 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(f1) // => 3

	// 関数を返す関数
	f2 := ReturnFunc()
	f2() // => I'm a function

	// 関数を引数に取る関数
	// コールバック関数
	CallFunction(func(){
		fmt.Println("I'm a function")
	}) // => I'm a function

	// クロージャー
	f3 := Later()
	fmt.Println(f3("Golang")) // => ""
	fmt.Println(f3("is")) // => "Golang"
	fmt.Println(f3("simple")) // => "is"
	fmt.Println(f3("language")) // => "simple"

	// ジェネレーター
	ints := integers()
	fmt.Println(ints()) // => 1
	fmt.Println(ints()) // => 2
	fmt.Println(ints()) // => 3
}

func Sum(x int, y int) int {
	return x + y
}

func Diff(x, y int) int {
	return x - y
}

func Div(x, y int) (int, int){
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Print() {
	fmt.Println("lesson6")
	return
}

func ReturnFunc() func() {
	return func(){
		fmt.Println("I'm a function")
	}
}

func CallFunction(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
