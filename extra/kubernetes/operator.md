# k8s operator开发

## 环境准备

+ go version v1.15+ (kubebuilder v3.0 < v3.1).
+ go version v1.16+ (kubebuilder v3.1 < v3.3).
+ go version v1.17+ (kubebuilder v3.3+).
+ docker version 17.03+.
+ kubectl version v1.11.3+.
+ Access to a Kubernetes v1.11.3+ cluster.

## 安装 kubebuilder

### Mac/Linux系统安装方法
```sh
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```

### Windows系统安装方法

非常不幸, 该工具官方没有windows的包

![](./images/kubebuilder-nowin.png)


所以我们需要通过源码安装
```sh
# 确保Proxy已经设置ok
export GOPROXY=https://goproxy.io,direct
# clone kubebuilder镜像仓库, 注意我直接使用的master
git clone git@github.com:kubernetes-sigs/kubebuilder.git
# 注意 提前搞定你的make命令安装
cd kubebuilder
make build
# 编译完成后, kubebuilder的可执行文件在 当前目录的bin下
```

最后将可执行文件 copy到 git bash安装目录下的 /usr/bin下面, windows由于权限问题, 需要你最近以管理员身份copy

copy完成后, 重新开启一个新的 git bash, 然后kubebuilder 这个命令就可以正常使用了

关于kubebuilder的使用说明一定要阅读:
```sh
$ kubebuilder -h
CLI tool for building Kubernetes extensions and tools.

Usage:
  kubebuilder [flags]
  kubebuilder [command]

Examples:
The first step is to initialize your project:
    kubebuilder init [--plugins=<PLUGIN KEYS> [--project-version=<PROJECT VERSION>]]

<PLUGIN KEYS> is a comma-separated list of plugin keys from the following table     
and <PROJECT VERSION> a supported project version for these plugins.

                        Plugin keys | Supported project versions
------------------------------------+----------------------------
          base.go.kubebuilder.io/v3 |                          3
   declarative.go.kubebuilder.io/v1 |                       2, 3
               go.kubebuilder.io/v2 |                       2, 3
               go.kubebuilder.io/v3 |                          3
 kustomize.common.kubebuilder.io/v1 |                          3

For more specific help for the init command of a certain plugins and project version
configuration please run:
    kubebuilder init --help --plugins=<PLUGIN KEYS> [--project-version=<PROJECT VERSION>]

Default plugin keys: "go.kubebuilder.io/v3"
Default project version: "3"


Available Commands:
  alpha       Alpha-stage subcommands
  completion  Load completions for the specified shell
  create      Scaffold a Kubernetes API or webhook
  edit        Update the project configuration
  help        Help about any command
  init        Initialize a new project
  version     Print the kubebuilder version

Flags:
  -h, --help                     help for kubebuilder
      --plugins strings          plugin keys to be used for this subcommand execution
      --project-version string   project version (default "3")

Use "kubebuilder [command] --help" for more information about a command.
```

## 创建项目

我们需要使用kubebuilder来为我们生成Operator开发的框架代码

kubebuilder 提供了一个 init命令用来初始化一个新的Operator工程目录，具体用法如下:
```sh
$ kubebuilder init -h
Initialize a new project including the following files:
  - a "go.mod" with project dependencies
  - a "PROJECT" file that stores project configuration
  - a "Makefile" with several useful make targets for the project
  - several YAML files for project deployment under the "config" directory
  - a "main.go" file that creates the manager that will run the project controllers

Usage:
  kubebuilder init [flags]

Examples:
  # Initialize a new project with your domain and name in copyright
  kubebuilder init --plugins go/v3 --domain example.org --owner "Your name"

  # Initialize a new project defining a specific project version
  kubebuilder init --plugins go/v3 --project-version 3


Flags:
      --component-config         create a versioned ComponentConfig file, may be 'true' or 'false'
      --domain string            domain for groups (default "my.domain")
      --fetch-deps               ensure dependencies are downloaded (default true)
  -h, --help                     help for init
      --license string           license to use to boilerplate, may be one of 'apache2', 'none' (default "apache2")
      --owner string             owner to add to the copyright
      --project-name string      name of this project
      --project-version string   project version (default "3")
      --repo string              name to use for go module (e.g., github.com/user/repo), defaults to the go package of the current working directory.
      --skip-go-version-check    if specified, skip checking the Go version

Global Flags:
      --plugins strings   plugin keys to be used for this subcommand execution
```

关键参数说明:
+ --plugins, 指定生成代码的插件, 默认使用 "go.kubebuilder.io/v3"
+ --project-version 支持的项目版本, 有2,3 默认3
+ --repo module path, 就是 go mod init 指定的go module的名称
+ --domain, 组织名称, 用于API Group等
+ --owner, operater所有者, 一般填写开发者邮箱

然后开始初始化项目:
```sh
$ kubebuilder init --domain magedu.com --repo gitee.com/go-course/k8s-operator

Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.11.0
go: downloading k8s.io/client-go v0.23.0
go: downloading github.com/evanphx/json-patch v4.12.0+incompatible
go: downloading gomodules.xyz/jsonpatch/v2 v2.2.0
go: downloading github.com/prometheus/common v0.28.0
go: downloading k8s.io/component-base v0.23.0
go: downloading github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
go: downloading golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
go: downloading github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369
Update dependencies:
$ go mod tidy
go: downloading github.com/go-logr/zapr v1.2.0
go: downloading github.com/onsi/gomega v1.17.0
go: downloading go.uber.org/zap v1.19.1
go: downloading github.com/Azure/go-autorest/autorest v0.11.18
go: downloading github.com/Azure/go-autorest/autorest/adal v0.9.13
go: downloading go.uber.org/goleak v1.1.12
go: downloading cloud.google.com/go v0.81.0
go: downloading gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
go: downloading github.com/Azure/go-autorest/logger v0.2.1
go: downloading github.com/form3tech-oss/jwt-go v3.2.3+incompatible
Next: define a resource with:
$ kubebuilder create api
```

## 创建API




## 参考

+ [kubebuilder 官方文档](https://book.kubebuilder.io/introduction.html)
+ [operator-sdk 官方文档](https://sdk.operatorframework.io/docs/)