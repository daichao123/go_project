package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8081")
	ClientHandleError(err, "net Dial")
	defer conn.Close() // 关闭TCP连接
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, err := reader.ReadString('\n')
		inputInfo := strings.Trim(readString, "\r\n")
		if err != nil {
			log.Fatal(err)
		}
		// 发送数据
		conn.Write([]byte(inputInfo))
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		ClientHandleError(err, "conn read")
		if string(buf[:]) == "OK off line!" {
			break
		}
		fmt.Println(string(buf[:n]))
	}
	fmt.Println("test over")
}

func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
