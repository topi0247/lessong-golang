package lesson10

import (
	"fmt"
)

// 構造体
type User struct {
	Name string
	Age int
	// X, Y int // 複数のフィールドを持つこともできる
}

type T struct {
	User User
}

type T2 struct {
	User
}

type Users []*User

/*
	type Users struct{
		Users []*User
	}
	上記はtype Users []*Userと同じ
*/

type MyInt int // 独自型

func (mi MyInt) Print(){
	fmt.Println(mi)
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 100
}

func UpdateUser2(user *User) {
	user.Name = "B"
	user.Age = 200
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func NewUser(name string, age int) *User{
	return &User{Name: name, Age: age}
}

func Lesson10() {
	var user1 User
	fmt.Println(user1) // => { 0}
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1) // => {user1 10}

	user2 := User{}
	fmt.Println(user2) // => { 0}
	user2.Name = "user2"
	user2.Age = 20
	fmt.Println(user2) // => {user2 20}

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3) // => {user3 30}

	user4 := User{"user4", 40}
	fmt.Println(user4) // => {user4 40}

	// user5 := User{50, "user5"}
	// fmt.Println(user5) // => エラー
	// cannot use 50 (untyped int constant) as string value in struct literal
	// cannot use "user5" (untyped string constant) as int value in struct literal

	user6 := User{Name: "user6"}
	fmt.Println(user6) // => {user6 0}

	user7 := new(User) // ポインタ型
	fmt.Println(user7) // => &{ 0}

	user8 := &User{}
	fmt.Println(user8) // => &{ 0}

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1) // => {user1 10}
	fmt.Println(user8) // => &{B 200}

	user9 := User{Name: "user9"}
	user9.SayName() // => user9
	user9.SetName("user99")
	user9.SayName() // => user9
	user9.SetName2("user999")
	user9.SayName() // => user999

	user10 := &User{Name: "user10"}
	user10.SetName2("user100")
	user10.SayName() // => N

	// 構造体の埋込
	t := T{User: User{Name: "user", Age: 10}}
	fmt.Println(t) // => {{user 0}}
	fmt.Println(t.User) // => {user 0}
	fmt.Println(t.User.Name) // => user
	// fmt.Println(t.Name) // => エラー
	// t.Name undefined (type T has no field or method Name)

	t2 := T2{User: User{Name: "user", Age: 10}}
	fmt.Println(t2) // => {{user 0}}
	fmt.Println(t2.User) // => {user 0}
	fmt.Println(t2.User.Name) // => user
	fmt.Println(t2.Name) // => user

	t.User.SetName2("user2")
	fmt.Println(t) // => {{user2 0}}

	// struct型コンストラクタ
	user11 := NewUser("user11", 11)
	fmt.Println(user11) // => &{user11 11}
	fmt.Println(*user11) // => {user11 11}

	// structスライス
	user12 := User{Name: "user12", Age: 12}
	user13 := User{Name: "user13", Age: 13}
	user14 := User{Name: "user14", Age: 14}
	user15 := User{Name: "user15", Age: 15}

	users := Users{}
	users = append(users, &user12)
	users = append(users, &user13)
	users = append(users, &user14, &user15)

	fmt.Println(users) // => [0xc0000b2000 0xc0000b2040 0xc0000b2080 0xc0000b20c0]

	for _, user := range users {
		fmt.Println(*user)
	} // => {user12 12} {user13 13} {user14 14} {user15 15}

	users2 := make([]*User, 0)
	users2 = append(users2, &user12, &user13, &user14, &user15)
	for _, user := range users2 {
		fmt.Println(*user)
	} // => {user12 12} {user13 13} {user14 14} {user15 15}

	// structのマップ
	m := map[int]User{
		1: User{Name: "user1", Age: 1},
		2: User{Name: "user2", Age: 2},
	}

	fmt.Println(m) // => map[1:{user1 1} 2:{user2 2}]

	m2 := map[User]string {
		{Name: "user1", Age: 1}: "Tokyo",
		{Name: "user2", Age: 2}: "LA",
	}
	fmt.Println(m2) // => map[{user1 1}:Tokyo {user2 2}:LA]

	m3 := make(map[int]User)
	fmt.Println(m3) // => map[]
	m3[1] = User{Name: "user1", Age: 1}
	fmt.Println(m3) // => map[1:{user1 1}]

	for _, user := range m {
		fmt.Println(user)
	} // => {user1 1} {user2 2}

	// struct 独自型
	var mi MyInt
	fmt.Println(mi) // => 0
	fmt.Printf("%T\n", mi) // => lesson10.MyInt

	mi.Print() // => 0
}