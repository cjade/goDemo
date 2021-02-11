package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

func main() {

	start := time.Now()

	ch := make(chan []string)
	for _, url := range os.Args[1:] {
		go findLinks(url, ch)
	}
	for range os.Args[1:] {
		links := <-ch
		for _, val := range links {
			fmt.Println(val)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func findLinks(url string, ch chan<- []string) {
	start := time.Now()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.146 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {
		ch <- []string{fmt.Sprint(err)}
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		ch <- []string{fmt.Sprintf("get url: %s StatusCode: %d", url, resp.StatusCode)}
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		ch <- []string{fmt.Sprintf("html Parse err %s", err)}
		return
	}

	ch <- visit1(nil, doc)
	fmt.Printf("%.2fs elapsedsss\n", time.Since(start).Seconds())
	return
}

func visit1(links []string, d *html.Node) []string {
	if d.Type == html.ElementNode && d.Data == "a" {
		for _, a := range d.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := d.FirstChild; c != nil; c = c.NextSibling {
		links = visit1(links, c)
	}
	return links
}
