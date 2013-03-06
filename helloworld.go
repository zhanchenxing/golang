package main

import "fmt"
import "time"
import "os"
import "hash/crc32"
import "io"
import "encoding/hex"

type World struct {
}

func (w *World) String() string {
	return "世界"
	return "world"
}

type ByteSize float64
const (
	_ = iota 		// ignore first value
	KB ByteSize = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	return fmt.Sprintf("%v", int(b) )
}

func main(){
	fmt.Println("Hello, world!")
	fmt.Printf("hello, %s\n", "世界" )
	fmt.Printf("hello, %s\n", new ( World )  )

	fmt.Printf("KB=%v, MB=%v, GB=%v\n", KB, MB, GB )

	fmt.Printf("KB=%v\n", KB )

	var bs ByteSize
	bs = 2048
	fmt.Printf("bs=%v\n", bs )


	start := time.Now()
	// fetch( "http://www.google.com.hk" )
	time.Sleep( time.Millisecond * 3 )
	fmt.Println( time.Since(start) )

	fmt.Fprintf(os.Stdout, "%s\n", "Hello, world from os.Stdout " )
	fmt.Printf("%s\n", "hello, world")

	

	h := crc32.NewIEEE()

	fmt.Fprintf(h, "Hello, world\n" )
	fmt.Printf("hash=%#x\n", h.Sum32() )

	file, err := os.Create("./output.txt")
	if err != nil {
		fmt.Printf("Open file failed: %v", err )
		return
	}
	defer file.Close()

	multiWriter := io.MultiWriter( file, os.Stdout )
	fmt.Fprintf( multiWriter, "hello, %s\n", "world")

	dumper := hex.Dumper( os.Stdout )
	defer dumper.Close()
	fmt.Fprintf( dumper , "%s", "hello, world!")
	
}
