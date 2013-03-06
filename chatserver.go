package main

import (
	"net"
	"io"
	"log"
	"fmt"
)

const listenAddr = "localhost:4000"

/*
func match( c net.Conn ){
	written, err := io.Copy( c, c )
	fmt.Println("written:", written, "error:", err)
	c.Close()
}
*/

var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser) {
	fmt.Fprint(c, "Waiting for a partner...")
	select {
	case partner <- c:
		// now handled by the other goroutine
	case p := <-partner:
		chat(p, c)
	}
}


func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Found one! Say hi.")
	fmt.Fprintln(b, "Found one! Say hi.")
	go io.Copy(a, b)
	io.Copy(b, a)
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
		go match( c )
	}

}
