package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

var aa *string

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// 第一种获取命令行参数： key,默认值，说明
	aa = flag.String("aa", "哈哈", "dfdf")

	// 第二种获取命令行参数：地址，key,默认值，说明
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	cmdLine.Parse(os.Args[1:])
	fmt.Println("dsd")
}
