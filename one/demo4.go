package main

import "fmt"

type s string

func main() {
	var as s
	as = "dad"
	fmt.Println(as.String())
}

func (c s) String() string {
	return fmt.Sprintf("%g", c)
}
