package main

import (
	//	"bufio"
	//	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(client net.Conn) {
	for {
		buf := make([]byte, 4096)
		nbytes, err := client.Read(buf)
		if nbytes == 0 || err != nil {
			return
		}
		client.Write(buf)
	}
}
