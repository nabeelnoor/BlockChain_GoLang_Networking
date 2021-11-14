//nodes.go
package main

import (
	"fmt"
	"net"
	//assume others to be present here
	// a2 "github.com/ehteshamz/assignment02IBC"
)

func main() {
	satoshiAddress := ":2020"
	/* Add code below to connect to Satoshi */
	callAddress := fmt.Sprintf("localhost%s", satoshiAddress)
	conn, err := net.Dial("tcp", callAddress)

	if err != nil {
		// handle error
	}

	/* Add code below to send NodeID to Satoshi */
	recvdSlice := make([]byte, 30)
	conn.Read(recvdSlice)
	fmt.Print(string(recvdSlice))
	var buffer string
	fmt.Scanln(&buffer)
	conn.Write([]byte(buffer))

	/*Provide code below to receive and print the chain */
	recvdSlice2 := make([]byte, 500)
	for {
		fmt.Println("In Output Phase")
		conn.Read(recvdSlice2)
		fmt.Println(string(recvdSlice2))
	}

	select {}
}
