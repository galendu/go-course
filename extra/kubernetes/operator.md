# k8s operator开发

## 需求与设计

之前讲过使用Traefik + etcd的 内外网关打通的方案, 我们的服务在启动的时候可以通过框架直接注册到etcd中去, 那如果是其他框架或者语言, 比如 php, python, java 也想使用 Traefik+etcd这套方案这么办?

直接能想到的就是 每个框架和语言 都实现一个套 注册到etcd的功能,  实现起来也并不难, 但是这就面临 对别人服务的侵入性, 别人不一定愿意接受

现在的微服务开发，基本都基于容器部署, 而容器管理平台k8s也是当今 容器编排工具的标准, k8s本身是知道当前系统中有哪些service服务的, 那能否通过感知service的变化，来把service信息动态写入的etcd喃?

那我们的解决方案就很简单了:

k8s service <---watch---- service operater ------> etcd

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

通过脚手架为我们提供的create api来创建 CRD 相关的Resource 和控制器:
```
$ kubebuilder create api -h
Scaffold a Kubernetes API by writing a Resource definition and/or a Controller.

If information about whether the resource and controller should be scaffolded  
was not explicitly provided, it will prompt the user if they should be.        

After the scaffold is written, the dependencies will be updated and
make generate will be run.

Usage:
  kubebuilder create api [flags]

Examples:
  # Create a frigates API with Group: ship, Version: v1beta1 and Kind: Frigate
  kubebuilder create api --group ship --version v1beta1 --kind Frigate

  # Edit the API Scheme
  nano api/v1beta1/frigate_types.go

  # Edit the Controller
  nano controllers/frigate/frigate_controller.go

  # Edit the Controller Test
  nano controllers/frigate/frigate_controller_test.go

  # Generate the manifests
  make manifests

  # Install CRDs into the Kubernetes cluster using kubectl apply
  make install

  # Regenerate code and run against the Kubernetes cluster configured by ~/.kube/config
  make run


Flags:
      --controller           if set, generate the controller without prompting the user (default true)
      --force                attempt to create resource even if it already exists
      --group string         resource Group
  -h, --help                 help for api
      --kind string          resource Kind
      --make make generate   if true, run make generate after generating files (default true)
      --namespaced           resource is namespaced (default true)
      --plural string        resource irregular plural form
      --resource             if set, generate the resource without prompting the user (default true)
      --version string       resource Version

Global Flags:
      --plugins strings   plugin keys to be used for this subcommand execution
```

### 生成样例代码

```sh
kubebuilder create api --group traefik --version v1 --kind TraefikService

Create Resource [y/n]
y
Create Controller [y/n]
y
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
api\v1\traefikservice_types.go
controllers\traefikservice_controller.go
Update dependencies:
$ go mod tidy
Running make:
$ make generate
/e/Projects/Golang/go-course-projects/k8s-operator/bin/controller-gen object:headerFile="hack\\boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests
```

样例生成完了后，我们会看到我们项目新增了一些文件:

![](./images/create-api.png)

+ api/v1 目录下主要存放是我们API Object, 就是我们的Resource对象相关信息
+ config/crd 目录下是我们crd的描述文件, 我们需要把自定义资源(CRD)的描述信息注册给k8s时需要的
+ rbac 目录下 存放着 关于CRD资源的 role定义的样例文件(editor/viewer)
+ samples 目录下 存放着 CRD的一个样例文件, 后面部署完成后可以 直接编辑下 apply到k8s集群中去
+ controllers 目录下 存放着 我们所有的 Object 的Controller 代码

### CRD开发

按照之前的设计，我们其实是没有必要定义CRD的, 下面关于CRD的定义和安装 是出于教学演示目的, 如果你只想做项目，可以忽略这部分，甚至删除kubebuilder为我们生成关于CRD定义部分的相关代码

#### CRD 设计

我们需要定义Traefik Service, 我们来看看Traefik Service一个service 实例定义:
```
<etcd_prefix>/<entry_point>/services/loadBalancer/servers/<index>/url   <url_value>

traefik etcd配置的前缀, provider配置时 有设置
services: 表示 web entrypoint的 services配置
loadBalancer: cmdb 服务loadBalancer配置
servers: loadBalancer 下的实例配置
0(变量): index
```

因此我们定义的Service需要有如下属性:
+ entrypoint name
+ service name
+ service url

作为k8s的CRD 必须是一个runtime.Object, 也就是为我们生成的TraefikService对象, 我们需要编辑的是TraefikServiceSpec对象

由于 Name已经在 ObjectMeta 有声明了, 因此我们只需要添加 entrypoint 和 url

修改资源定义: api/v1/traefikservice_types.go 
```go
// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TraefikServiceSpec defines the desired state of TraefikService
type TraefikServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of TraefikService. Edit traefikservice_types.go to remove/update
	Entrypoint string `json:"entrypoint"`
	URL        string `json:"url"`
}
```

#### CRD代码生成

我们通过make 提供 
 
+ manifests 重新生成修改后的 CRD定义描述
+ generate 重新生成代码

```
$ make manifests generate

/e/Projects/Golang/go-course-projects/k8s-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
/e/Projects/Golang/go-course-projects/k8s-operator/bin/controller-gen object:headerFile="hack\\boilerplate.go.txt" paths="./..."
```

我们看到CRD的描述文件已经有变化了:

![](./images/treafik-crd-define.png)

#### 安装CRD



#### 验证CRD


### Crontroller开发


## 参考

+ [kubebuilder 官方文档](https://book.kubebuilder.io/introduction.html)
+ [operator-sdk 官方文档](https://sdk.operatorframework.io/docs/)