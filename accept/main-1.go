package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	log.Println("Start Server")
	l, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// io.Copy is blocking to write
		n, err := io.Copy(os.Stderr, conn)
		log.Printf("Copied %d, err = %v", n, err)
	}
}
