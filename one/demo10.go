package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/axgle/mahonia"
	"golang.org/x/net/html"
)

func main() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.booktxt.net/2_2219/356905707.html", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.146 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, " fetch :%v\v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	// b, err := ioutil.ReadAll(resp.Body)

	//解决中文乱码
	// bodystr := mahonia.NewDecoder("gbk").ConvertString(string(b))
	// fmt.Printf("%s", bodystr)

	// fmt.Println(bodystr)
	// 写入文件
	// ioutil.WriteFile("site.txt", bodystr, 0644)
	// 写入文件
	// data := []byte(bodystr)
	// ioutil.WriteFile("site.txt", data, 0644)

	body, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf(" html err %s", err)
	}
	// log.Fatalln(body)
	var content string
	getContent(&content, body)

	// var re = regexp.MustCompile(`.*?<a.*?href=".*?".*?>(.+?)</a>.*?`)
	var re = regexp.MustCompile(`.*?<div.*?id="content".*?>(.+?)</div>.*?`)
	aa := re.ReplaceAllString(content, "$1")

	var text []string
	for _, line := range strings.Split(string(aa), "<br/>") {
		if line != "\n" {
			line = strings.Trim(line, " ")
			line = strings.Trim(line, "<!---115--->")
			text = append(text, line)
		}

	}

	fmt.Println(text[2 : len(text)-3])
	// data := []byte(text)
	// ioutil.WriteFile("text1.txt", []byte(text), 0644)
	// fmt.Fprintf(os.Stderr, " fetch :%v\v", aa)
	// os.Exit(1)

	// fmt.Println(content)

}

func getContent(content *string, n *html.Node) {
	var buf bytes.Buffer
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, d := range n.Attr {
			if d.Key == "id" && d.Val == "content" {
				w := io.Writer(&buf)
				html.Render(w, n)
				*content = mahonia.NewDecoder("gbk").ConvertString(string(buf.String()))
			}

		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getContent(content, c)
	}
}
