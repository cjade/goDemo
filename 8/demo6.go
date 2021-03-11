package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fileSizes := make(chan int64)
	walkDir(dir, fileSizes)
	fmt.Println(fileSizes)
	close(fileSizes)
	// for f := range fileSizes {
	// 	fmt.Printf("f: %v s: ", f)
	// }

}

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

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stdout, "du1: %v \n", err)
		return nil
	}
	return entries
}
