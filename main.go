package main

import (
	"./chapters"
	"fmt"
)

func main() {
	//chapter1_123()
	//chapter1_4()
	chapter1_5()
}
func chapter1_123() {
	fmt.Println("====================================第一章练习====================================")

	//打印命令行参数
	//go run main.go -id=2 -name="wanghuiyang" -age=12 -detals="细节介绍"
	chapters.Echo()
	chapters.Echo2()
	chapters.Echo3()
	chapters.EchoIndexVal()

	//统计重复行
	chapters.Dup() //命令行输入
	//go run main.go  ./chapter1重复行.txt
	chapters.DupFromFile()  //统计文件内容或者命令行内容的重复行
	chapters.DupFromFile2() //一次性读入大块内存来统计
}

func chapter1_4() {
	fmt.Println("====================================第一章练习====================================")
	// 执行命令 go run main.go web 打开了8000端口，web服务，访问即可看到动画
	chapters.Paint()
}
func chapter1_5() {
	fmt.Println("====================================第一章练习====================================")
	//go run main.go http://www.google.com
	//chapters.Fetch()
	chapters.Fetch2()
}
