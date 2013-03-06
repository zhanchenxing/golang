package main

import (
	"fmt"
	"net"
	"log"
	"io"
	"os"
)


func reading ( c net.Conn, q chan int ){
	written, err := io.Copy( os.Stdout, c )
	fmt.Println("reading exit: written=", written, ", err=", err )
	q <- 1
}

func writting( c net.Conn, q chan int ){
	count, err := io.Copy( c, os.Stdin )
	fmt.Println("writting exit: count=", count, ", err=", err )
	q <- 1
}

func main(){
	quit := make( chan int )
	c, err := net.Dial("tcp", "localhost:4000" )
	if err != nil {
		log.Fatal(err)
	}

	go reading( c, quit )
	go writting( c, quit )

	<-quit
	<-quit
	fmt.Println("quit")
}



