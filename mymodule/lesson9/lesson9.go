package lesson9

import (
	"fmt"
)

func Lesson9() {
	// ポインタ
	var n int = 100
	fmt.Println(n) // 100

	fmt.Println(&n)	// 0xc00001205

	Double(n)

	fmt.Println(n)	// 100

	var p *int = &n
	fmt.Println(p)	// 0xc00001205
	fmt.Println(p)	// 0xc00001205

	*p = 300
	fmt.Println(n) // 300
	n = 500
	fmt.Println(*p) // 500

	Doublev2(&n)
	fmt.Println(n)	// 1000

	var sl []int = []int{1, 2, 3}
	Doublev3(sl)
	fmt.Println(sl)	// [2 4 6]
}

func Double(n int) {
	n = n * 2
}

func Doublev2(n *int) {
	*n = *n * 2
}

func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}
