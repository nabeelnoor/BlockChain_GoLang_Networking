package main

import (
	"fmt"
	"net"
)

func Send(conn net.Conn) {
	var buffer string
	for {
		fmt.Println("In Input Phase")
		fmt.Scanln(&buffer)
		conn.Write([]byte(buffer))
	}
}

func Recv(conn net.Conn) {
	recvdSlice := make([]byte, 30)
	for {
		fmt.Println("In Output Phase")
		conn.Read(recvdSlice)
		fmt.Println(string(recvdSlice))
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		// handle error
	}
	recvdSlice := make([]byte, 30)
	conn.Read(recvdSlice)
	fmt.Print(string(recvdSlice))
	var buffer string
	fmt.Scanln(&buffer)
	conn.Write([]byte(buffer))

	go Send(conn)
	go Recv(conn)
	for {
	}

}
