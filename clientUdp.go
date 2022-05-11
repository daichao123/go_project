package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("UDP", ":8888")
	clientHandleError(err, "net dial")
	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)
	for {
		//从标准输入中读取消息
		line, _, _ := reader.ReadLine()
		conn.Write(line)
		n, _ := conn.Read(buffer)
		if string(buffer[0:n]) == "ok offline" {
			conn.Write([]byte("嘻嘻嘻"))
			break
		}
		fmt.Printf("收到服务端消息%s\n", string(buffer[0:n]))

	}
}

func clientHandleError(err error, when string) {
	if err != nil {
		fmt.Printf("发生错误%s", when)
		os.Exit(1)
	}
}
