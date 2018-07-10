package chapters

import (
	"flag"
	"fmt"
	"strings"
)

//指针练习
//-n标识忽略输出时结尾的换行符，-s标识用sep替换默认参数输出时使用的空格分割符
func Echo4() {
	var n = flag.Bool("n", false, "忽略结尾处的换行")
	var sep = flag.String("s", ",", "分割符")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

//new创建一个未命名的T类型变量，返回地址，地址类型为*T
//每次new返回的地址不同，即每次new到的变量都是具有唯一地址的不同变量
func newInt() *int {
	//等价于new创建
	//var dummy int
	//return &dummy
	return new(int)
}

//多重赋值，求最大公约数
func GCD(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
