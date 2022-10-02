package main

import (
	"fmt"
	"net"
)

func handleInput (conn net.Conn) {
	var input string
	fmt.Scanln(&input)
	conn.Write([]byte(input))
}

func handleRead (conn net.Conn) {
	buffer := make([]byte, 4096)
	mLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Received " + string(buffer[:mLen]))
}

func handleConnection (conn net.Conn) {

	for {

		go handleInput(conn)
		go handleRead(conn)
	
	}

}

func main () {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		
		go handleConnection(conn)
			
	}
}