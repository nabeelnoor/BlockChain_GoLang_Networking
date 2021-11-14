package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		// handle error
	}
	recvdSlice := make([]byte, 11)
	conn.Read(recvdSlice)
	fmt.Println(string(recvdSlice))

}
