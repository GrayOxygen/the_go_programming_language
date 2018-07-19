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
	mustCopy(os.Stdout, conn)
}
func mustCopy(dst io.Writer, src io.ReadCloser) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
