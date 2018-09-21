package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("failed to accept: %v\n", err)
			continue
		}
		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	defer conn.Close()

	remote, err := net.Dial("tcp", "gophercon.com:443")
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()

	// client to remote
	go io.Copy(remote, conn)

	// remote to client
	io.Copy(conn, remote)
}
