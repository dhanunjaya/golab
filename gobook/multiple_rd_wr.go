package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func getPageSize(url string) (int, error) {
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

func main() {
	urls := []string{"http://www.google.com", 
				"https://news.ycombinator.com", 
				"http://www.cricbuzz.com"}
	size := 0
	var err error
	for _, url := range urls {
		size, err = getPageSize(url)
		if err != nil {
			fmt.Println("%s\n", err)
		}
		fmt.Printf("%s %d\n", url, size)
	}
}
