// Unix里echo命令的实现
// 打印命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// 程序的命令行参数可从os包的Args变量获取
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
