package main

import (
	"fmt"
	"log"
	"runtime"
)

func rt_string(arg1 string, arg2 string) string {
	return arg1 + arg2
}

func exec_for(arg int) int {
	sum := 0
	for count := 0; count < arg; count++ {
		sum += count
		fmt.Println(sum)
	}
	return sum
}

func exec_if(arg1 int, arg2 int) int {
	var judge = 1
	if arg1%2 == judge {
		return arg2
	} else {
		return arg1
	}
}

func exec_switch(arg string) string {
	switch arg {
	case "JS":
		return "This is JS"
	case "Python":
		return "This is Python"
	case "Go":
		return "Yes!! This is GoLang"
	default:
		return "I don't know what you talking about"
	}
}

func exec_defer() {
	defer fmt.Println("Golang") //defer1
	defer fmt.Println("Ruby")   //defer2
	fmt.Println("JS")
}

func exec_pointer() {
	name := "Nakamura"
	fmt.Println(&name)

	namePoint := &name
	fmt.Println(namePoint)
	fmt.Println(*namePoint)
}

type Person struct {
	name string
	age  int
}

type User struct {
	Person
	company string
}

func exec_struct() {
	var Nakamura Person
	Nakamura.name = "AOI"
	Nakamura.age = 22

	var sato = Person{name: "taro", age: 23}
	fmt.Println(Nakamura, sato)

	tanaka := User{
		Person: Person{
			name: "tanaka",
			age:  24,
		},
		company: "Go",
	}
	fmt.Println(tanaka)

}

func exec_arrays() {
	var arr1 [2]string
	var arr2 [2]string = [2]string{"Go", "Python"}
	var arr3 = [...]string{"Go", "JS"}
	fmt.Println(arr1, arr2, arr3)
	arr1[0] = "C"
	arr1[1] = "Java"
	fmt.Println(arr1, arr2, arr3)
}

func exec_slice() {
	// 参照用の配列
	var refer_array [3]string = [3]string{"Go", "Python", "JS"}
	fmt.Println(refer_array)

	// スライス　<-可変長の配列的な? リスト?
	// 配列とは異なり, 長さの指定なし
	var slice1 []string
	var slice2 []string = []string{"Python", "JS"}
	fmt.Println(slice1, slice2)

	// 配列から要素を取り出し参照する形での宣言
	var slice3 = refer_array[0:1]
	fmt.Println(slice3)

	// make()の利用
	var slice4 = make([]string, 2, 2)
	fmt.Println(slice4)

	// 配列とは異なり要素の追加が可能
	slice5 := []string{"JavaScript"}
	new_slice5 := append(slice5, "TypeScript")
	fmt.Println(slice5, new_slice5)

}

func exec_maps() {
	// 連想配列
	// makeでの宣言
	map1 := make(map[string]string)
	map1["name"] = "Nakamura"
	map1["Gen"] = "male"
	fmt.Println(map1)

	// 初期値を指定して宣言
	map2 := map[string]int{"age": 22, "userId": 1}
	fmt.Println(map2)

	var map3 = map[int]string{1: "Nakamura", 2: "sato"}
	fmt.Println(map3)
	map3[3] = "tanaka"
	fmt.Println(map3)
	delete(map3, 2)
	fmt.Println(map3)

}

func exec_range() {
	//スライスとマップの作成
	var slice1 = []string{"Golang", "Ruby"}
	var map1 = map[string]string{"Lang1": "Golang", "Lang2": "Ruby"}

	for index, value := range slice1 {
		fmt.Println(index, value)
	}

	for key, value := range map1 {
		fmt.Println(key, value)
	}

	for _, value := range map1 {
		fmt.Println(value)
	}

}

type character struct {
	name        string
	yearOfBirth int
	birthPlace  string
}

type performance struct {
	character
	drama string
}

func (p performance) profile() {
	fmt.Println(p.name, "\n出生年:", p.yearOfBirth, "\n出身地:", p.birthPlace, "\n出演ドラマ:", p.drama)
}

func exec_method() {
	p1 := performance{
		character: character{
			"新垣結衣",
			1988,
			"沖縄県",
		},
		drama: "逃げるは恥だが役に立つ",
	}

	p2 := performance{
		character: character{
			"本田翼",
			1992,
			"東京都",
		},
		drama: "ラジエーションハウス",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	p1.profile()
	p2.profile()
}

// ここからインタフェースの内容
type Stringfy interface {
	ToString() string
}

type Book struct {
	name   string
	birth  int
	author string
}

func (book *Book) ToString() string {
	return fmt.Sprintf("%s(%d歳:%s)", book.name, book.birth, book.author)
}

type Animal struct {
	name string
	kind string
}

func (animal *Animal) ToString() string {
	return fmt.Sprintf("%s[%s]", animal.name, animal.kind)
}

func exec_interface() {
	vs := []Stringfy{
		&Book{
			name:   "Nakamura",
			birth:  2021,
			author: "Aoi",
		},
		&Animal{
			name: "rabbit",
			kind: "哺乳類",
		},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

}

// インタフェース 補足情報
// レシーバについて, 値レシーバとポインタレシーバ
type Users struct {
	Name string
}

// メソッド（値レシーバ）
// func (u Users) setName() {
// 	fmt.Printf("u address is %p\n", &u)
// 	u.Name = "gopher"
// }

// メソッド（値レシーバ）
func (u Users) setUserNameValue() {
	u.Name = "gopherValue"
}

// メソッド（ポインタレシーバ）
func (u *Users) setUserNamePointer() {
	u.Name = "gopherPointer"
}

func exec_receiver() {
	u1 := new(Users)      //uは*User型
	u1.setUserNameValue() //暗黙的に(*u).setName()と解釈される。
	fmt.Println(u1.Name)

	u2 := Users{}           //uはUser型
	u2.setUserNamePointer() //暗黙的に(&u).setName()と解釈される。
	fmt.Println(u2.Name)
}

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("World")
}

func exec_go_routine() {
	go hello()
	go world()
	// mainとhelloとworldのルーチンが動いている
	log.Println(runtime.NumGoroutine())
}

func goRoutine1(vs []int, ch chan int) {
	sum := 0
	for _, v := range vs {
		sum += v
	}
	ch <- sum
}

func goRoutine2(vs []int, ch chan int) {
	sum := 1
	for _, v := range vs {
		sum = sum * v
	}
	ch <- sum
}

func exec_channel() {
	vs := []int{1, 2, 3, 4, 5}
	ch := make(chan int, 2)

	log.Println(runtime.NumGoroutine())

	go goRoutine1(vs, ch)
	log.Println(runtime.NumGoroutine())
	p_1 := <-ch
	go goRoutine2(vs, ch)
	log.Println(runtime.NumGoroutine())
	p_2 := <-ch

	fmt.Println(p_1, p_2)

}

const NUM = 22

func main() {
	// fmt.Println("Hello golang from docker!")
	var printString string
	printString = rt_string("Hello ", "world")
	fmt.Println(printString)
	// fmt.Println(NUM)
	// fmt.Println(exec_for(8))
	// fmt.Println(exec_if(3, 5))
	// fmt.Println(exec_switch("Python"))
	// fmt.Println(exec_switch("Go"))
	// fmt.Println(exec_switch("go"))
	// exec_defer()
	// exec_pointer()
	// exec_struct()
	// exec_arrays()
	// exec_slice()
	// exec_maps()
	// exec_range()
	// exec_method()
	// exec_interface()
	// exec_receiver()
	// exec_go_routine()
	exec_channel()
}
