package main

import (
	"time"
	"sort"
	"fmt"
	"net/http"
	"io"
	"os"
	"bytes"
	"strings"
	"github.com/go-errors/errors"
)

type dollars float32
type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s:  %s\n", item, price)
	}
}

func main() {
	//db := database{"shoes": 90, "socks": 50}
	//mux := http.NewServeMux()
	//mux.Handle("/hello", http.HandlerFunc(db.list))
	//log.Fatal(http.ListenAndServe("localhost:8080", mux))
	//类型断言
	var w io.Writer
	w = os.Stdout
	f, ok1 := w.(*os.File)
	c, ok2 := w.(*bytes.Buffer)
	fmt.Println(f, ok1)
	fmt.Println(c, ok2)
	//使用类型断言识别错误
	err := errors.Errorf("%s", "错误：：：file does not exist")
	fmt.Println(IsNotExist(err))
}
func IsNotExist(err error) bool {
	return strings.Contains(err.Error(), "file does not exist")
}
func test() {
	names := []string{"王", "辉", "阳", "wang", "王"}
	sort.Sort(StringInterface(names))
	fmt.Println(names)
	var time = length("4m40s")
	fmt.Println(time)

}
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type StringInterface []string

func (p StringInterface) Len() int {
	return len(p)
}

func (p StringInterface) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p StringInterface) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
