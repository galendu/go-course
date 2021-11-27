# Workflow API Server

![](./images/ci_cd.png)

API Server的核心职责比较简单: 就是往etcd里面写数据, 写什么数据，怎么写, 这就是我们的业务逻辑


## workflow 组件与流程

+ API Server: 负责将Pipeline对象写入Etcd
+ Scheduler: 负责调度Pipeline定义的任务到具体的Node节点执行
+ Node: Watch做任务, 发现有任务调度给自己后, 执行任务

因此我们的项目骨架如下:

+ api: api server 项目
    + app: app模块层
    + client: api server grpc 客户端
    + cmd: api server cli
    + protocol: api server 暴露的API服务,包含 grpc和http
+ scheduler: scheduler项目
    + algorithm: 调度算法包
    + cmd: 调度器 cli
    + controller: 调度器的控制器, watch list 对象变化
        + cronjob: 赋值cronjob类型任务的调度，预留未实现
        + node: node对象的控制器, watch 注册上的Node节点, 方便调度器选择执行的node
        + pipeline: watch pipeline对象变化, 赋值控制pipeline状态变化, 负责pipeline 中 task任务的创建, 创建后由 task调度器负责调度
        + step: 负责watch step对象的变化, 将具体的step 调度给对应的node节点执行
+ node: node项目
    + cmd: node cli
    + controller: node节点控制器, 负责watch 对象变化
        + step: 服务watch step 状态变化
            + engine: 控制调用runner来执行任务, 并管理所有任务
            + runner: 负责执行具体的认为
                + docker: 负责调用docker来执行任务
                + k8s: 负责调用k8s来执行任务
                + local: 负责在本地执行任务
            + store: 负责记录task 执行日志
+ conf: 整个项目的配置文件
+ version: 整个项目的版本 
+ common: 项目通用工具包