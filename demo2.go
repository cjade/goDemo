package main

import "fmt"

type T struct {
	neme  string
	value int
}

func main() {
	var a = 1
	fmt.Printf("%5T \n", a)
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d \n", i)
	}
	fmt.Println("end")
}
