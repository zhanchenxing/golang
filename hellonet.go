package main

import (
	"fmt"
	"net"
	"log"
)

const listenAddr = "localhost:4000"

func main(){
	l, err := net.Listen( "tcp", listenAddr )
	if err != nil {
		fmt.Println("Listen failed:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint( c,  "hello!" )
		c.Close()
	}
}
