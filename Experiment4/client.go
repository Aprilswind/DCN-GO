package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	serverIp := "127.0.0.1"
	serverPort := "8888"

	conn, err := net.Dial("tcp", serverIp+":"+serverPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	readBuf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	readLen, err := conn.Read(readBuf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(readBuf[:readLen]))
}
