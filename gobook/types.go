package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	err  error
	body []byte
}

func (w *webPage) getPage() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()
	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
		return
	}
}

func (w *webPage) isOk() bool {
	return w.err == nil
}

func main() {
	w := &webPage{url: "http://www.google.com"}
	w.getPage()
	if w.isOk() {
		fmt.Printf("url: %s, error: %s, length: %d\n", w.url, w.err, len(w.body))
	} else {
		fmt.Println("something went wrong")
	}
}
