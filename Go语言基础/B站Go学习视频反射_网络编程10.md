### 反射

1. 反射可以在运行时动态获取变量的各种信息, 比如变量的类型(type), 类别(kind)
2. 如果时结构体变量, 还可以获取到结构体本身的信息(字段, 方法)
3. 通过反射, 可以修改变量的值, 可以调用关联的方法
4. 使用反射, import("reflect)

反射重要函数

1. `reflect.TypeOf(变量名)` 获取变量类型, 返回 `reflect.Type` 类型
2. `reflect.ValueOf(变量名)` 获取变量的值, 返回 `reflect.Value`类型, 是一个结构体类型
3. `变量, interface{} reflect.Value` 可以相互转换


```go
var stu Student
var num int

// 用于做反射
func test(b interface{}){
	// 如何将interface{} 转成 reflect.Value
	rVal := reflect.ValueOf(b)
	// 类型转回去
	iVal := rVal.Interface()
	// 如何将 interface{} 转成 原来的变量类型
	// 类型断言
	v := iVal.(Student)
}

test(stu)
```

注意事项和细节

1. `reflect.value.Kind` 获取变量的类别, 返回的是一个常量
2. Type 是类型, Kind 是类别, 可能相同, 可能也不同
3. 通过反射可以让变量在 interface{} 和 Reflect.Value 之间相互转换
4. 使用反射的方式来获取变量的值, 要求数据类型匹配, 比如 x 是 int, 应该使用 `reflect.Value().Int()`
5. 通过反射来修改变量, 注意当使用 `SetXXX`方法来设置需要通过对应指针类型来完成, 这样才能改变传入的值, 同时需要 `reflect.Value.Elem()` 方法
6. `reflect.Value.Elem()`

```go
func main() {
	// 通过反射修改值
	var num int = 10
	reflect1(&num)

	fmt.Println(num)
}

func reflect1(b interface{}) {
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(20)
}

```

反射最佳实践

1. 使用反射来遍历结构体的字段, 调用结构体的方法, 并获取结构体标签的值
   - `func (v Value) Method(i int) Value` 默认按方法名排序对应的 i 值, 从 0 开始
   - `func (v Value) Call(in []Value) []Value` 传入参数和返回参数是 []reflect.Value

---

### 网络编程 TCP

Golang 的主要设计目标之一就是面向大规模后端服务程序. 网络通信这块是服务端程序不可少的一部分


网络编程两种:

1. TCP socket编程: 底层是基于 `Tcp/ip协议`, 比如: QQ 聊天
2. `b/s`结构是 `http编程`, 使用浏览器去访问服务器时, 使用的是 `http协议`, http底层依旧是用`tcp socket`实现的


基础知识

**协议TCP/IP**

- TCP Transmission Control protocol

TCP/IP 网络通信协议, 简单的说: 就是由网络层IP协议和传输层的TCP协议组成的,
Tcp/ip模型

1. 应用层 (application) smtp ftp telnet http App
2. 传输层 (transport) 解释数据 Tcp头
3. 网络层 (ip) 定位ip 地址和确定链接路径 Ip头
4. 链路层 (link) 与硬件驱动对话 帧头 | 桢尾


IP地址

每个 internet 上的主机和路由器都有一个 ip 地址, 它包括网络号和主机号
ip地址有 ipv4(32位) `4个字节` 或者ipv6(128位), 可以通过 ipconfig 查看

端口 prot 

指TCP/IP协议中的端口, 是逻辑意义上的端口. 如果把ip地址比作房子, 端口就是门,

- 0 是保留端口
- 1-1024是固定端口, 又叫有名端口, 一般程序员不能用
- 1025 - 65535 是动态端口

1. 只要做服务程序, 必须监听一个端口
2. 该端口就是其他程序和该服务通讯的通道
3. 一台电脑有 65535 个端口 1 - 66535
4. 一旦端口被程序监听, 其他程序不能监听该端口
5. 尽量少开端口, 一个端口只能被一个程序监听
6. 使用 `netstat -an` 查看本机有那些端口在监听
7. 使用 `netstat -anb` 查看端口的 pid, 结合任务管理器关闭不安全的端口



### tcp socket 编程的客户端和服务端

服务端的处理流程

1. 监听端口 8888
2. 接受客户端的 tcp 链接, 建立客户端和服务端的链接
3. 创建 `goroutine`, 处理该链接的请求

客户端的处理流程

1. 建立与服务端的链接
2. 发送请求数据, 接受服务器端返回的数据
3. 关闭链接
























