// udpclient.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"time"
)

var uaddr = flag.String("uaddr", "127.0.0.1:29000", "udp target service address")

func main() {
	flag.Parse()
	p := make([]byte, 256)

	conn, err := net.Dial("udp", *uaddr)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	message := fmt.Sprintf("Hi UDP Server, How are you doing? %v", time.Now().Format(time.RFC850))
	// fmt.Fprintf(conn, "Hi UDP Server, How are you doing?\n")
	fmt.Fprintf(conn, message)
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
