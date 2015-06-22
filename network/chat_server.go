package main

import (
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	var clients []net.Conn
	input := make(chan []byte, 10)

	go func() {
		for {
			message := <-input
			for _, client := range clients {
				client.Write(message)
			}
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		clients = append(clients, conn)
		go handleClient(conn, input)
	}
}

func handleClient(client net.Conn, input chan []byte) {
	for {
		buf := make([]byte, 4096)
		nbytes, err := client.Read(buf)
		if nbytes == 0 || err != nil {
			return
		}
		input <- buf
	}
}
