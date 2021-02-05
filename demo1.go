package main

import (
	"flag"
	"fmt"
)

//go run demo1.go -name="te" -aa="dsd"

// 第一种获取命令行参数： key,默认值，说明
var aa = flag.String("aa", "哈哈", "dfdf")

var name string

func init() {
	// 第二种获取命令行参数：地址，key,默认值，说明
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	fmt.Println("Hello, 世界")
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("Hello, %s!\n", *aa)
}
