package main

import (
	"the_go_programming_language/chapters"
)

func main() {
	//chapters.Clock1()
	//chapters.ServerEcho()
	//chapters.PipeLine()
	//chapters.PipeLine2()
	//in := make(chan string)
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		in <- "go1" + strconv.Itoa(i)
	//	}()
	//}
	//go func() {
	//	for i := 51; i < 100; i++ {
	//		in <- "go2" + strconv.Itoa(i)
	//	}
	//}()
	//go func() {
	//	for i := 101; i < 222; i++ {
	//		in <- "go3" + strconv.Itoa(i)
	//	}
	//}()
	//go func() {
	//	for x := range in {
	//		fmt.Println("输出：：：", x)
	//	}
	//	close(in)
	//}()
	//time.Sleep(10000)
	chapters.CountDown1()
}
