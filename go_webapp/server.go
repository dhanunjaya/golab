package main

import (
	"fmt"
	"net/http"
//	"strings"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("Scheme-URL: ", r.URL.Scheme)
	fmt.Println("User-URL: ", r.URL.Host)
	fmt.Println("Path-URL: ", r.URL.Path)
	fmt.Println("Protocol: ", r.Proto)
	fmt.Println("ProtoMajor: ", r.ProtoMajor)
	fmt.Println("ProtoMinor: ", r.ProtoMinor)
	r.ParseForm()
	fmt.Println("Form map[string][]string: ", r.Form)
	fmt.Fprintf(w, "Hello Dhanunjaya Naidu Ravada!\n")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
