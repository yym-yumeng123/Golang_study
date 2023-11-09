### 文件

文件是数据源的一种, 最主要的作用就是保存数据

文件在程序中以`流`的形式操作

Go程序(内存) -> 输出流(写文件) -> 文件
文件 -> 输入流(读文件) -> Go程序

- 流: 数据在数据源和程序之间经历的路径
- 输入流: 数据从数据源(文件)到程序(内存)的路径
- 输出流: 数据从程序(内存)到数据源(文件)的路径

#### 基本介绍

`os.File` 封装了所有文件相关操作, File是一个结构体

- File 代表一个打开的文件对象
- `Create` 创建一个文件
- `Open` 打开一个文件

常用的文件操作函数

1. 打开一个文件, `os.Open(name string) (*File, error)`
2. 关闭一个文件 `File.close()`

读文件操作应用实例

1. 读取文件的内容并显示在终端(带缓冲区的方式), 使用 `os.Open, file.Close bufio.NewReader(), reader.ReadString`函数和方法
2. 读取文件显示在终端(使用`os`一次性整个文件读入到内存中), 适用于文件不大的情况
   - `os.ReadFile`

写文件操作应用实例

`OpenFile(name string, flag int, perm FileMode) (file *File, err error)`

- `os.OpenFile` 是一个更一般性的文件打开函数
- flag 参数: 文件打开模式 只读/只写/读写等
- perm 参数: 权限控制 linux/unix 下面用

判断文件是否存在

golang 判断文件或文件夹是否存在的方法为使用 `os.Stat()`

- 返回的错误为 nil, 文件或文件夹存在
- 返回的错误类型使用 `os.IsNotExit()` 为true, 说明不存在
- 返回的错误为其它类型,不确定是否存在

```go
func PathExists(path string) (bool, error) {
	_, err := os.Stat()
	if err == nil {
		return  true, nil
	}
	if os.IsNotExit(err) {
		return false, nil
}
	return false err
}
```

拷贝文件

`io.Copy(dst Writer, src Reader) (writte int64, err error)`


---


### 命令行参数

基本介绍

`os.Args` 是一个string 的切片 , 用来存储所有的命令行, 比较原生, 解析参数不是特别的方便, 特别是带有指定参数形式的命令行

go 另外提供了 `flag` 包, 可以方便的解析命令行参数, 参数顺序可以随意


### json 基本介绍

一种轻量级的数据交换格式, 可以序列化和反序列化

JSON 键值对用来保存数据, 都用双引号包裹

1. web编程中的应用, 数据传输使用 json
2. tcp编程中的应用, 例如Go写了聊天系统


结构体 map 切片的序列化

```go
// 序列化: 有key-value结构的数据类型(结构体 map, 切片) 序列化成 json 字符串的操作

dataMap, _ := json.Marshal(map)

// 反序列化: 将json字符串反序列化成对应的数据类型
// 保证反序列化后的数据类型和原来序列化前的数据类型一致
err := json.Unmarshal([]byte(str), &monster)
```























