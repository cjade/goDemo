package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 获取当前时间戳毫秒
func getUnixMillisecond() int64 {
	// return time.Now().UnixNano() / int64(time.Millisecond)
	// 纳秒
	return time.Now().UnixNano()
}

func main() {
	startTime := getUnixMillisecond()
	s, sep := "", " "

	for _, arg := range os.Args[1:] {
		s += arg + sep
		// sep = " "
	}
	fmt.Println(s)
	endTime := getUnixMillisecond()

	time := endTime - startTime

	fmt.Println("for耗时", time)

	joinStartTime := getUnixMillisecond()
	fmt.Println(strings.Join(os.Args[1:], " "))
	joinEndTime := getUnixMillisecond()
	joinTime := joinEndTime - joinStartTime
	fmt.Println("join耗时", joinTime)
}
