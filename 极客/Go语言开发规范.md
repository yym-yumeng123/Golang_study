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


### API 风格