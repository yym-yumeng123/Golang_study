### 版本规范

主版本号.次版本号.修订号（X.Y.Z），其中 X、Y 和 Z 为非负的整数，且禁止在数字前方补零

- 主版本号（MAJOR）：当做了不兼容的 API 修改
  - 主版本号为零（0.y.z）的软件处于开发初始阶段，一切都可能随时被改变，这样的公共 API 不应该被视为稳定版
- 次版本号（MINOR）：当做了向下兼容的功能性新增及修改。这里有个不成文的约定需要你注意，偶数为稳定版本，奇数为开发版本
  - 次版本号 Y（x.Y.z | x > 0）必须在有向下兼容的新功能出现时递增，在任何公共 API 的功能被标记为弃用时也必须递增，当有改进时也可以递增
- 修订号（PATCH）：当做了向下兼容的问题修正
  - 修订号 Z（x.y.Z | x > 0）必须在只做了向下兼容的修正时才递增，这里的修正其实就是 Bug 修复。

```go
v1.2.3

// 先行版本号
1.0.0-alpha
1.0.0-alpha.1
1.0.0-0.3.7
1.0.0-x.7.z.92
```

### commit 规范

| 类型       | 类别   | 说明      |
|----------|------|---------|
| feat     | prod | 新增功能    |
| fix      | prod | bug修复   |
| pref     | prod | 提高代码性能  |
| style    | dev  | 格式类的变更  |
| refactor | prod | 其它代码类变更 |
| test     | dev  | 测试      |
| docs     | dev  | 文档      |

**合并提交**

- `git rebase` 重写历史
  - `git rebase -i commitId` 进入交互页面

| 命令        | 目的                          |
|-----------|-----------------------------|
| p,pick    | 不做任何处理                      |
| r, reword | 保留该commit,但修改提交信息           |
| e, edit   | 保留, 但rebase暂停,允许修改这个commit  |
| s, squash | 保留, 但会将当前commit与上一个commit合并 |
| f, fixup  | 与上面相同,不会保存当前commit提交信息      |
| x, exec   | 执行其它shell命令                 |
| d, drop   | 删除该 commit                  |

- 操作示例
  - 新建一个分支, 例如 branchA
  - 在 branchA, 合并所有commit `git rebase -i commitId`
  - 可以将 branch 合并到 master `git checkout master; git merge branchA`

**commit message**

- `git commit --amend`：修改最近一次 commit 的 message
- `git rebase -i`：修改某次 commit 的 message。


### 目录规范

- 命名清晰
- 功能明确
- 全面性
- 可观测性
- 可扩展性

**平铺式目录结构**

当项目是代码框架 / 库时，比较适合采用平铺式目录结构。

```shell
$ ls glog/
glog_file.go  glog.go  glog_test.go  LICENSE  README
```

**结构化目录结构**

Go社区比较推荐的

```md
├── api // /api 目录中存放的是当前项目对外提供的各种不同类型的 API 接口定义文件
│   ├── openapi
│   └── swagger
├── build // 存放安装包和持续集成相关的文件
│   ├── ci // 存放 CI（travis，circle，drone）的配置文件和脚本
│   ├── docker // 存放子项目各个组件的 Dockerfile 文件
│   │   ├── iam-apiserver
│   │   ├── iam-authz-server
│   │   └── iam-pump
│   ├── package // 存放容器（Docker）、系统（deb, rpm, pkg）的包配置和脚本
├── CHANGELOG
├── cmd // 一个项目有很多组件，可以把组件 main 函数所在的文件夹统一放在/cmd 目录下
│   ├── iam-apiserver
│   │   └── apiserver.go
│   ├── iam-authz-server
│   │   └── authzserver.go
│   ├── iamctl
│   │   └── iamctl.go
│   └── iam-pump
│       └── pump.go
├── configs // 存放跟应用部署相关的文件
├── CONTRIBUTING.md
├── deployments // 用来存放 Iaas、PaaS 系统和容器编排部署配置和模板（Docker-Compose，Kubernetes/Helm，Mesos，Terraform，Bosh）
├── docs // 存放设计文档、开发文档和用户文档等
│   ├── devel
│   │   ├── en-US
│   │   └── zh-CN
│   ├── guide
│   │   ├── en-US
│   │   └── zh-CN
│   ├── images
│   └── README.md
├── examples // 存放应用程序或者公共包的示例代码
├── githooks
├── go.mod
├── go.sum
├── init // 存放初始化系统（systemd，upstart，sysv）和进程管理配置文件
├── internal // 存放私有应用和库代码
│   ├── apiserver // 该目录中存放真实的应用代码。这些应用的共享代码存放在/internal/pkg 目录下
│   │   ├── api
│   │   │   └── v1 // HTTP API 接口的具体实现，主要用来做 HTTP 请求的解包、参数校验、业务逻辑处理、返回
│   │   │       └── user
│   │   ├── apiserver.go
│   │   ├── options // 应用的 command flag
│   │   ├── service // 存放应用复杂业务处理代码
│   │   ├── store
│   │   │   ├── mysql // 一个应用可能要持久化的存储一些数据
│   │   │   ├── fake
│   │   └── testing
│   ├── authzserver
│   │   ├── api
│   │   │   └── v1
│   │   │       └── authorize
│   │   ├── options
│   │   ├── store
│   │   └── testing
│   ├── iamctl // 对于一些大型项目，可能还会需要一个客户端工具
│   │   ├── cmd
│   │   │   ├── completion
│   │   │   ├── user
│   │   └── util
│   ├── pkg // 存放项目内可共享，项目外不共享的包。这些包提供了比较基础、通用的功能，例如工具、错误码、用户验证等功能
│   │   ├── code // 项目业务 Code 码。
│   │   ├── options
│   │   ├── server
│   │   ├── util
│   │   └── validation // 一些通用的验证函数
├── LICENSE
├── Makefile
├── _output
│   ├── platforms
│   │   └── linux
│   │       └── amd64
├── pkg // 存放可以被外部应用使用的代码库，其他项目可以直接通过 import 导入这里的代码
│   ├── util
│   │   └── genutil
├── README.md
├── scripts // 用来存放脚本文件，实现构建、安装、分析等不同功能
│   ├── lib
│   ├── make-rules
├── test // 存放其他外部测试应用和测试数据
│   ├── testdata
├── third_party
│   └── forked
└── tools // 存放这个项目的支持工具
```

### Go项目开发流程

- 需求阶段
- 设计阶段
- 开发阶段
  - 要思考怎么尽可能自动生成代码
  - 最终合并代码到 master 之前，要确保代码是经过充分测试的
- 测试阶段
- 发布阶段
- 运营阶段

### 写出优雅的Go项目

**代码结构**

1. 按层拆分

- View: 视图
- Controller 控制器 负责根据用户从 View 层输入的指令，选取 Model 层中的数据
- Model（模型），是应用程序中用于处理数据逻辑的部分。

```md
$ tree --noreport -L 2 layers
layers
├── controllers
│   ├── billing
│   ├── order
│   └── user
├── models
│   ├── billing.go
│   ├── order.go
│   └── user.go
└── views
    └── layouts
```

**代码规范**

Go社区 [Uber Go语言规范](https://github.com/xxjwxc/uber_go_guide_cn)

**代码质量**

- 单元测试

### Go常用设计模式

创建者模式: 在创建对象的同时隐藏创建逻辑的方式，而不是使用 new 运算符直接实例化对象

- 单例模式: 单例模式指的是全局只有一个实例，并且它负责创建自己的对象
  - 全局共享一个实例，且只需要被初始化一次的场景, 数据库实例、全局配置、全局任务池
- 工厂模式: 在你不公开内部实现的情况下，让调用者使用你提供的各种功能

```go
// 工厂模式
type Person struct {
  Name string
  Age int
}

func (p Person) Greet() {
  fmt.Printf("Hi! My name is %s", p.Name)
}

func NewPerson(name string, age int) *Person {
  return &Person{
    Name: name,
    Age: age,
  }
}
```

结构性模式: 关注类和对象的组合

- 策略模式 定义一组算法，将每个算法都封装起来，并且使它们之间可以互换
- 模板模式 定义一个操作中算法的骨架，而将一些步骤延迟到子类中

行为型模式

- 代理模式: 可以为另一个对象提供一个替身或者占位符，以控制对这个对象的访问。


### API 风格 RESTful

REST 有一系列规范，满足这些规范的 API 均可称为 `RESTful API`, 把所有内容都视为资源

- 以资源 (resource) 为中心，所有的东西都抽象成资源，所有的行为都应该是在资源上的 CRUD 操作
  - 资源对应着面向对象范式里的对象，面向对象范式以对象为中心。
  - 资源使用 URI 标识，每个资源实例都有一个唯一的 URI 标识
- 资源是有状态的，使用 JSON/XML 等在 HTTP Body 里表征资源的状态。
- 客户端通过四个 HTTP 动词，对服务器端资源进行操作，实现“表现层状态转化”
- 无状态，这里的无状态是指每个 RESTful API 请求都包含了所有足够完成本次操作的信息，服务器端无须保持 session


URI 设计:

- 资源名使用名词而不是动词，并且用名词复数表示
  - Collection：一堆资源的集合。例如我们系统里有很多用户（User）, 这些用户的集合就是 Collection
  - Member：单个特定资源。例如系统中特定名字的用户，就是 Collection 里的一个 Member
- URI 结尾不应包含/
- URI 中不能出现下划线 _，必须用中杠线 -代替
- URI 路径用小写，不要用大写。
- 避免层级过深的 URI。超过 2 层的资源嵌套会很乱，建议将其他资源转化为?参数

统一分页 / 过滤 / 排序 / 搜索功能

- 分页：在列出一个 Collection 下所有的 Member 时，应该提供分页功能，例如/users?offset=0&limit=20
- 过滤：如果用户不需要一个资源的全部状态属性，可以在 URI 参数里指定返回哪些属性，例如/users?fields=email,username,address
- 排序：用户很多时候会根据创建时间或者其他因素，列出一个 Collection 中前 100 个 Member，这时可以在 URI 参数中指明排序参数，例如/users?sort=age,desc
- 搜索：当一个资源的 Member 太多时，用户可能想通过搜索, 建议按模糊匹配来搜索。


### API风格 RPC API

`RPC（Remote Procedure Call）`，即远程过程调用，是一个计算机通信协议

服务端实现了一个函数，客户端使用 RPC 框架提供的接口，像调用本地函数一样调用这个函数，并获取返回值。RPC 屏蔽了底层的网络通信细节，使得开发人员无需关注网络编程的细节，可以将更多的时间和精力放在业务逻辑本身的实现上，从而提高开发效率


RPC 调用过程

1. Client 通过本地调用, 调用 Client Stub
2. Client Stub将参数打包成一个小希, 然后发送这个消息
3. Client 所在的 OS 将消息发送给 Server
4. Server 端接收到消息后，将消息传递给 Server Stub
5. Server Stub 将消息解包（也叫 Unmarshalling）得到参数。
6. Server Stub 调用服务端的子程序（函数），处理完后，将最终结果按照相反的步骤返回给 Client。

**gRPC 介绍**

gRPC 是由 Google 开发的高性能、开源、跨多种编程语言的通用 RPC 框架，
基于 `HTTP 2.0` 协议开发，
默认采用 `Protocol Buffers` 数据序列化协议

- 支持多种语言，例如 Go、Java、C、C++、C#、Node.js、PHP、Python、Ruby 等
- 基于 IDL（Interface Definition Language）文件定义服务，通过 proto3 工具生成指定语言的数据结构、服务端接口以及客户端 Stub
- 通信协议基于标准的 HTTP/2 设计，支持双向流、消息头压缩、单 TCP 的多路复用、服务端推送等特性。
- 支持 Protobuf 和 JSON 序列化数据格式。Protobuf 是一种语言无关的高性能序列化框架，可以减少网络传输流量，提高通信效率。

**Protocol Buffers**

- 更快的数据传输速度：protobuf 在传输时，会将数据序列化为二进制数据，和 XML、JSON 的文本传输格式相比，这可以节省大量的 IO 操作，从而提高数据传输速度
- 跨平台多语言：protobuf 自带的编译工具 protoc 可以基于 protobuf 定义文件，编译出不同语言的客户端或者服务端，供程序直接调用，因此可以满足多语言需求的场景
- 具有非常好的扩展性和兼容性，可以更新已有的数据结构，而不破坏和影响原有的程序
- 基于 IDL 文件定义服务，通过 proto3 工具生成指定语言的数据结构、服务端和客户端接口

### Makefile

高效管理项目， Makefile 来管理是目前的最佳实践


### 静态代码检查

```go
go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1
```

### 生成 swagger api 文档

```go
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

### 科学的错误码

1. 常见的错误码设计方式

- 不论请求成功或失败，始终返回200 http status code，在 HTTP Body 中包含用户账号没有找到的错误信息
- 返回http 404 Not Found错误码，并在 Body 中返回简单的错误信息
- 返回http 404 Not Found错误码，并在 Body 中返回详细的错误信息

2. 错误码设计建议

- 有区别于http status code的业务码，业务码需要有一定规则，可以通过业务码判断出是哪类错误。
- 请求出错时，可以通过http status code直接感知到请求出错
- 返回的错误信息，需要是可以直接展示给用户的安全信息，也就是说不能包含敏感信息
- 返回的数据格式应该是固定的、规范的。
- 错误信息要保持简洁，并且提供有用的信息

3. 业务Code码设计

- 可以非常方便地定位问题和定位代码行
- 错误码包含一定的信息，通过错误码可以判断出错误级别、错误模块和具体错误信息
- Go 中的 HTTP 服务器开发都是引用 net/http 包，该包中只有 60 个错误码，基本都是跟 HTTP 请求相关的错误码，在一个大型系统中，这些错误码完全不够用，而且这些错误码跟业务没有任何关联，满足不了业务的需求。引入业务的 Code 码，则可以解决这些问题
- 业务开发过程中，可能需要判断错误是哪种类型，以便做相应的逻辑处理，通过定制的错误可以很容易做到这点
- Code 码设计规范：纯数字表示，不同部位代表不同的服务，不同的模块。

100101

- 10: 服务。
- 01: 某个服务下的某个模块。
- 01: 模块下的错误码序号，每个模块可以注册 100 个错误。

4. 设置HTTP Status Code

Go net/http 提供了 60个错误码, 分为如下5类:

- 1XX - 指示信息 表示请求已接收, 继续处理
- 2XX - 请求成功 表示成功处理了请求的状态代码
- 3XX - 请求被重定向 表示要完成请求，需要进一步操作。通常，这些状态代码用来重定向
- 4XX - 请求错误 表示请求可能出错，妨碍了服务器的处理，通常是客户端出错，需要客户端做进一步的处理
- 5XX - 服务器错误 示服务器在尝试处理请求时发生内部错误。这些错误可能是服务器本身的错误，而不是客户端的问题。
- 200 成功
- 400 客户端出问题
- 500 服务端出问题
- 401 认证失败
- 403 授权失败
- 404 资源找不到

### 设计错误包

```go
type withCode struct {
	err error // 错误
	code int // 业务错误码
	cause error // cause error
	*stack // 错误堆栈
}
```

### 日志包

- 标准库 log 包
- glog
- logrus
- zap

### 构建应用

- 命令行参数解析: 主要用来解析命令行参数，这些命令行参数可以影响命令的运行效果
  - `Pflag` 包
- 配置文件解析：一个大型应用，通常具有很多参数，为了便于管理和配置这些参数，通常会将这些参数放在一个配置文件中，供程序读取并解析。
  - `Viper` 包
- 应用的命令行框架：应用最终是通过命令来启动的。这里有 3 个需求点，一是命令需要具备 Help 功能，这样才能告诉使用者如何去使用；二是命令需要能够解析命令行参数和配置文件；三是命令需要能够初始化业务代码，并最终启动业务进程。也就是说，我们的命令需要具备框架的能力，来纳管这 3 个部分。
  - `Cobra` 包

### 构建应用实战

**应用的三大基本功能**

Go 后端服务，基本上可以分为 API 服务和非 API 服务两类。

- API 服务：通过对外提供 HTTP/RPC 接口来完成指定的功能。比如订单服务，通过调用创建订单的 API 接口，来创建商品订单。
- 非 API 服务：通过监听、定时运行等方式，而不是通过 API 调用来完成某些任务。比如数据处理服务，定时从 Redis 中获取数据，处理后存入后端存储中。再比如消息处理服务，监听消息队列（如 NSQ/Kafka/RabbitMQ），收到消息后进行处理。

启动流程基本一致，都可以分为三步：

1. 应用框架的构建
2. 应用初始化

### Web服务核心功能

- Web 服务最核心的功能是路由匹配

- 参数解析、参数校验、逻辑处理、返回结果
- HTTP 参数解析和返回
  - path 路径参数
  - query 字符串参数
  - form 表单参数
  - http header 参数 header
  - body 消息体参数 

**Gin**

1. HTTP/HTTPS 支持
```go
insecureServer := &http.Server{
  Addr:         ":8080",
  Handler:      router(),
  ReadTimeout:  5 * time.Second,
  WriteTimeout: 10 * time.Second,
}
...
err := insecureServer.ListenAndServe()
```

```go
secureServer := &http.Server{
  Addr:         ":8443",
  Handler:      router(),
  ReadTimeout:  5 * time.Second,
  WriteTimeout: 10 * time.Second,
}
...
err := secureServer.ListenAndServeTLS("server.pem", "server.key")
```

2. JSON数据支持

Gin 支持多种数据通信格式，例如 `application/json、application/xml`。
可以通过`c.ShouldBindJSON`函数，将 Body 中的 JSON 格式数据解析到指定的 Struct 中，
通过`c.JSON`函数返回 JSON 格式的数据。

3. 路由匹配

- 精确匹配 `/products/:name`
- 模糊匹配 `products/*name`

4. 路由分组

```go
v1 := router.Group("/v1", gin.BasicAuth(gin.Accounts{"foo": "bar", "colin": "colin404"}))
{
    productv1 := v1.Group("/products")
    {
        // 路由匹配
        productv1.POST("", productHandler.Create)
        productv1.GET(":name", productHandler.Get)
    }

    orderv1 := v1.Group("/orders")
    {
        // 路由匹配
        orderv1.POST("", orderHandler.Create)
        orderv1.GET(":name", orderHandler.Get)
    }
}

v2 := router.Group("/v2", gin.BasicAuth(gin.Accounts{"foo": "bar", "colin": "colin404"}))
{
    productv2 := v2.Group("/products")
    {
        // 路由匹配
        productv2.POST("", productHandler.Create)
        productv2.GET(":name", productHandler.Get)
    }
}
```

5. 一进程多服务

```go
var eg errgroup.Group
insecureServer := &http.Server{...}
secureServer := &http.Server{...}

eg.Go(func() error {
  err := insecureServer.ListenAndServe()
  if err != nil && err != http.ErrServerClosed {
    log.Fatal(err)
  }
  return err
})
eg.Go(func() error {
  err := secureServer.ListenAndServeTLS("server.pem", "server.key")
  if err != nil && err != http.ErrServerClosed {
    log.Fatal(err)
  }
  return err
}

if err := eg.Wait(); err != nil {
  log.Fatal(err)
})
```


6. 参数解析 参数校验 逻辑处理 返回结果

```go
func (u *productHandler) Create(c *gin.Context) {
  u.Lock()
  defer u.Unlock()

  // 1. 参数解析
  var product Product
  if err := c.ShouldBindJSON(&product); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // 2. 参数校验
  if _, ok := u.products[product.Name]; ok {
    c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("product %s already exist", product.Name)})
    return
  }
  product.CreatedAt = time.Now()

  // 3. 逻辑处理
  u.products[product.Name] = product
  log.Printf("Register product %s success", product.Name)

  // 4. 返回结果
  c.JSON(http.StatusOK, product)
}
```

```go
// Gin 提供了一些函数，来分别读取这些 HTTP 参数，每种类别会提供两种函数
// 一种函数可以直接读取某个参数的值
// 另外一种函数会把同类 HTTP 参数绑定到一个 Go 结构体中

gin.Default().GET("/:name/:id", nil)

name := c.Param("name")
action := c.Param("action")

type Person struct {
  ID string `uri:"id" binding:"required,uuid"`
  Name string `uri:"name" binding:"required"`
}

if err := c.ShouldBindUri(&person); err != nil {
  // normal code
  return
}
```

Gin 在绑定参数时，是通过结构体的 tag 来判断要绑定哪类参数到结构体中的

- 路径参数 uri `ShouldBindUri BindUri`
- 查询字符串参数 form `ShouldBindQuery BindQuery`
- 表单参数 form `ShouldBind`
- http头参数 header `shouldBindHeader BindHeader`
- 消息体参数, 根据Content-Type `shouldBindJSON BindJSON`


针对每种参数类型，Gin 都有对应的函数来获取和绑定这些参数。这些函数都是基于如下两个函数进行封装的：

- `ShouldBindWith(obj interface{}, b binding.Binding) error`

很多 ShouldBindXXX 函数底层都是调用 ShouldBindWith 函数来完成参数绑定的,
该函数会根据传入的绑定引擎，将参数绑定到传入的结构体指针中
如果绑定失败，只返回错误内容，但`不终止 HTTP` 请求


- `MustBindWith(obj interface{}, b binding.Binding) error`

很多 BindXXX 函数底层都是调用 MustBindWith 函数来完成参数绑定的
该函数会根据传入的绑定引擎，将参数绑定到传入的结构体指针中
如果绑定失败，`返回错误并终止请求，返回 HTTP 400 错误`

7. 中间件

- 中间件做成可加载的，通过配置文件指定程序启动时加载哪些中间件
- 只将一些通用的、必要的功能做成中间件。
- 在编写中间件时，一定要保证中间件的代码质量和性能

```go
// 中间件作用于所有的HTTP请求
router.Use(gin.Logger(), gin.Recovery()) 
```

8. 认证 RequestID 跨域

```go
router := gin.New()

// 认证
router.Use(gin.BasicAuth(gin.Accounts{"foo": "bar", "colin": "colin404"}))

// RequestID
router.Use(requestid.New(requestid.Config{
    Generator: func() string {
        return "test"
    },
}))

// 跨域
// CORS for https://foo.com and https://github.com origins, allowing:
// - PUT and PATCH methods
// - Origin header
// - Credentials share
// - Preflight requests cached for 12 hours
router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https://foo.com"},
    AllowMethods:     []string{"PUT", "PATCH"},
    AllowHeaders:     []string{"Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
        return origin == "https://github.com"
    },
    MaxAge: 12 * time.Hour,
}))
```

### 访问认证

- 认证（Authentication，英文缩写 authn）：用来验证某个用户是否具有访问系统的权限
。如果认证通过，该用户就可以访问系统，从而创建、修改、删除、查询平台支持的资源
- 授权（Authorization，英文缩写 authz）：用来验证某个用户是否具有访问某个资源的权限，
如果授权通过，该用户就能对资源做增删改查等操作

四种基本的认证方式

1. Basic 认证（基础认证），是最简单的认证方式; 简单地将用户名:密码进行 base64 编码后，
放到 HTTP Authorization Header 中。HTTP 请求到达后端服务后，
后端服务会解析出 Authorization Header 中的 base64 字符串，
解码获取用户名和密码，并将用户名和密码跟数据库中记录的值进行比较，
如果匹配则认证通过
2. Digest 认证（摘要认证），是另一种 HTTP 认证协议，它与基本认证兼容，但修复了基本认证的严重缺陷

- 绝不会用明文方式在网络上发送密码
- 可以有效防止恶意用户进行重放攻击
- 可以有选择地防止对报文内容的篡改。

3. OAuth

OAuth（开放授权）是一个开放的授权标准，
允许用户让第三方应用访问该用户在某一 Web 服务上存储的私密资源（例如照片、视频、音频等），
而无需将用户名和密码提供给第三方应用

OAuth2.0 一共分为四种授权方式，分别为密码式、隐藏式、凭借式和授权码模式

4. Bearer

Bearer 认证，也称为令牌认证，是一种 HTTP 身份验证方法。
Bearer 认证的核心是 bearer token

### GORM

```go
// 定义了一个 GORM 模型（Models）
type Product struct {
    gorm.Model
    Code  string `gorm:"column:code"`
    Price uint   `gorm:"column:price"`
}

// TableName maps to mysql table name.
func (p *Product) TableName() string {
    return "product"
}
```

```go
// 1. Auto migration for given models
// 只对新增的字段或索引进行变更
db.AutoMigrate(&Product{})
```

```go
// 插入表记录
// 2. Insert the value into database
if err := db.Create(&Product{Code: "D42", Price: 100}).Error; err != nil {
    log.Fatalf("Create error: %v", err)
}
PrintProducts(db)
```

```go
// 获取符合条件的记录
// 3. Find first record that match given conditions
product := &Product{}
if err := db.Where("code= ?", "D42").First(&product).Error; err != nil {
    log.Fatalf("Get product error: %v", err)
}
```

```go
// 更新表记录
// 4. Update value in database, if the value doesn't have primary key, will insert it
product.Price = 200
if err := db.Save(product).Error; err != nil {
    log.Fatalf("Update product error: %v", err)
}
PrintProducts(db)
```

```go
// 删除
// 5. Delete value match given conditions
if err := db.Where("code = ?", "D42").Delete(&Product{}).Error; err != nil {
    log.Fatalf("Delete product error: %v", err)
}
PrintProducts(db)
```

1. 模型定义

GORM 使用模型(Models) 来映射一个数据库表, 默认情况下, 使用 ID 作为主键;
使用 CreatedAt、UpdatedAt、DeletedAt 字段追踪创建、更新和删除时间

```go
type Animal struct {
  AnimalID int64     `gorm:"column:animalID;primarykey"` // 将列名设为 `animalID`
  Birthday time.Time `gorm:"column:birthday"`            // 将列名设为 `birthday`
  Age      int64     `gorm:"column:age"`                 // 将列名设为 `age`
}

func (a *Animal) TableName() string {
  return "animal"
}
```

2. 连接数据库

GORM 支持连接池，底层是用 database/sql 包来维护连接池的，连接池设置如下：
```go
sqlDB, err := db.DB()
sqlDB.SetMaxIdleConns(100)              // 设置MySQL的最大空闲连接数（推荐100）
sqlDB.SetMaxOpenConns(100)             // 设置MySQL的最大连接数（推荐100）
sqlDB.SetConnMaxLifetime(time.Hour)    // 设置MySQL的空闲连接最大存活时间（推荐10s）
```