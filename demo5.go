// 作用域

package main

import (
	"fmt"
	"log"
	"os"
)

func f() {}

var g = "g"

var cwd string

func init() {
	// 虽然cwd在外部已经声明过，但是:=语句还是将cwd和err重新声明为新的局部变量。因为内部声明的cwd将屏蔽外部的声明，因此下面的代码并不会正确更新包级声明的cwd变量。
	cwd, err := os.Getwd() // compile error: unused: cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	//由于当前的编译器会检测到局部声明的cwd并没有使用，然后报告这可能是一个错误，但是这种检测并不可靠。因为一些小的代码变更，例如增加一个局部cwd的打印语句，就可能导致这种检测失效。
	// log.Printf("Working directory = %s", cwd)
}

func main() {
	// 内部声明屏蔽了外部同名的声明，让外部的声明的名字无法被访问
	f := "f"

	fmt.Println(f)

	fmt.Println(g) // "g"; 包级别变量
	// fmt.Println(h) // compile error: undefined: h

	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}

}
