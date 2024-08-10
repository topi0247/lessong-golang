package lesson2

import "fmt"

// test := 500 // これはエラーになる
var test int = 500

func Lesson2() {
	// var 変数名 データ型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "We are fighting dreamers!"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		x int    = 200
		y string = "高みを目指して"
	)
	fmt.Println(x, y)

	var num int
	var str string
	fmt.Println(num, str)

	num = 300
	str = "fighting dreamers!"
	fmt.Println(num, str)

	autoStr := "なりふり構わず"
	fmt.Println(autoStr)

	// autoStr := "再定義は出来ない"
	// fmt.Println(autoStr)

	// autoStr = 300
	// fmt.Println(autoStr)

	fmt.Println(test)

	Lesson2_2()

	// 使われない変数があるとエラー
	// var notUse int = 100
}

func Lesson2_2() {
	fmt.Println("別の関数")
}
