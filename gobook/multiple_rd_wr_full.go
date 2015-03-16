package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPage(url string) (size int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	return len(body), nil
}

func worker(urlCh chan string, sizeCh chan string) {
	url := <-urlCh
	length, err := getPage(url)
	if err != nil {
		sizeCh <- fmt.Sprintf("%s has length %d\n", url, length)
	} else {
		sizeCh <- fmt.Sprintf("%s has length %d\n", url, length)
	}
}

func generator(url string, urlCh chan string) {
	urlCh <- url
}

func main() {
	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.bing.com"}
	urlCh := make(chan string)
	sizeCh := make(chan string)

	for _, value := range urls {
		go generator(value, urlCh)
	}

	for i := 0; i < 10; i++ {
		go worker(urlCh, sizeCh)
		fmt.Printf("%s\n", <-sizeCh)
	}
}
