package main

import (
	"log"
	"net"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleListener(conn)
	}
}

func handleListener(conn net.Conn) {
	defer conn.Close()
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	t := time.Now()
	msg := t.Format(time.RFC1123)
	conn.Write([]byte(msg))
}
