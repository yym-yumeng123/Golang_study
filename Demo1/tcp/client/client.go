package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("client dial err", err)
		return
	}
	fmt.Println(conn)
	// 功能1: 客户端发送单行数据, 然后退出
	// os.stdio 代表标准输入[终端]
	v := bufio.NewReader(os.Stdin)
	// 从终端读取一行用户输入, 发送服务器
	con, err := v.ReadString('\n')
	if err != nil {
		fmt.Println("readString err=", err)
	}
	// con 发送服务器
	n, err := conn.Write([]byte(con))
	if err != nil {
		fmt.Println("conn,Write err=", err)
	}
	fmt.Printf("客户端发生了 %d 字节的数据", n)

}
