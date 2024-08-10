package lesson4

import "fmt"

// 定数
// 頭文字を大文字にするとグローバル宣言で他のパッケージでも使える
const Pi = 3.14

// 一括代入できる
// 下記の場合、A, B, C は 1、D, E, F は "A" になる
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// これはオーバーフロー
// cannot use 9223372036854775807 + 1 (untyped int constant 9223372036854775808) as int value in variable declaration (overflows)
// var Big int = 9223372036854775807 + 1
// これはオーバーフローを回避できる
const Big = 9223372036854775807 + 1

const (
	c0 = iota // 0
	c1
	c2
)

func Lesson4() {
	fmt.Println(Pi)

	// 定数は再代入できない
	// mymodule/lesson4/lesson4.go:13:2: cannot assign to Pi (neither addressable nor a map index expression)
	// Pi = 3.14159
	// fmt.Println(Pi)

	fmt.Println(A, B, C, D, E, F) // => 1 1 1 A A A

	fmt.Println(Big - 1) // => 9223372036854775807

	fmt.Println(c0, c1, c2) // => 0 1 2
}