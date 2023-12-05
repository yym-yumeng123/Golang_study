### 创建 "Hello World"

```go
// main.go 全小写字母形式命名, .go 扩展名为结尾

// 定义包名 
// 整个 Go 程序中仅允许存在一个名为 main 的包
package main

// 导入包
import "fmt"

// main 函数
// 当你运行一个可执行的 Go 程序的时候，所有的代码都会从这个入口函数开始运行
func main() {
	// fmt 输出到终端
	fmt.Println("Hello World")
}
```

```shell
# 运行 main.go
go run main.go

# 编译 main.go, 获得一个二进制可执行文件
go build main.go
./main.go
```

### Go Module

Go 语言的构建模式历经了三个迭代和演化过程，分别是最初期的 `GOPATH`、1.5 版本的 `Vendor` 机制，以及现在的 `Go Module`

`Go module` 构建模式是在 Go 1.11 版本正式引入的

Go 1.16 版本中，Go module 已经成为了 Go 默认的包依赖管理机制和 Go 源码构建机制

Go Module 的核心是一个名为 go.mod 的文件，在这个文件中存储了这个 module 对第三方依赖的全部信息

- 一个 Go Module 是一个 Go 包的集合. module 是有版本的
- 在 Go Module 模式下, 通常一个代码仓库对应一个 Go Module, 一个 Go Module 的顶层目录下会放置一个 go.mod 文件，每个 go.mod 文件会定义唯一一个 module
- go.mod 文件所在的顶层目录也被称为 module 的根目录

```shell
# 使用 go mod init 在当前目录创建一个 go.mod 文件
# go.mod 文件将当前项目变为一个 Go Module
go mod init blog  # 用前端类比, npm init -y, 创建一个 package.json

# go mod tidy 命令会扫描 Go 源码，并自动找出项目依赖的外部 Go Module 以及版本，
# 下载这些依赖并更新本地的 go.mod 文件
# 生成一个 go.sum 的文件记录直接依赖和间接依赖包的相关版本的 hash 值
go mod tidy # 前端 => npm install

# 特殊情况, 在 vendor 目录下，创建了一份这个项目的依赖包的副本
go mod vendor
```

```go
// go.mod
module blog

go 1.21.0

require github.com/sirupsen/logrus v1.8.1
```

上面我用了前端类比, 是为了更方便的理解, 当然了 go 和 node 包的管理方式完全不同

1. 为当前 module 添加一个依赖

```shell
# go get 命令将我们新增的依赖包下载到了本地 module 缓存里，
# 并在 go.mod 文件的 require 段中新增了一行内容
go get github.com/google/uuid
```

2. 升级/降级依赖的版本

Go Module 的版本号采用了语义版本规范，也就是版本号使用 vX.Y.Z 的格式。
其中 X 是主版本号，Y 为次版本号（minor），Z 为补丁版本号（patch）

以 `logrus` 为例

```shell
# 查询版本
go list -m -versions github.com/sirupsen/logrus

# 执行带有版本号的 go get 命令
# 选择 v1.7.0 版本的
go get github.com/sirupsen/logrus@v1.7.0

# 使用 go mod tidy
# 1. 先试用 go mode edit 明确告知依赖的版本
go mod edit -require=github.com/sirupsen/logrus@v1.7.0
go mod tidy


# 升级
go get github.com/sirupsen/logrus@v1.7.1
```

3. 添加一个主版本号大于 1 的依赖

如果新旧版本的包使用相同的导入路径，那么新包与旧包是兼容的

如果我们要为 Go 项目添加主版本号大于 1 的依赖，
我们就需要使用“语义导入版本”机制，在声明它的导入路径的基础上，加上版本号信息

```go
package main

import (
  _ "github.com/go-redis/redis/v7" // “_”为空导入
  "github.com/google/uuid"
  "github.com/sirupsen/logrus"
)

func main() {
  logrus.Println("hello, go module mode")
  logrus.Println(uuid.NewString())
}
```

```shell
# 加上版本号信息 v7
go get github.com/go-redis/redis/v7
```

4. 升级依赖版本到一个不兼容版本

语义导入版本的原则，不同主版本的包的导入路径是不同的

```go
import (
  _ "github.com/go-redis/redis/v8"
  "github.com/google/uuid"
  "github.com/sirupsen/logrus"
)
```

```shell
go get github.com/go-redis/redis/v8
```

5. 移除一个依赖

```shell
# 列出当前 module 的所有依赖
go list -m all

# 代码中删除 import 导入的依赖
# 自动分析源码依赖，而且将不再使用的依赖从 go.mod 和 go.sum 中移除
go mod tidy
```

### 入口函数和包初始化

1. main.main 函数: Go应用入口函数

```go
// 可执行程序的 main 包必须定义 main 函数，否则 Go 编译器会报错
package main

func main() {
    // 用户层执行逻辑
    ... ...
}
```

2. init 函数: Go包的初始化函数

进行包初始化的 `init` 函数

- 重置包级变量值
- 实现对包级变量的复杂初始化
- 在 init 函数实现 "注册模式"

```go
// init 函数的执行就都会发生在 main 函数之前
// init 函数也是一个无参数无返回值的函数
func init() {
    // 包初始化逻辑
    ... ...
}
```

Go 包可以拥有不止一个 init 函数，
每个组成 Go 包的 Go 源文件中，也可以定义多个 init 函数

先传递给 Go 编译器的源文件中的 init 函数，会先被执行；
而同一个源文件中的多个 init 函数，会按声明顺序依次执行

3. Go 包的初始化次序

- 依赖包按“深度优先”的次序进行初始化；
- 每个包内按以“常量 -> 变量 -> init 函数”的顺序进行初始化；
- 包内的多个 init 函数按出现次序进行自动调用