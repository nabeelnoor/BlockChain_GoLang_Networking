package main

import (
	"log"
	"net"
	"fmt"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
	fmt.Println("asdasd")
}

func handleConnection(c net.Conn) {

	log.Println("A client has connected", c.RemoteAddr())
	c.Write([]byte("Hello world"))
	fmt.Println("handle executed")

}
