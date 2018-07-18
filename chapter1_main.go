package main

import (
	"./chapters"
	"fmt"
)

func main() {
	//chapter1_123()
	//chapter1_4()
	chapter1_567()

}

func chapter1_123() {
	fmt.Println("====================================第一章练习====================================")

	//打印命令行参数
	//go run chapter1_main.go -id=2 -name="wanghuiyang" -age=12 -detals="细节介绍"
	chapters.Echo()
	chapters.Echo2()
	chapters.Echo3()
	chapters.EchoIndexVal()

	//统计重复行
	chapters.Dup() //命令行输入
	//go run chapter1_main.go  ./chapter1重复行.txt
	chapters.DupFromFile()  //统计文件内容或者命令行内容的重复行
	chapters.DupFromFile2() //一次性读入大块内存来统计
}

func chapter1_4() {
	fmt.Println("====================================第一章练习====================================")
	// 执行命令 go run chapter1_main.go web 打开了8000端口，web服务，访问即可看到动画
	chapters.Paint()
}
func chapter1_567() {
	fmt.Println("====================================第一章练习====================================")
	//go run chapter1_main.go http://www.google.com
	//chapters.Fetch()
	//chapters.Fetch2()
	chapters.Fetch3()

	//go run chapter1_main.go http://www.sina.com http://www.hao123.com http://www.baidu.com
	chapters.Fetch4() //并发请求多个url
}
