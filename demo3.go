package main

import (
	"fmt"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	var time = time.Now().Unix()
	var sum = add(12, 15)

	a, b := swap("hello", "world")

	fmt.Printf("a = %s  b = %s \n", a, b)

	fmt.Printf("现在时间戳 %d \n", sum)
	fmt.Printf("现在时间戳 %d \n", time)
}
