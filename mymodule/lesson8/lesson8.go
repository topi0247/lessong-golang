package lesson8

import (
	"fmt"
	// "time"
)

func Lesson8() {
	// スライスの宣言
	// 配列の場合は要素数を指定したが、スライスの場合は指定しない
	var sl []int
	fmt.Println(sl) // => []

	// スライスの初期化
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2) // => [100 200]

	// 暗黙的な宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4) // => [0 0 0 0 0]

	sl2[0] = 1000
	fmt.Println(sl2) // => [1000 200]

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5[0]) // => 1
	fmt.Println(sl5[2:4]) // => [3 4]
	fmt.Println(sl5[:2]) // => [1 2]
	fmt.Println(sl5[2:]) // => [3 4 5]
	fmt.Println(sl5[:]) // => [1 2 3 4 5]
	fmt.Println(sl5[len(sl5)-1]) // => 5
	fmt.Println(sl5[1:len(sl5)-1]) // => [2 3 4]

	// スライスのappend
	sl = []int{100, 200}
	fmt.Println(sl) // => [100 200]
	sl = append(sl, 300)
	fmt.Println(sl) // => [100 200 300]
	sl = append(sl, 400, 500)
	fmt.Println(sl) // => [100 200 300 400 500]

	// make関数
	sl6 := make([]int, 5)
	fmt.Println(sl6) // => [0 0 0 0 0]

	// len関数
	fmt.Println(len(sl6)) // => 5

	// cap関数
	fmt.Println(cap(sl6)) // => 5
	sl7 := make([]int, 5, 10)
	fmt.Println(sl7) // => [0 0 0 0 0]
	fmt.Println(len(sl7)) // => 5
	fmt.Println(cap(sl7)) // => 10
	sl7 = append(sl7, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sl7) // => [0 0 0 0 0 1 2 3 4 5 6 7]
	fmt.Println(len(sl7)) // => 12
	fmt.Println(cap(sl7)) // => 20
	// len = 実際に値が格納されている要素数
	// cap = 容量、スライスの要素数の上限
	// capは容量を超えると倍の容量を再確保する

	// スライスのコピー
	// スライスは参照型
	sl8 := []int{1, 2, 3}
	sl9 := sl8

	sl9[0] = 1000
	fmt.Println(sl8) // => [1000 2 3]
	// スライスは参照型なので、値渡しではなく参照渡しになる

	// 基本型の場合
	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i, i2) // => 10 100

	// スライスをコピーする場合
	sl10 := []int{1, 2, 3, 4, 5}
	sl11 := make([]int, 5, 10)
	fmt.Println(sl10, sl11) // => [1 2 3 4 5] [0 0 0 0 0]
	n := copy(sl11, sl10)
	fmt.Println(sl10, sl11, n) // => [1 2 3 4 5] [1 2 3 4 5] 5

	// スライスのループ
	sl12 := []string{"A", "B", "C"}
	fmt.Println(sl12) // => [A B C]
	for i, v := range sl12 {
		fmt.Println(i, v)
	} // => 0 A, 1 B, 2 C

	// 中の要素だけを使う
	for _, v := range sl12 {
		fmt.Println(v)
	} // => A, B, C

	// インデックスだけを使う
	for i := range sl12 {
		fmt.Println(i)
	}

	// 古典技法
	for i := 0; i < len(sl12); i++ {
		fmt.Println(sl12[i])
	}

	// 可変長引数
	fmt.Println(Sum(1, 2, 3, 4, 5)) // => 15
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // => 55
	fmt.Println(Sum()) // => 0

	sl13 := []int{1, 2, 3}
	fmt.Println(Sum(sl13...)) // => 6

	// map
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m) // => map[A:100 B:200]

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2) // => map[A:100 B:200]

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3) // => map[1:A 2:B]

	m4 := make(map[int]string)
	fmt.Println(m4) // => map[]
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4) // => map[1:JAPAN 81:USA]

	fmt.Println(m["A"]) // => 100
	fmt.Println(m4[2]) // => USA
	fmt.Println(m4[3]) // => 空文字

	s, ok := m4[1]
	fmt.Println(s, ok) // => JAPAN true
	// 2つ目の戻り値は、取り出しに成功したかどうか
	// 存在しない場合は、空文字とfalseが返る
	_, ok2 := m4[3]
	if !ok2 {
		fmt.Println("Key Not Found")
	}
	fmt.Println(ok2) // => false

	m4[2] = "US"
	fmt.Println(m4) // => map[1:JAPAN 2:US]

	m4[3] = "CHINA"
	fmt.Println(m4) // => map[1:JAPAN 2:US 3:CHINA]

	delete(m4, 2) // 第一引数に対象のmapを指定し、第二引数に削除したいキーを指定
	fmt.Println(m4) // => map[1:JAPAN 3:CHINA]

	fmt.Println(len(m4)) // => 2

	// mapのループ
	m5 := map[string]int{
		"Apple":	100,
		"Banana":	200,
	}

	for k, v := range m5 {
		fmt.Println(k, v)
	} // => Apple 100, Banana 200

	for _, v := range m5 {
		fmt.Println(v)
	} // => 100, 200

	for k := range m5 {
		fmt.Println(k)
	}

	// チャネル
	// チャネルは、複数のゴルーチン間でデータの受け渡しを行うための仕組み
	var ch1 chan int
	// このままだと読み取り＋書き込みができない
	ch1 = make(chan int)
	// 読み取り＋書き込みができる

	// var ch2 <-chan int // 受信専用
	// var ch3 chan<- int // 送信専用

	ch2 := make(chan int)
	// バッファ付きチャネル
	fmt.Println(cap(ch1)) // => 0
	fmt.Println(cap(ch2)) // => 0
	// バッファ付きチャネルは、バッファのサイズを指定することができる
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3)) // => 5

	// チャネルの送受信
	ch3 <- 1 // 送信
	fmt.Println(len(ch3)) // => 1
	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3)) // => 3

	i3 := <-ch3 // 受信
	fmt.Println(i3) // => 1
	fmt.Println("len", len(ch3)) // => 2
	i4 := <-ch3
	fmt.Println(i4) // => 2
	fmt.Println("len", len(ch3)) // => 1
	// チャネルは受信すると要素が削除される（要素から取り出してる、キューのようなもの）

	fmt.Println(<-ch3) // => 3
	fmt.Println("len", len(ch3)) // => 0

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	// ch3 <- 6 // バッファがいっぱいの場合は、エラーになる
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch3) // => 1
	ch3 <- 6
	fmt.Println("len", len(ch3))
	fmt.Println(<-ch3) // => 2
	fmt.Println("len", len(ch3))

	// チャネルとゴルーチン
	// ch4 := make(chan int)
	// fmt.Println(<-ch4)
	// チャネルに何も入っていないと、エラーになる
	// fatal error: all goroutines are asleep - deadlock!

	// ch5 := make(chan int)

	// go receiver(ch4)
	// go receiver(ch5)

	// i5 := 0
	// for i5 < 100 {
	// 	ch4 <- i5
	// 	ch5 <- i5
	// 	time.Sleep(50 * time.Millisecond)
	// 	i5++
	// } // 0 0 1 1 2 2 3 3 4 4 5 5 6 6 7 7 8 8 9 9 ...

	// チャネルを明示的に閉じる
	ch6 := make(chan int, 2)
	close(ch6)
	// ch6 <- 1
	// panic: send on closed channel
	// fmt.Println(<-ch6) // => 1, 送信はできるが、受信はできない
	i, ok3 := <-ch6
	fmt.Println(i, ok3) // => 0 false

	ch7 := make(chan int, 2)
	ch7 <- 1
	close(ch7)
	i6, ok4 := <-ch7
	fmt.Println(i6, ok4) // => 1 true
	i7, ok5 := <-ch7
	fmt.Println(i7, ok5) // => 0 false

	// ch8 := make(chan int, 2)
	// go receiver2("Goroutine1", ch8)
	// go receiver2("Goroutine2", ch8)
	// go receiver2("Goroutine3", ch8)

	// i8 := 0
	// for i8 < 100 {
	// 	ch8 <- i8
	// 	i8++
	// }
	// close(ch8)
	// time.Sleep(3 * time.Second)
	// どのゴルーチンがどの順番で処理されるかは、実行環境によって異なる
	// 毎回異なる順番で処理される

	// チャネルのfor
	// ch9 := make(chan int, 3)
	// ch9 <- 1
	// ch9 <- 2
	// ch9 <- 3
	// for i := range ch9 {
	// 	fmt.Println(i)
	// } // => 1, 2, 3
	// fatal error: all goroutines are asleep - deadlock!

	ch10 := make(chan int, 3)
	ch10 <- 1
	ch10 <- 2
	ch10 <- 3
	close(ch10)
	for i := range ch10 {
		fmt.Println(i)
	} // => 1, 2, 3

	// チャネルのselect
	// ch11 := make(chan int, 2)
	// ch12 := make(chan string, 2)

	// ch12 <- "A"

	// // v1 := <-ch11 // ch11は空なのでエラーが走る
	// // // fatal error: all goroutines are asleep - deadlock!
	// // v2 := <-ch12
	// // fmt.Println(v1) // => 0
	// // fmt.Println(v2) // => A
	// // // チャネルが一つエラーが起きるだけでそれ以降の処理が止まる

	// // select文を使うと、上記の問題を解決できる
	// select {
	// 	case v1 := <-ch11:
	// 		fmt.Println(v1 + 1000)
	// 	case v2 := <-ch12:
	// 		fmt.Println(v2 + "!?")
	// }

	ch13 := make(chan int, 2)
	ch14 := make(chan string, 2)

	ch14 <- "A"
	ch13 <- 100
	ch14 <- "B"
	ch13 <- 200

	select {
		case v3 := <-ch13:
			fmt.Println(v3 + 1000)
		case v4 := <-ch14:
			fmt.Println(v4 + "!?")
	} // => 1100, A!?, どちらかが実行される

	ch15 := make(chan int, 2)
	ch16 := make(chan string, 2)
	select{
		case v5 := <-ch15:
			fmt.Println(v5 + 1000)
		case v6 := <-ch16:
			fmt.Println(v6 + "!?")
		default:
			fmt.Println("No Data")
	} // => No Data

	ch17 := make(chan int)
	ch18 := make(chan int)
	ch19 := make(chan int)

	// レシーバーの無名関数
	go func() {
		for {
			i9 := <-ch17
			ch18 <- i9 * 2
		}
	}()

	go func() {
		for {
			i10 := <-ch18
			ch19 <- i10 - 1
		}
	}()

	n2 := 0
	L:
		for {
			select {
				case ch17 <- n2:
					n2++
				case i11 := <-ch19:
					fmt.Println("received", i11)
				default:
					if n2 > 100 {
						break L // ラベルを指定してbreak
					}
			}
			// if n2 > 100 {
			// 	break
			// } // defaultじゃない場合はこちらでbreakできる
		}
}

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func receiver(c chan int){
	for {
		i := <-c
		fmt.Println(i)
	}
}

func receiver2(name string, ch chan int){
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}