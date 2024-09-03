package lesson7

import (
	"fmt"
	"strconv"
	"os"
	"time"
)

func Lesson7() {
	// 条件分岐
	x := 1
	if x == 1 {
		fmt.Println("x is 1")
	} else if x == 2 {
		fmt.Println("x is 2")
	} else {
		fmt.Println("x is not 1")
	}

	// 短縮文
	if y := 2; y == 2 {
		fmt.Println("y is 2")
	} else {
		fmt.Println("y is not 2")
	}

	// スコープが優先される
	x2 := 0
	if x2 := 2; true {
		fmt.Println(x2) // => 2
	}
	fmt.Println(x2) // => 0

	// エラーハンドリング
	var str string = "A"
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err) // => strconv.Atoi: parsing "A": invalid syntax
	}
	fmt.Printf("i = %T\n", i) // => i = int

	// for ループ文
	// 下記は無限ループ
	// for {
	// 	fmt.Println("loop")
	// }

	// while文っぽいの
	index := 0
	for {
		index++
		if index == 3 {
			break
		}
		fmt.Println("index is", index)
	} // => index is 1, index is 2

	// 条件付き, while文っぽいの
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	} // => 0 1 2 3 4 5 6 7 8 9

	// C言語っぽいの
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	} // => 0 1 2 3 4 5 6 7 8 9

	// スキップしたり抜けたり
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // 3 はスキップ
		}
		if i == 5 {
			break // 5 で抜ける
		}
		fmt.Println(i)
	} // => 0 1 2 4

	// 配列
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	} // => 1 2 3

	// 範囲式, each文っぽいの
	arr2 := [3]int{1, 2, 3}
	for i, v := range arr2 {
		fmt.Println(i, v)
	} // => 0 1, 1 2, 2 3

	sl := []string{"A", "B", "C"}
	for i, v := range sl {
		fmt.Println(i, v)
	} // => 0 A, 1 B, 2 C

	// map, 辞書型っぽいの, 連想配列
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	} // => apple 100, banana 200

	// switch文
	n := 1
	switch n { // データ型はcaseと合わせる必要がある
	case 1:
		fmt.Println("n is 1")
	case 2:
		fmt.Println("n is 2")
	default:
		fmt.Println("n is not 1 nor 2")
	} // => n is 1

	// 短縮文
	switch n2 := 2; n2 {
	case 1:
		fmt.Println("n2 is 1")
	case 2:
		fmt.Println("n2 is 2")
	default:
		fmt.Println("n2 is not 1 nor 2")
	} // => n2 is 2

	// 複数の条件
	n3 := 3
	switch n3 {
	case 1, 2:
		fmt.Println("n3 is 1 or 2")
	case 3, 4:
		fmt.Println("n3 is 3 or 4")
	default:
		fmt.Println("n3 is not 1 nor 2")
	} // => n3 is 3 or 4

	// 条件式
	n4 := 4
	switch {
	case n4 == 1:
		fmt.Println("n4 is 1")
	case n4 == 2:
		fmt.Println("n4 is 2")
	default:
		fmt.Println("n4 is not 1 nor 2")
	} // => n4 is not 1 nor 2

	// 演算子と組み合わせるとエラー
	// switch n5 := 5; n5 {
	// case 1:
	// 	fmt.Println("n5 is 1")
	// case 2:
	// 	fmt.Println("n5 is 2")
	// case n5 > 3:
	// 	fmt.Println("n5 is greater than 3")
	// default:
	// 	fmt.Println("n5 is not 1 nor 2")
	// } // cannot convert n5 > 3 (untyped bool value) to type int

	// 型スイッチ
	anything("Hello")
	anything(1)

	var x3 interface{} = 3
	// i2 := x3.(int) // 型アサーション
	i2, isInt := x3.(int)
	fmt.Println(i2, isInt) // => 3 true

	// f := x3.(float64) // panic: interface conversion: interface {} is int, not float64
	f, isFloat64 := x3.(float64)
	fmt.Println(f, isFloat64) // => 0 false

	// 型スイッチ
	if x3 == nil {
		fmt.Println("None")
	} else if i3, isInt := x3.(int); isInt { // 短縮文: 型アサーション, 条件式
		fmt.Println(i3, "x3 is int")
	} else if s, isString := x3.(string); isString { // 短縮文: 型アサーション, 条件式
		fmt.Println(s, "x3 is string")
	} else {
		fmt.Println("I don't know")
	}

	switch x3.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}  // => int

	switch v := x3.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println("I don't know")
	} // => 3 int


	// ラベル付きfor
	Loop:
		for {
			for {
				for {
					fmt.Println("start")
					break Loop
				}
				fmt.Println("break 2")
			}
			fmt.Println("break 1")
		}
		fmt.Println("end") // => start end

		for i := 0; i < 3; i++ {
			for j := 1; j < 3; j++ {
				if j > 1 {
					continue
				}
				fmt.Println(i, j, i*j)
			}
			fmt.Println("end")
		} // => 0 1 0, end, 1 1 1, end, 2 1 2, end

		Loop2:
			for i := 0; i < 3; i++ {
				for j := 1; j < 3; j++ {
					if j > 1 {
						continue Loop2
					}
					fmt.Println(i, j, i*j)
				}
				fmt.Println("end")
			} // => 0 1 0, 1 1 1, 2 1 2

	// defer
	TestDefer() // => START, END

	defer func() {
		fmt.Println("defer1")
		fmt.Println("defer2")
		fmt.Println("defer3")
	}() // => main関数が終わった後に defer1, defer2, defer3

	RunDefer() // => 3, 2, 1

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello")) // => test.txt に Hello が書き込まれる

	// // panic & recover
	// defer func() {
	// 	if x := recover(); x != nil {
	// 		fmt.Println(x)
	// 	}
	// }() // => runtime error
	// panic("runtime error") // => panic: runtime error

	// // Goの並行処理
	// go sub() // ゴルーチン

	// for {
	// 	fmt.Println("main loop")
	// 	time.Sleep(200 * time.Millisecond)
	// }

	// init
	fmt.Println("main")
}

func anything(a interface{}) {
	fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v, "string")
	case int:
		fmt.Println(v, "int")
	default:
		fmt.Println("I don't know")
	} // => Hello string, 1 int
}

func TestDefer(){
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer(){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

// init関数はmain関数よりも先に実行される
func init() {
	fmt.Println("init")
}