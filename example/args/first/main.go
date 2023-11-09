package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 获取命令行参数
	fmt.Println("命令行的参数有", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("args[%v]=%v", i, arg)
	}

	// 定义变量, 用于接受命令行参数
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名, 默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码, 默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名")
	flag.IntVar(&port, "port", 3306, "默认为3306")

	// 转换
	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
