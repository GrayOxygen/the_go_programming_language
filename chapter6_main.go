package main

import (
	"fmt"
	"the_go_programming_language/chapters"
)

func main() {
	var p chapters.Point
	p = [2]int{1, 2}
	fmt.Println((&p).Where())

}
