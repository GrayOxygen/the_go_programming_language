package main

import (
	"fmt"
	"math"
	"unicode/utf8"
	"the_go_programming_language/chapters"
	"reflect"
)

func main() {
	////chapter3_123()
	//fmt.Println(chapters.CommaNotRecursion("ahao我们好niyeyao好起来123456"))
	////字符串和整形互相转换
	//x := 123
	//y := fmt.Sprintf("%d", x)
	//fmt.Println(y, reflect.TypeOf(y), strconv.Itoa(x))
	//z := fmt.Sprintf("x的二进制为%b", 2)
	//xx, xxerr := strconv.Atoi("123abc")
	//yy, yyerr := strconv.ParseInt("123", 10, 64)
	//fmt.Println(z, "---", xx, "---", xxerr, "---", yy, "---", yyerr)

	///
	//ptr := [2]byte{1, 2}
	//zero(&ptr)
	//
	//months := [...]string{1: "Jan"}
	//fmt.Println(months)
	//
	////无法获取map元素，因为不是一个变量，&persons["wdx"]会报错，主要是rehash时，bucket存的值可能变化
	//persons := make(map[string]person)
	//persons["wdx"] = person{Name: "王大仙"}
	//persons["xm"] = person{Name: "小明"}
	//fmt.Println(persons["wdx"], reflect.TypeOf(persons["wdx"]))
	//var students map[string]int
	//fmt.Println(students == nil, len(students) == 0)
	////报错，因为nil map没有任何键值
	////students["wang"] = 100
	//val, ok := students["x"]
	////不存在的键，查出的值就是零值
	//fmt.Println(val, ok)
	//
	//fmt.Println(utf8.UTFMax)
	var typeList []int64
	typeList = nil
	fmt.Println(typeList, len(typeList))
}

type person struct {
	Name string
}

func zero(ptr *[2]byte) {
	fmt.Println(reflect.TypeOf(ptr), ptr[1])
	for i := range ptr {
		ptr[i] = 0
	}
}
func chapter3_123() {
	//golang中%的结果以被除数的符号为准
	fmt.Println(5 % -3)
	var apples int32 = 1
	var oranges int16 = 2
	var compote int = int(apples) + int(oranges)
	fmt.Println(compote)
	var z float64
	fmt.Println(z/z == math.NaN())

	//rune表示字符或者码点，也是int32类型的别名
	var name string = "David王"
	fmt.Println(len(name), utf8.RuneCountInString(name))
	//UTF-8编码的字符串的码点序列
	chineseStr := "中文"
	fmt.Printf("十六进制% x\n", chineseStr)
	r := []rune(chineseStr)
	fmt.Printf("unicode码十六进制%x\n", r)
	//根据数值转字符串
	fmt.Println(string(1234567))
	fmt.Println(string(0x4eac))

	chapters.Comma("123abc我们都是大哥吹")
	//字符串和字节可以互相转换
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s, b, s2)

}
