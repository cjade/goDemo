package main

import (
	"fmt"
	"os"
)

// 优化

func main() {
	// var s, sep string
	s, sep := "", " "

	for _, arg := range os.Args[1:] {
		s += arg + sep
		// sep = " "
	}
	fmt.Println(s)
}
