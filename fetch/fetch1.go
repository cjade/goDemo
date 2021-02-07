// 打印命令行输入的url内容
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, " fetch :%v\v", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		// 第一种
		// b, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }

		// fmt.Printf("%s", b)

		// 第二种
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, " io.Copy :%v\v", err)
			os.Exit(1)
		}

	}
}
