package main

import (
	"net"
	"log"
	"os"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8999")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy2(os.Stdout, conn)
	//输入流赋值到连接，使得能够在客户端启动后能够输入
	mustCopy2(conn, os.Stdin)
}

func mustCopy2(dst io.Writer, src io.ReadCloser) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
