package main

import (
	//assume others to be present here
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

func stringifyBlocks(chainHead *a2.Block) string {
	var ret string
	ret += fmt.Sprintf("\n\n--------------------------Listing Blocks (most recent first) ... ---------------------\n")
	var currPtr = chainHead
	for currPtr != nil { //for block iteration
		ret += fmt.Sprintf("\n-----------------Block-----------------\n")
		// fmt.Println("Following are its transactions:-")
		for i := 0; i < len(currPtr.Data); i++ { //for iterations of transactions
			ret += fmt.Sprintf("Transaction %d : {Title:%s Sender:%s Receiver:%s Amount: %d} \n", (i + 1), currPtr.Data[i].Title, currPtr.Data[i].Sender, currPtr.Data[i].Receiver, currPtr.Data[i].Amount)
		}
		currPtr = currPtr.PrevPointer
	}
	ret += fmt.Sprintf("--------------------------------------------------------------------------------------\n\n ")
	return ret
}

func centralRoutine(clientCh chan Node) {
	//this is the area where block chain is formed,update and send to clients
	/*Complete the code below*/
	// invoking centralRoutine - being used for storing connections and handling clientChannel

	/*
		handling list of connection
		inform other clients about nodes
	*/
	clientsSlice := make([]Node, 0, 20)
	var chainHead *a2.Block
	chainHead = a2.PremineChain(chainHead, 2) //premine x number of blocks
	tempo := stringifyBlocks(chainHead)
	fmt.Printf(tempo)
	for {
		select {
		case newClient := <-clientCh:
			clientsSlice = append(clientsSlice, newClient)
			log.Println("new client has arrived")

			//adding its block in block chain
			curTitle := fmt.Sprintf("SatoshiTo%s", newClient.NodeID)
			GiftbySatoshi := []a2.BlockData{{Title: curTitle, Sender: "Satoshi", Receiver: newClient.NodeID, Amount: 10}}
			chainHead = a2.InsertBlock(GiftbySatoshi, chainHead)

			//send updated block chain to all
			updatedBL := stringifyBlocks(chainHead)
			//send to all
			for _, someClient := range clientsSlice {
				someClient.connection.Write([]byte(updatedBL))
			}
			log.Printf("new message:-\n %s", updatedBL)
		}
	}
}

func satoshiClientHandler(conn net.Conn, clientCh chan Node) {
	log.Println("Receiving node ID from the node")
	/*Complete the code below*/
	//please enter your node id
	//recv node id
	//send through channel to main sub routine

	//- invoking satoshiClientHandler - being used for handling each client node
	enterID := make([]byte, 120)
	conn.Write([]byte("Please Enter Your NodeId:"))
	n, _ := conn.Read(enterID)
	ClientNode := Node{connection: conn, NodeID: string(enterID[0:n])}
	clientCh <- ClientNode

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
	var clientChannel = make(chan Node)
	ln, err := net.Listen("tcp", satoshiAddress)
	if err != nil {
		log.Fatal(err)
	}
	go centralRoutine(clientChannel)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go satoshiClientHandler(conn, clientChannel)
	}
}
