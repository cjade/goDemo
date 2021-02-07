// 并发的获取多个url
//这个程序会同时去获取所有的URL，所以这个程序的总执行时间不会超过执行时间最长的那一个任务
// go run fetch/fetch2.go http://books.studygolang.com/gopl-zh/ch1/ch1-05.html http://books.studygolang.com/gopl-zh/ch1/ch1-06.html http://gopl.io

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	// 把响应的Body内容拷贝到ioutil.Discard输出流中
	// 可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
