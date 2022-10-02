package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(">>>Error<<<", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(">>>Error<<<", err)
		}
		for {
			var input string
			fmt.Print("Scanning >>> ")
			fmt.Scanln(&input)
			conn.Write([]byte(input))
			buffer := make([]byte, 1024)
			mLen, err := conn.Read(buffer)
			if err != nil {
				fmt.Println(">>>Error<<<", err)
			}
			fmt.Println("Received", string(buffer[:mLen]))
		}
	}
}

