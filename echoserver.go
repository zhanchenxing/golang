package main

import (
	"net"
	"io"
	"log"
	"fmt"
)

const listenAddr = "localhost:4000"

func echo( c net.Conn ){
	written, err := io.Copy( c, c )
	fmt.Println("written:", written, "error:", err)
	c.Close()
}

func main(){
	l, err := net.Listen("tcp", listenAddr )
	if err != nil {
		log.Fatalf("Listen failed: %v", err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal("Accept failed:", err)
		}
		go echo( c )
	}

}
