package main

import (
	//assume others to be present here
	"encoding/gob"
	"fmt"
	"log"
	"net"

	a2 "github.com/nabeelnoor/assignment02IBC"
)

type Node struct {
	NodeID string
	/*Complete the code below */
	connection net.Conn
}

func centralRoutine(clientCh chan Node) {
	/*
		Complete the code below
		invoking centralRoutine - being used for storing connections and handling clientChannel
		this is the area where block chain is formed,update and send to clients
		handling list of connection
		inform other clients about nodes
	*/
	clientsSlice := make([]Node, 0, 20) //slice to store information of clients that connect to satoshi
	//creation of blockchain and premine x number of blocks, let x=2
	var chainHead *a2.Block
	chainHead = a2.PremineChain(chainHead, 2) //premine x number of blocks

	//here block chain is update(adding of transaction) and send to all clients whenever new client is connected
	for {
		select {
		case newClient := <-clientCh: //if there is new client arrived
			clientsSlice = append(clientsSlice, newClient) //append new client information
			log.Println("new client has arrived")

			//adding transaction for new client and adding its block in block chain
			curTitle := fmt.Sprintf("SatoshiTo%s", newClient.NodeID)
			GiftbySatoshi := []a2.BlockData{{Title: curTitle, Sender: "Satoshi", Receiver: newClient.NodeID, Amount: 10}}
			chainHead = a2.InsertBlock(GiftbySatoshi, chainHead)

			//send updated block chain to all clients
			for _, someClient := range clientsSlice {
				enc := gob.NewEncoder(someClient.connection)
				enc.Encode(chainHead)
				//someClient.connection.Write([]byte(updatedBL))
			}
			a2.ListBlocks(chainHead)
		}
	}
}

func satoshiClientHandler(conn net.Conn, clientCh chan Node) {
	log.Println("Receiving node ID from the node")
	/*	//- invoking satoshiClientHandler - being used for handling each client node
		Complete the code below
		1.please enter your node id
		2.recv node id
		3.send through channel to centralRoutine
	*/

	//Asking nodeID from node
	enterID := make([]byte, 120)
	conn.Write([]byte("Please Enter Your NodeId:"))
	n, _ := conn.Read(enterID)                                         //Recieving that nodeID
	ClientNode := Node{connection: conn, NodeID: string(enterID[0:n])} //making node struct for new client
	clientCh <- ClientNode                                             //sending to "Node struct" to centralRoutine through channel

	/*just reading from node (not required in question) as client only send nodeid and then recv
	blockchain on occassion of updates
	*/
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			conn.Close()
			break
		}
	}
	log.Println("Client Expired")
}

/* Update the code below (write on the right side) to add any missing information including
- any global variables, do NOT write to them from multiple goroutines
- invoking centralRoutine - being used for storing connections and handling clientChannel
- invoking satoshiClientHandler - being used for handling each client node
- any other info needed ...
*/

func main() {
	satoshiAddress := ":2020"
	var clientChannel = make(chan Node)          //making channel to communicate between sub routines
	ln, err := net.Listen("tcp", satoshiAddress) //listening at port 2020
	if err != nil {
		log.Fatal(err)
	}
	go centralRoutine(clientChannel) //invoking centralRoutine (thread that run in parallel)
	for {
		conn, err := ln.Accept() //accepting dial from nodes
		if err != nil {
			log.Println(err)
			continue
		}
		go satoshiClientHandler(conn, clientChannel) //invoking satoshiClientHandler (thread that run in parallel)
	}
}
