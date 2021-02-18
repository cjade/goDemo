package main

import "fmt"

func main() {
	f := square
	fmt.Printf("%d", f) // "9"
}

func square(n int) int { return n * n }
