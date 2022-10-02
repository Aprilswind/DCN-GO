package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(">>>Error<<<", err)
	}

	for {
		buffer := make([]byte, 1024)
		mLen, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(">>>Error<<<", err)
		}
		fmt.Println("Received", string(buffer[:mLen]))
		var input string
		fmt.Print("Scanning >>> ")
		fmt.Scanln(&input)
		conn.Write([]byte(input))
	}

}
