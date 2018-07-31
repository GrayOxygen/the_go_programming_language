package chapters

import (
	"net"
	"log"
	"io"
	"time"
	"bufio"
	"fmt"
	"strings"
)

func Clock1() {
	listener, err := net.Listen("tcp", "localhost:8999")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn1(conn)
	}
}

func handleConn1(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

//并发回声服务器
func ServerEcho() {
	listener, err := net.Listen("tcp", "localhost:8999")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn2(conn)
	}
}
func handleConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func PipeLine() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for x := 0; x < 6; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	go func() {
		//channel没有数据时，自动退出循环
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	for {
		fmt.Println(<-squares)
	}
}

//单向通道，双向通道可转换为单向通道，反过来不行
func PipeLine2() {
	naturals := make(chan int)
	squarers := make(chan int)
	go counter(naturals)
	go squarer(squarers, naturals)
	printer(squarers)
}
func counter(out chan<- int) {
	for x := 0; x < 11; x++ {
		out <- x
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func CountDown1() {
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
}
