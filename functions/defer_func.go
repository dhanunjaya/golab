// This Program is from Effective go and It talks the usage of defer

package main

import (
	"fmt"
	"io" 
	"os"
)

func Contents(filename string)(string, error) {
		f, err := os.Open(filename)
		if err != nil {
				return "", err
		}
		defer f.Close()
		var result []byte
		buf := make([]byte, 100)
		for {
				nbytes, err := f.Read(buf[0:])
				result = append(result, buf[0:nbytes]...)
				if err != nil {
						if err == io.EOF {
								break
						}
						return "", err
				}
		}
		return string(result), nil
}

func main() {
		contents, _ := Contents("/etc/hosts")
		fmt.Println(contents)
}
