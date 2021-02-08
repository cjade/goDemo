package main

import (
	"fmt"
)

func main() {
	// a := []int{1, 2, 3, 4, 5, 6}
	// reverse(a) // [6 5 4 3 2 1]
	// reverse(a[:2])  // [2 1 3 4 5 6]
	// reverse(a[2:]) // [1 2 6 5 4 3]
	// reverse(a[2:4]) // [1 2 4 3 5 6]
	// fmt.Println(a)

	// var s []int
	// s = []int{}
	// fmt.Println(len(s) == 0)
	// asd()

	// ma()

	// data := []string{"one", "", "three"}
	// fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	// fmt.Printf("%q\n", data)           // `["one" "three" "three"]`

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"

}

// 要删除slice中间的某个元素并保存原有的元素顺序，可以通过内置的copy函数将后面的子slice向前依次移动
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 去除空字符
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {

		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func ma() {
	var ma = make([]int, 5, 9)
	ma[0] = 3
	fmt.Println(len(ma))
	fmt.Println(cap(ma))
	fmt.Println(ma)
}

// 翻转
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}

func asd() {
	x := []int{1, 2, 3}
	var z []int
	for i := 0; i < 10; i++ {
		z = appendInt(x, i)
		fmt.Printf("%d cap = %d \t %v \n", i, cap(z), z)
		x = z
	}

}

// 向slice追加元素 y->x
func appendInt(x []int, y int) []int {
	z := []int{}
	xl := len(x)
	l := xl + 1
	c := cap(x)
	if l <= c {
		z = x[:l]
	} else {
		zcap := l
		// 分配两倍长度的容量
		if zcap < 2*xl {
			zcap = 2 * xl
		}
		z = make([]int, l, zcap)
		copy(z, x)
	}
	z[xl] = y
	return z
}
