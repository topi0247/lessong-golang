package lesson3

import (
	"fmt"
	"strconv"
)

func Lesson3() {
	/*
	   整数型
	   https://www.w3schools.com/go/go_integer_data_type.php
	   同じint型でも、int8, int16, int32, int64などがあり、型が異なると計算できない
	*/
	var i int = 100
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	fmt.Println(i, i8, i16, i32, i64) // => 100 127 32767 2147483647 9223372036854775807

	/*
	   浮動小数点型
	   https://www.w3schools.com/go/go_float_data_type.php
	   同じfloat型でも、float32, float64などがあり、型が異なると計算できない
	*/
	var f32 float32 = 3.1415926535
	var f64 float64 = 3.1415926535
	fmt.Println(f32, f64) // => 3.1415927 3.1415926535

	/*
	   論理値
	   https://www.w3schools.com/go/go_boolean_data_type.php
	*/
	var t, f bool = true, false
	fmt.Println(t, f) // => true false

	/*
	   文字列型
	   https://www.w3schools.com/go/go_string_data_type.php
	*/
	var s string = "fighting dreamers!"
	fmt.Println(s) // => We are fighting dreamers!
	// 改行込み
	var s2 string = `信じるがままに
    Oli oli oli oh!
        Just go my way!`
	fmt.Println(s2)

	// 型を知る方法
	fmt.Printf("%T\n", i) // => int

	// 型パース
	fmt.Printf("%T\n", float32(i)) // => float32

	// ゼロで割る
	// fmt.Println(1 / 0) // => nvalid operation: division by zero

	/*
	   byte(uint8)型
	   https://www.w3schools.com/go/go_byte_data_type.php
	*/
	byteA := []byte{72, 73}
	fmt.Println(byteA)         // => [72 73]
	fmt.Println(string(byteA)) // => HI
	c := []byte("HI")
	fmt.Println(c) // => [72 73]

	/*
	   配列型
	   https://www.w3schools.com/go/go_arrays.php
	*/
	var arr [3]int
	fmt.Println(arr)        // => [0 0 0]
	fmt.Printf("%T\n", arr) // => [3]int

	var arr2 [5]string = [5]string{"Right", "here,", "right", "now", "(bang!)"}
	fmt.Println(arr2) // => [Right here, right now (bang!)]

	arr3 := [...]string{"ぶっ放せ", "like", "a", "弾丸ライナー！"}
	fmt.Println(arr3)        // => [ぶっ放せ like a 弾丸ライナー！]
	fmt.Printf("%T\n", arr3) // => [4]string

	/*
	   インターフェース型
	*/
	var iX interface{}
	fmt.Println(iX) // => <nil>
	iX = 100
	fmt.Println(iX) // => 100
	iX = "Right here, right now (bang!)"
	fmt.Println(iX) // => Right here, right now (bang!)
	// iX = 3.1415926535
	// fmt.Printf(iX + 1) // => mymodule/lesson3/lesson3.go:90:14: invalid operation: iX + 1 (mismatched types interface{} and int)

	// 型変換
	var ften float64 = 10.0
	var ten int = int(ften)
	fmt.Println(ten) // => 10
	var str2 string = "100"
	fmt.Println(str2 + "yen") // => 100yen
	i, _ = strconv.Atoi(str2)
	fmt.Println(i + 1) // => 101
}
