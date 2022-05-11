package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("UDP", ":8888")
	handleError(err, "net ResolveUDPAddr")
	conn, err := net.ListenUDP("UDP", addr)
	buffer := make([]byte, 1024)
	for {
		n, udpAddr, err := conn.ReadFromUDP(buffer)
		handleError(err, "conn.ReadFromUDP")
		s := string(buffer[0:n])
		if s == "offline" {
			conn.WriteToUDP([]byte("ok offline"), udpAddr)
			break
		}
		fmt.Printf("收取来自：%v的消息：%s", udpAddr, s)
		conn.WriteToUDP([]byte("收到消息"), udpAddr)
	}

}

func handleError(err error, when string) {
	if err != nil {
		fmt.Printf("发生错误%s", when)
		os.Exit(1)
	}
}
