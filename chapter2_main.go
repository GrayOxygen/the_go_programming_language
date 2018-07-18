package main

import (
	"fmt"
	"the_go_programming_language/chapters"
)

func main() {
	//typepoint()
	//go run chapter2_main.go -n -s 分割 shi ni hello 00 yeah
	//chapters.Echo4()

	fmt.Println(chapters.GCD(5, 3))

}

//变量指针练习
func typepoint() {
	var x, y int
	//指针类型的零值不为nil
	fmt.Println(&x == &x, &x == &y, &x == nil)
	//函数返回的局部变量是非常安全的，f的v即使在调用后依然存在，指针p依然引用它，每次调用f都会返回一个不同的值
	fmt.Println(f() == f())
	res1 := f()
	res2 := f()
	fmt.Println(res1 == res2, res1, res2)

}

var p = f()
var v int

func f() *int {
	v := 1
	return &v
}

//值传递，间接更新值，指针不变
func incr(p *int) int {
	*p++
	return *p
}
