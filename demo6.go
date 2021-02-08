// 字符串

package main

import "fmt"

func main() {
	s := "let`s go go go"

	fmt.Println(len(s))

	fmt.Println(s)
	s = "dsad"
	fmt.Println(s[0:len(s)])

	n := 0
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		n++
	}

	fmt.Println(n)
}
