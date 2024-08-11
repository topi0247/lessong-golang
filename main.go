package main

import (
	"fmt"
	"app/mymodule/lesson1"
	"app/mymodule/lesson2"
	"app/mymodule/lesson3"
	"app/mymodule/lesson4"
	"app/mymodule/lesson5"
)

func main() {
	fmt.Println("====== Lesson1 ======")
	lesson1.Lesson1()
	fmt.Println("====== Lesson2 ======")
	lesson2.Lesson2()
	fmt.Println("====== Lesson3 ======")
	lesson3.Lesson3()
	fmt.Println("====== Lesson4 ======")
	lesson4.Lesson4()
	fmt.Println("====== Lesson5 ======")
	lesson5.Lesson5()
}
