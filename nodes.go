//nodes.go
package main

import (
	"encoding/gob"
	"fmt"
	"net"

	//assume others to be present here
	a2 "github.com/nabeelnoor/assignment02IBC"
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
	conn.Read(recvdSlice) //reading "asking of nodeid from satoshi"
	fmt.Print(string(recvdSlice))
	var buffer string
	fmt.Scanln(&buffer)        //input of nodeid
	conn.Write([]byte(buffer)) //sending nodeid to Satoshi

	/*Provide code below to receive and print the chain */
	//recv blockChain from satoshi
	var chainHead *a2.Block
	dec := gob.NewDecoder(conn) //making of decoder
	for {
		fmt.Println("In Output Phase")
		dec.Decode(&chainHead) //recieving blockchain
		a2.ListBlocks(chainHead)
	}

	select {}
}
