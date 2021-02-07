//我们会展示一个微型服务器，这个服务器的功能是返回当前用户正在访问的URL
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handel)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
