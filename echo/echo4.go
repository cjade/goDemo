package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 打印命令行每个参数和值

// go run echo/echo4.go name=tom age=18



func main() {
	startTime := getUnixMillisecond()
	for index, arg := range os.Args[1:] {
		command := strings.Split(arg, "=")
		fmt.Println(index, command[0], command[1])
	}
	endTime := getUnixMillisecond()

	time := endTime - startTime

	fmt.Println("for耗时", time)

	joinStartTime := getUnixMillisecond()
	fmt.Println(strings.Join(os.Args[1:], " "))
	joinEndTime := getUnixMillisecond()
	joinTime := joinEndTime - joinStartTime
	fmt.Println("join耗时", joinTime)
}
