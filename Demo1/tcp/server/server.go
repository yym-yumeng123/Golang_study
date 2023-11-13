package main

import (
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	// 循环接收客户端发送的数据
	defer conn.Close() // 关闭 conn
	for {
		// 创建每一个新的切片
		buf := make([]byte, 1024)
		// 等待客户端通过 conn 发送信息
		fmt.Println("服务器在等待客户端的输入")
		n, err := conn.Read(buf) // 从 conn读取
		if err != io.EOF {
			fmt.Printf("客户端已退出")
			return
		}

		// 显示到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	// 网络开发, net 包 tcp 协议
	// 使用的是网络协议 tcp
	// 本地监听 8080 端口
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		// handle error
		fmt.Println("listen err=", err)
	}

	defer ln.Close()

	for {
		// 等待客户端链接
		fmt.Println("等待客户端链接...")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept()", err)
		} else {
			fmt.Println("Accept suc con", conn, conn.RemoteAddr())
		}

		go handleConnection(conn)
	}

}
