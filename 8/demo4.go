package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	ch := make(chan string, 3)
	ch <- "A"
	// ch <- "B"
	ch <- "C"
	fmt.Println("----------")
	fmt.Println(cap(ch))
	fmt.Println(len(ch))
	fmt.Println("----------")
	close(ch)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
	fmt.Println("----------")
	for a := range ch {
		fmt.Println(a)
	}

}

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}
func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
