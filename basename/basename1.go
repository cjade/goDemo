package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KiB
)

// a.go => a, a/b/c.go => c, a/b.c.go => b.c
func main() {
	s := "e.g., a"
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	fmt.Println(s)
	a := 1 << 20
	fmt.Println(a)

	aa := [32]byte{1, 2, 3}

	zero(&aa)
	fmt.Println(aa)

}

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
