package main

import (
	// "bufio"
	"log"
	"net"
	"fmt"
	//"o«s"
)

//Client is something
type Client struct {
	connection net.Conn
	nickname   string
}

func handleConnection(c net.Conn, msgchan chan string, addchan chan Client) {
	nickName := make([]byte, 120)
	c.Write([]byte("Please Enter Your NickName:\n"))
	c.Read(nickName)
	//clientReader := bufio.NewReader(c)
	//nickName, _, _ := clientReader.ReadLine()

	newClient := Client{connection: c, nickname: string(nickName) + ": "}
	fmt.Println("Sending new client in channel")
	addchan <- newClient

	buf := make([]byte, 4096)
	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		msgchan <- newClient.nickname + string(buf[0:n])
		// ...
	}
	log.Println("Client Expired")
}

func printMessages(msgchan chan string, addchan chan Client) {
	clientsSlice := make([]Client, 0, 20)
	for {
		select {
		case newClient := <-addchan:
			clientsSlice = append(clientsSlice, newClient)
			log.Println("new client has arrived")

		case msg := <-msgchan:
			for _, someClient := range clientsSlice {
				someClient.connection.Write([]byte(msg))
			}
			log.Printf("new message: %s", msg)
		}
	}
}

func main() {

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	addchan := make(chan Client) //making of channel
	msgchan := make(chan string)
	go printMessages(msgchan, addchan) //making of subroutine that print msg

	for { //this loop handles new connection
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn, msgchan, addchan)
	}
}
