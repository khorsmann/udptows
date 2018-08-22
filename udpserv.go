// udpserv
package main

import (
	"bytes"
	"log"
	"net"
)

// how can i use vars from client.go?
var (
	cr = []byte{'\n'}
	sp = []byte{' '}
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		log.Fatal("Couldn't send response: ", err)
	}
}

func udpserv(hub *Hub) {
	message := make([]byte, 256)
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal("UDP some error: ", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(message)
		log.Printf("Read a message from <%v> <%s> \n", remoteaddr, message)
		if err != nil {
			log.Fatal("UDP some error: ", err)
			continue
		}
		go sendResponse(ser, remoteaddr)

		// send message to hub as broadcast
		// Hi UDP Server, How are you doing?
		message = bytes.TrimSpace(bytes.Replace(message, cr, sp, -1))

		// seems to work but how can i trim "unused stuff" from udpclient?
		// with Trim all Null-Characters
		message = bytes.Trim(message, "\x00")
		hub.broadcast <- message

	}
}
