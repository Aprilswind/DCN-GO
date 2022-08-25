package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 1, 103),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("connection failed!", err)
		return
	}
	// defer socket.Close()
	senddata := []byte("hello server!")
	_, err = socket.Write(senddata)
	if err != nil {
		fmt.Println("send data failed!", err)
		return
	}
	data := make([]byte, 4096)
	read, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read data failed!", err)
		return
	}
	fmt.Println(read, remoteAddr)
	fmt.Printf("%s\n", data)
}
