# 微服务网关

## 微服务架构

### 基于网关的架构

没有网关之前，服务直接是直接访问的, 如果外部需要访问某个服务, 需要知道每个服务的地址：

![](./images/service-call.png)

有了微服务后, 我们只需要知道网关的地址就可以了, 服务具体在哪儿，由网关负责寻找

![](./images/api-gw.png)

传统的网关比如Nginx, 服务的发现都需要手动配置, 并不支持自动发现

![](./images/nginx.png)


因此我们需要选择一个微服务网关, 让我们的服务可以自动注册

常见的微服务网关有:
+ Kong(openresty), 基于Lua脚本进行扩展
+ APISix(openresty), 基于Lua脚本进行扩展
+ Traefik, Go 云原生网关

### 为什么选择Traefik

我们选择一个微服务网关大体的需求如下:

+ 服务启动时可以直接注册到注册中心(REST API/GRPC)
+ 网关可以通过注册中心发现内部服务
+ 网关可以选择配置哪些服务队伍提供访问
+ 内部服务通过注册中心直接访问，也可以走网关访问
+ 网关可以配置灰度发布‘
+ 网关可以配置链路追踪
+ 网关有流量监控


Treafik在综合比较下是一个不错的方案:

+ 完善的API接口与文档
+ 高性能,稳定性,高可用
+ 架构简单，易于维护
+ 完善的流量管理策略(Path匹配, 灰度，流量镜像)
+ 灵活的扩展定制能力(比如Nginx Lua)
+ 能对接注册中心实现服务发现
+ 与微服务生态比较亲近, 比如k8s
+ 支持GRPC代理

## Traefik 介绍

### 基本概念

网关对性能和可靠性要求比较高, 这里选用Traefik做完我们微服务的网关, 也方便我们二次开发

下面是Traefik的流程示意图:

![](./images/traefik-summary.jpeg)


traefik通过路由规则(Routes) 来配置 Endpint和service进行流量的分发

![](./images/traefik-flow.jpeg)

在路由里面我们通过配置路由规则进行流量转发, 比如Host规则就是匹配Host进行调度

![](./images/traefik-routes.jpeg)

最后把 流量路由到我们的 服务组内

![](./images/traefik-service.jpeg)

下面是一个服务组的定义

```yaml
## Dynamic configuration
http:
  services:
    my-service:
      loadBalancer:
        servers:
        - url: "http://<private-ip-server-1>:<private-port-server-1>/"
        - url: "http://<private-ip-server-2>:<private-port-server-2>/"
```

### 配置介绍

Traefik里面的配置由2部分组成:
+ 静态配置: EntryPoints 和 Provider 需要在启动时配置好, [详细说明](https://doc.traefik.io/traefik/reference/static-configuration/overview/)
+ 动态配置: 路由规则和服务注册 可以动态发现, [详细说明](https://doc.traefik.io/traefik/reference/dynamic-configuration/file/)

![](./images/traefik-config.jpeg)

Traefik和其他网关不同之处，在于其灵活的服务配置(服务发现)

![](./images/traefik_provider.jpeg)

其中最灵活的是Etcd, 我们将由程序自己实现服务的注册, 这样我们对外通过Traefik暴露我们的服务, 对内通过Etcd作为服务注册中心, 直接调用


#### 配置EntryPoint

下面是Yaml文件的配置:
```yaml
## Static configuration
entryPoints:
  web:
    address: ":80"

  websecure:
    address: ":443"

  grpc:
    address: ":18080"
```

下面是基于环境变量的配置:
```
TRAEFIK_ENTRYPOINTS_<NAME>:
Entry points definition. (Default: false)

TRAEFIK_ENTRYPOINTS_<NAME>_ADDRESS:
Entry point address.
```

#### 配置 Etcd Provider

下面是Etcd相关配置: [Traefik & Etcd](https://doc.traefik.io/traefik/providers/etcd/)

```yaml
providers:
  etcd:
    endpoints:
      - "127.0.0.1:2379"
    rootKey: "traefik"
    username: "foo"
    password: "bar"
    tls:
      ca: path/to/ca.crt
      caOptional: true
      cert: path/to/foo.cert
      key: path/to/foo.key
      insecureSkipVerify: true
```

如果要采用环境变量设置需要查阅: traefik支持的[配置变量](https://doc.traefik.io/traefik/reference/static-configuration/env/) 
该配置的参数都是 TRAEFIK_PROVIDERS_ETCD 打头的: 
```
TRAEFIK_PROVIDERS_ETCD:
Enable Etcd backend with default settings. (Default: false)

TRAEFIK_PROVIDERS_ETCD_ENDPOINTS:
KV store endpoints (Default: 127.0.0.1:2379)

TRAEFIK_PROVIDERS_ETCD_PASSWORD:
KV Password

TRAEFIK_PROVIDERS_ETCD_ROOTKEY:
Root key used for KV store (Default: traefik)

TRAEFIK_PROVIDERS_ETCD_USERNAME:
KV Username
```

#### API配置

```yaml
api:
  insecure: true
  dashboard: true
  debug: true
```

```
TRAEFIK_API:
Enable api/dashboard. (Default: false)

TRAEFIK_API_DASHBOARD:
Activate dashboard. (Default: true)

TRAEFIK_API_DEBUG:
Enable additional endpoints for debugging and profiling. (Default: false)

TRAEFIK_API_INSECURE:
Activate API directly on the entryPoint named traefik. (Default: false)
```

## 网关设计

那基于Traefik如何实现服务的自动发现喃?

traefik支持以etcd做完配置中心, 因此我们自己基于Traefik的格式 开发一套注册中心 可以对接Traefik了


## 安装Traefik

etcd的安装参考上节, 下面介绍Traefik的搭建

这里采用docker安装, 首先拉去官方镜像:
```go
docker pull traefik
```

准备好配置文件: traefik.yaml:
```yaml
api:
  insecure: true
  dashboard: true
  debug: true

entryPoints:
  web:
    address: ":80"

  websecure:
    address: ":443"

  grpc:
    address: ":18080"

providers:
  etcd:
    endpoints:
      - "<your host ip>:2379"
    rootKey: "traefik"
```

启动
```
# 其中 8080 是 traefik dashboard的地址
# 80 是web,  18080 是grpc, 443不测试 故不暴露
docker run -d -p 8080:8080 -p 80:80 -p 18080:18080 \
    -v $PWD/traefik.yml:/etc/traefik/traefik.yml traefik:latest
```

然后访问: http://localhost:8080/dashboard 就可以看到Traefik dashboard了

![](./images/traefik-db.png)

更详细的安装文档请求参考: [Install Traefik](https://doc.traefik.io/traefik/getting-started/install-traefik/)

## 服务注册与访问

我们的服务想要接入到Traefik, 需要2步:
+ 服务注册: 配置services
+ 路由匹配: 配置Router

我们先不忙 对接Etcd Provider, 我们先通过配置文件来看看, 如何配置

### 配置详解

以cmdb为例, cmdb提供2个服务:
+ HTTP RESTful API: 80
+ GRPC RPC:  18080

```yaml
## dynamic configuration ##

# web entrypoint 配置
# web entrypoint 监听的地址 在traefik启动的时候已经配置好了: 80
http:
  # services 配置, 具体配置请查看: https://doc.traefik.io/traefik/routing/services/
  services:
    # cmdb service 配置
    cmdb-api:
      loadBalancer:
        # cmdb 服务的实例
        servers:
        - url: http://127.0.0.1:8080
    # 注册 cmdb 的grpc服务, gprc采用http2协议, h2c兼容, 具体配置请参考 https://doc.traefik.io/traefik/user-guides/grpc/
    cmdb-grpc:
      loadBalancer:
        servers:
        - url: h2c://127.0.0.1:18080

  # routers 配置, 具体配置请参考: https://doc.traefik.io/traefik/routing/routers/
  routers:
    # cmdb 服务的路由
    cmdb:
      entryPoints:
        - "web"
      service: cmdb-api
      rule: PathPrefix(`/cmdb/api/v1`)
    # cmdb的所有服务都有统一的前缀, 通过gprc生成的文件可以看到
    cmdb:
      entryPoints:
        - "grpc"
      service: cmdb-grpc
      rule: PathPrefix(`infraboard.cmdb`)
```

### 基于KV的服务发现

现在Traefik使用etcd provider动态发现配置, 因此我们需要将上面的配置转化为符合etcd provider的格式写入etcd中

完整的文档请参考: [KV Configuration Reference](https://doc.traefik.io/traefik/reference/dynamic-configuration/kv/)


#### Service配置

服务注册 核心配置的是服务的地址, 也就是URL, 比如我有2个cmdb实例, 注册到etcd里面, key value结构大致如下: 
```
traefik/http/services/Service01/loadBalancer/servers/0/url	foobar
traefik/http/services/Service01/loadBalancer/servers/1/url  foobar
```

由此我们可以看出traefik的 key结构设计和yaml的结构设计是一样的, 只是etcd里面的使用/做完分层的方式:
```
<etcd_prefix>/<entry_point>/services/loadBalancer/servers/<index>/url   <url_value>
```

+ traefik etcd配置的前缀, provider配置时 有设置
+ web(变量): entrypoint 名称
+ services: 表示 web entrypoint的 services配置
+ cmdb(变量): 表示是cmdb服务的配置
+ loadBalancer: cmdb 服务loadBalancer配置
+ servers: loadBalancer 下的实例配置
+ 0(变量):  index
+ url: 实例的地址

可以测试下: 
```
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/Service01/loadBalancer/servers/0/url foobar
```

![](./images/add-service.png)

#### Router配置

对于Router而言，核心需要配置的是:
+ entry point
+ rule
+ service

```
traefik/http/routers/Router1/entryPoints/0	foobar
traefik/http/routers/Router1/rule	foobar
traefik/http/routers/Router1/service	foobar
```

#### 完整配置

下面是cmdb和keyauth服务的服务发现配置

```
# cmdb 和 keyauth services 配置
traefik/http/services/cmdb/loadBalancer/servers/0/url	http://127.0.0.1:8060
traefik/http/services/cmdb/loadBalancer/servers/0/url	h2c://127.0.0.1:18060
traefik/http/services/keyauth/loadBalancer/servers/0/url	http://127.0.0.1:8050
traefik/http/services/keyauth/loadBalancer/servers/0/url	h2c://127.0.0.1:18050

# cmdb 和 keyauth router配置
traefik/http/routers/cmdb-api/entryPoints/0	web
traefik/http/routers/cmdb-api/rule	PathPrefix(`/cmdb/api/v1`)
traefik/http/routers/cmdb-api/service cmdb-api

traefik/http/routers/cmdb-grpc/entryPoints/0 grpc
traefik/http/routers/cmdb-grpc/rule PathPrefix(`infraboard.cmdb`)
traefik/http/routers/cmdb-grpc/service cmdb-grpc

traefik/http/routers/keyauth-api/entryPoints/0 web
traefik/http/routers/keyauth-api/rule	PathPrefix(`/keyauth/api/v1`)
traefik/http/routers/keyauth-api/service keyauth-api

traefik/http/routers/keyauth-grpc/entryPoints/0 grpc
traefik/http/routers/keyauth-grpc/rule PathPrefix(`infraboard.keyauth`)
traefik/http/routers/keyauth-grpc/service keyauth-grpc
```

#### 验证测试

我们手动操作etcd来把上面的配置写入: 
```
# cmdb http 配置, 注意填写本级的IP, 因为traefik是在容器里面访问的
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api/loadBalancer/servers/0/url http://192.168.31.16:8060
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/routers/cmdb-api/rule 'PathPrefix(`/cmdb/api/v1`)'
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/routers/cmdb-api/service cmdb-api
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/routers/cmdb-api/entryPoints/0	web

# cmdb grpc 配置
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-grpc/loadBalancer/servers/0/url h2c://192.168.31.16:18060
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/routers/cmdb-grpc/rule 'PathPrefix(`/infraboard.cmdb`)'
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/routers/cmdb-grpc/service cmdb-grpc
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/routers/cmdb-grpc/entryPoints/0 grpc

通过etcd我们可以看到key已经写入
```sh
$ docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl get --prefix  traefik 
traefik/http/routers/cmdb-api/entryPoints/0
web
traefik/http/routers/cmdb-api/rule
PathPrefix(`/cmdb/api/v1`)
traefik/http/routers/cmdb-api/service
cmdb-api
traefik/http/routers/cmdb-grpc/entryPoints/0
grpc
traefik/http/routers/cmdb-grpc/rule
PathPrefix(`/infraboard.cmdb`)
traefik/http/routers/cmdb-grpc/service
cmdb-grpc
traefik/http/services/Service01/loadBalancer/servers/0/url
foobar
traefik/http/services/cmdb-api/loadBalancer/servers/0/url
http://127.0.0.1:8060
traefik/http/services/cmdb-grpc/loadBalancer/servers/0/url
h2c://127.0.0.1:18060                                                         
```

![](./images/cmdb.png)


启动cmdb服务, 确认好 http和grpc服务监听的端口:
```
$ cd cmdb
$ make run
```

验证 http 访问
```sh
$ curl localhost/cmdb/api/v1/hosts
{"code":400,"namespace":"global","reason":"请求不合法","message":"token required"}
```

验证 grpc 访问: 运行之前的测试用例测试
```go
func TestClient(t *testing.T) {
	should := assert.New(t)
	conf := client.NewConfig("localhost:18080")
	conf.WithClientCredentials("nHerVBlrKIDurviMGUXVOQHC", "l5FB38Mw2JmxHgGm8rUcich2ZrGRVrl7")

	c, err := client.NewClient(conf)
	if should.NoError(err) {
		rs, err := c.Resource().Search(context.Background(), resource.NewSearchRequest())
		should.NoError(err)
		fmt.Println(rs)
	}
}
```

## 其他功能

###  灰度发布

灰度发布需要我们控制不通版本的集群的流量, traefik的Weighted Round Robin (service)提供该功能的支持

```yaml
## Dynamic configuration
http:
  services:
    cmdb-api-app:
      weighted:
        services:
        - name: appv1
          weight: 3
        - name: appv2
          weight: 1

    cmdb-api-v1:
      loadBalancer:
        servers:
        - url: "http://192.168.31.16:8060"

    cmdb-api-v2:
      loadBalancer:
        servers:
        - url: "http://192.168.31.16:8070"
```


服务注册
```
# appv1和appv2
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-v1/loadBalancer/servers/0/url http://172.22.111.202:8060
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-v2/loadBalancer/servers/0/url http://172.22.111.202:8070

# Weighted Round Robin (service)
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-app/weighted/services/0/name cmdb-api-v1
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-app/weighted/services/0/weight 3
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-app/weighted/services/1/name cmdb-api-v2
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-app/weighted/services/1/weight 1
```

![](./images/weighted.jpg)

配置导出规则:
```
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/routers/cmdb-api-app/rule 'PathPrefix(`/cmdb/api/v1`)'
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/routers/cmdb-api-app/service cmdb-api-app
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/routers/cmdb-api-app/entryPoints/0	web
```

![](./images/weighted-rule.png)

最后验证服务的访问情况

### 健康检查

现在我们的服务实例注册上去了, 一个完整的网关一定要包含实例的检查检查, 这样到后端实例故障的时候，才能正常摘除

traefik为服务提供的监控检查的功能:

![](./images/traefik-hc.png)

```yaml
  services:
    Service01:
      loadBalancer:
        sticky:
          cookie:
            name: foobar
            secure: true
            httpOnly: true
            sameSite: foobar
        servers:
        - url: foobar
        - url: foobar
        healthCheck:
          scheme: foobar
          path: foobar
          port: 42
          interval: foobar
          timeout: foobar
          hostname: foobar
          followRedirects: true
          headers:
            name0: foobar
            name1: foobar
        passHostHeader: true
        responseForwarding:
          flushInterval: foobar
        serversTransport: foobar
```

这是官方的配置样例
```
traefik/http/services/Service01/loadBalancer/healthCheck/followRedirects	true
traefik/http/services/Service01/loadBalancer/healthCheck/headers/name0	foobar
traefik/http/services/Service01/loadBalancer/healthCheck/headers/name1	foobar
traefik/http/services/Service01/loadBalancer/healthCheck/hostname	foobar
traefik/http/services/Service01/loadBalancer/healthCheck/interval	foobar
traefik/http/services/Service01/loadBalancer/healthCheck/path	foobar
traefik/http/services/Service01/loadBalancer/healthCheck/port	42
traefik/http/services/Service01/loadBalancer/healthCheck/scheme	foobar
traefik/http/services/Service01/loadBalancer/healthCheck/timeout
```

我们基于此来配置下app-v1/app-v2 的健康检查
```
# cmdb-api-v1
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-v1/loadBalancer/healthCheck/path	/
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-v1/loadBalancer/healthCheck/interval	5
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-v1/loadBalancer/healthCheck/timeout	1

# cmdb-api-v2
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-v2/loadBalancer/healthCheck/path	/
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-v2/loadBalancer/healthCheck/interval	5
docker exec -it -e "ETCDCTL_API=3" etcd  etcdctl put traefik/http/services/cmdb-api-v2/loadBalancer/healthCheck/timeout	1
```

![](./images/traefik-hce.png)

## 注册中心






