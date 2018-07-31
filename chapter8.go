package main

import (
	"flag"
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

func main() {
	//确定初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	//遍历文件树
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
}

//递归求得目录下所有文件的文件大小
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

//返回目录中的条目：所有直接子文件（和直接子文件夹）
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}

//打印大小
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB \n", nfiles, float64(nbytes)/1e9)
}
