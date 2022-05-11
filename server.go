package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

//func process(conn net.Conn) {
//	defer conn.Close()
//	for {
//		reader := bufio.NewReader(conn)
//		buffer := make([]byte, 1024)
//		n, err := reader.Read(buffer[:])
//		if err != nil {
//			fmt.Println("read from client failed, err: ", err)
//			break
//		}
//		str := string(buffer[:n])
//		if str == "q" {
//			conn.Write([]byte("bye!")) // 发送数据
//			break
//		}
//		fmt.Println("收到Client端发来的数据：", str)
//		conn.Write([]byte(str)) // 发送数据
//	}
//}

func main() {
	listen, err := net.Listen("tcp", ":8081")
	ServerHandleError(err, "net.listen")
	for {
		conn, err := listen.Accept()
		ServerHandleError(err, "listen Accept")
		go ChatWith(conn)
	}
}

func ChatWith(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		ServerHandleError(err, "conn read")
		clientMsg := strings.Trim(string(buffer[0:n]), "\r\n")
		fmt.Print(clientMsg)
		fmt.Printf("收到%v消息:%s\n\n", conn.RemoteAddr(), clientMsg)
		if clientMsg == "offline" {
			conn.Write([]byte("OK off line!"))
			break
		}
		conn.Write([]byte("已阅:" + clientMsg))
	}
	fmt.Printf("客户端%v下线\n", conn.RemoteAddr())
}

func ServerHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
