package main

import (
	"flag"
	"fmt"
	"os"
)

//go run demo1.go --help  该命令源码文件的参数说明

//go run demo1.go -name="te" -aa="dsd" 命令就可以为参数name,aa传值

var name string

var aaaaa *string

func init() {
	//自定义命令源码文件的参数使用说明,推荐使用这种方式
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	// 第一种获取命令行参数： key,默认值，说明
	aaaaa = flag.String("aa", "哈哈", "dfdf")

	// 第二种获取命令行参数：地址，key,默认值，说明
	flag.StringVar(&name, "name", "everyone", "The greeting object.")

}

func main() {
	// 自定义命令源码文件的参数使用说明,对flag.Usage的赋值必须在调用flag.Parse函数之前。
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	// 函数flag.Parse用于真正解析命令参数，并把它们的值赋给相应的变量。
	flag.Parse()
	fmt.Println("Hello, 世界")
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("Hello, %s!\n", *aaaaa)
}
