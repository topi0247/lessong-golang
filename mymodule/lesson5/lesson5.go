package lesson5

import "fmt"

func Lesson5() {
	// 足し算
	fmt.Println(1 + 2) // => 3
	fmt.Println("ABC" + "DRF") // => ABCDRF
	num := 0
	num += 5
	fmt.Println(num) // => 5
	num++
	fmt.Println(num) // => 6
	str := "Hello"
	str += " World"
	fmt.Println(str) // => Hello World

	// 引き算
	fmt.Println(1 - 2) // => -1
	num -= 3
	fmt.Println(num) // => 3
	num--
	fmt.Println(num) // => 2

	// 掛け算
	fmt.Println(2 * 3) // => 6

	// 割り算
	fmt.Println(10 / 3) // => 3

	// 余り
	fmt.Println(10 % 3) // => 1

	// 比較演算子
	fmt.Println(1 == 1) // => true
	fmt.Println(1 != 1) // => false
	fmt.Println(1 < 1) // => false
	fmt.Println(1 <= 1) // => true
	fmt.Println(1 > 1) // => false
	fmt.Println(1 >= 1) // => true

	// 論理演算子
	fmt.Println(true && true) // => true
	fmt.Println(true && false) // => false
	fmt.Println(true || false) // => true
	fmt.Println(!true) // => false
}
