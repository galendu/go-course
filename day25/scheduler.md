# 调度器开发

调度器的核心逻辑:
+ Watch Pipeline对象, 监听变化事件
+ 当有新的Pipeline对象被创建时, 修改Pipeline对象的 Scheduler属性, 为其挑选一个可以Scheduler Node来处理Pipline

![](./images/k8s-watch-list.png)

因此第一步是编写Informer, Watch Pipeline对象的变化

## PIpeline Informer

### 定义接口

我们先定义Informer的接口
```go
package pipeline

import (
	"context"

	"github.com/infraboard/workflow/api/app/pipeline"
	"github.com/infraboard/workflow/common/cache"
)

// Informer 负责事件通知
type Informer interface {
	Watcher() Watcher
	Lister() Lister
	Recorder() Recorder
	GetStore() cache.Store
}

type Lister interface {
	List(ctx context.Context, opts *pipeline.QueryPipelineOptions) (*pipeline.PipelineSet, error)
}

type Recorder interface {
	Update(*pipeline.Pipeline) error
}

// Watcher 负责事件通知
type Watcher interface {
	// Run starts and runs the shared informer, returning after it stops.
	// The informer will be stopped when stopCh is closed.
	Run(ctx context.Context) error
	// AddEventHandler adds an event handler to the shared informer using the shared informer's resync
	// period.  Events to a single handler are delivered sequentially, but there is no coordination
	// between different handlers.
	AddPipelineTaskEventHandler(handler PipelineEventHandler)
}

// PipelineEventHandler can handle notifications for events that happen to a
// resource. The events are informational only, so you can't return an
// error.
//  * OnAdd is called when an object is added.
//  * OnUpdate is called when an object is modified. Note that oldObj is the
//      last known state of the object-- it is possible that several changes
//      were combined together, so you can't use this to see every single
//      change. OnUpdate is also called when a re-list happens, and it will
//      get called even if nothing changed. This is useful for periodically
//      evaluating or syncing something.
//  * OnDelete will get the final state of the item if it is known, otherwise
//      it will get an object of type DeletedFinalStateUnknown. This can
//      happen if the watch is closed and misses the delete event and we don't
//      notice the deletion until the subsequent re-list.
type PipelineEventHandler interface {
	OnAdd(obj *pipeline.Pipeline)
	OnUpdate(old, new *pipeline.Pipeline)
	OnDelete(obj *pipeline.Pipeline)
}

// PipelineEventHandlerFuncs is an adaptor to let you easily specify as many or
// as few of the notification functions as you want while still implementing
// ResourceEventHandler.
type PipelineTaskEventHandlerFuncs struct {
	AddFunc    func(obj *pipeline.Pipeline)
	UpdateFunc func(oldObj, newObj *pipeline.Pipeline)
	DeleteFunc func(obj *pipeline.Pipeline)
}

// OnAdd calls AddFunc if it's not nil.
func (r PipelineTaskEventHandlerFuncs) OnAdd(obj *pipeline.Pipeline) {
	if r.AddFunc != nil {
		r.AddFunc(obj)
	}
}

// OnUpdate calls UpdateFunc if it's not nil.
func (r PipelineTaskEventHandlerFuncs) OnUpdate(oldObj, newObj *pipeline.Pipeline) {
	if r.UpdateFunc != nil {
		r.UpdateFunc(oldObj, newObj)
	}
}

// OnDelete calls DeleteFunc if it's not nil.
func (r PipelineTaskEventHandlerFuncs) OnDelete(obj *pipeline.Pipeline) {
	if r.DeleteFunc != nil {
		r.DeleteFunc(obj)
	}
}

type PipelineFilterHandler func(obj *pipeline.Pipeline) error
```

### Etcd实现

结下来我们使用etcd 来实现定义


#### 实现Lister

lister其实就是一个带 prefix的 Get操作

```go
type lister struct {
	log    logger.Logger
	client clientv3.KV
}

func (l *lister) List(ctx context.Context, opts *pipeline.QueryPipelineOptions) (*pipeline.PipelineSet, error) {
	listKey := pipeline.EtcdPipelinePrefix()
	l.log.Infof("list etcd pipeline resource key: %s", listKey)
	resp, err := l.client.Get(ctx, listKey, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	ps := pipeline.NewPipelineSet()
	for i := range resp.Kvs {
		// 解析对象
		pt, err := pipeline.LoadPipelineFromBytes(resp.Kvs[i].Value)
		if err != nil {
			l.log.Errorf("load pipeline [key: %s, value: %s] error, %s", resp.Kvs[i].Key, string(resp.Kvs[i].Value), err)
			continue
		}

		pt.ResourceVersion = resp.Header.Revision
		ps.Add(pt)
	}
	return ps, nil
}
```

#### 实现Recoder

recorder 其实就是一个 put操作

```go
type recorder struct {
	log    logger.Logger
	client clientv3.KV
}

func (l *recorder) Update(t *pipeline.Pipeline) error {
	if t == nil {
		return fmt.Errorf("update nil pipeline")
	}

	if l.client == nil {
		return fmt.Errorf("etcd client is nil")
	}

	objKey := t.EtcdObjectKey()
	objValue, err := json.Marshal(t)
	if err != nil {
		return err
	}
	if _, err := l.client.Put(context.Background(), objKey, string(objValue)); err != nil {
		return fmt.Errorf("update pipeline task '%s' to etcd3 failed: %s", objKey, err.Error())
	}
	return nil
}
```


#### 实现Watcher

watcher的逻辑稍微复杂些:
+ 需要Watch 所有Pipeline对象
+ 处理事件, 把对象缓存到本地的indexer中, 这里Indexer就是一个内存缓存
+ 调用Event处理函数，把事件传播出去


注意Watch是一个阻塞操作, 所以Watcher 是几个里面为一个需要Run的

```go
type shared struct {
	log       logger.Logger
	client    clientv3.Watcher
	indexer   cache.Indexer
	handler   informer.PipelineEventHandler
	filter    informer.PipelineFilterHandler
	watchChan clientv3.WatchChan
}

// AddPipelineEventHandler 添加事件处理回调
func (i *shared) AddPipelineTaskEventHandler(h informer.PipelineEventHandler) {
	i.handler = h
}

// Run 启动 Watch
func (i *shared) Run(ctx context.Context) error {
	// 是否准备完成
	if err := i.isReady(); err != nil {
		return err
	}

	// 监听事件
	i.watch(ctx)

	// 后台处理事件
	go i.dealEvents()
	return nil
}

func (i *shared) dealEvents() {
	// 处理所有事件
	for {
		select {
		case nodeResp := <-i.watchChan:
			for _, event := range nodeResp.Events {
				switch event.Type {
				case mvccpb.PUT:
					if err := i.handlePut(event, nodeResp.Header.GetRevision()); err != nil {
						i.log.Error(err)
					}
				case mvccpb.DELETE:
					if err := i.handleDelete(event); err != nil {
						i.log.Error(err)
					}
				default:
				}
			}
		}
	}
}

func (i *shared) isReady() error {
	if i.handler == nil {
		return errors.New("PipelineEventHandler not add")
	}
	return nil
}

func (i *shared) watch(ctx context.Context) {
	ppWatchKey := pipeline.EtcdPipelinePrefix()
	i.watchChan = i.client.Watch(ctx, ppWatchKey, clientv3.WithPrefix())
	i.log.Infof("watch etcd pipeline resource key: %s", ppWatchKey)
}

func (i *shared) handlePut(event *clientv3.Event, eventVersion int64) error {
	i.log.Debugf("receive pipeline put event, %s", event.Kv.Key)

	// 解析对象
	new, err := pipeline.LoadPipelineFromBytes(event.Kv.Value)
	if err != nil {
		return err
	}
	new.ResourceVersion = eventVersion

	old, hasOld, err := i.indexer.GetByKey(new.MakeObjectKey())
	if err != nil {
		return err
	}

	if i.filter != nil {
		if err := i.filter(new); err != nil {
			return err
		}
	}

	// 区分Update
	if hasOld {
		// 更新缓存
		i.log.Debugf("update pipeline: %s", new.ShortDescribe())
		if err := i.indexer.Update(new); err != nil {
			i.log.Errorf("update indexer cache error, %s", err)
		}
		i.handler.OnUpdate(old.(*pipeline.Pipeline), new)
	} else {
		// 添加缓存
		i.log.Debugf("add pipeline: %s", new.ShortDescribe())
		if err := i.indexer.Add(new); err != nil {
			i.log.Errorf("add indexer cache error, %s", err)
		}
		i.handler.OnAdd(new)
	}

	return nil
}

func (i *shared) handleDelete(event *clientv3.Event) error {
	key := event.Kv.Key
	i.log.Debugf("receive pipeline delete event, %s", key)

	obj, ok, err := i.indexer.GetByKey(string(key))
	if err != nil {
		i.log.Errorf("get key %s from store error, %s", key)
	}
	if !ok {
		i.log.Warnf("key %s found in store", key)
	}

	// 清除缓存
	if err := i.indexer.Delete(obj); err != nil {
		i.log.Errorf("delete indexer cache error, %s", err)
	}

	i.handler.OnDelete(obj.(*pipeline.Pipeline))
	return nil
}
```


#### 实现Informer

最后我们把这些组合起来我们的Informer就完成了

```go
// NewScheduleInformer todo
func NewInformerr(client *clientv3.Client, filter informer.PipelineFilterHandler) informer.Informer {
	return &Informer{
		log:     zap.L().Named("Pipeline"),
		client:  client,
		filter:  filter,
		indexer: cache.NewIndexer(informer.MetaNamespaceKeyFunc, informer.DefaultStoreIndexers()),
	}
}

// Informer todo
type Informer struct {
	log      logger.Logger
	client   *clientv3.Client
	shared   *shared
	lister   *lister
	recorder *recorder
	indexer  cache.Indexer
	filter   informer.PipelineFilterHandler
}

func (i *Informer) GetStore() cache.Store {
	return i.indexer
}

func (i *Informer) Debug(l logger.Logger) {
	i.log = l
	i.shared.log = l
	i.lister.log = l
}

func (i *Informer) Watcher() informer.Watcher {
	if i.shared != nil {
		return i.shared
	}
	i.shared = &shared{
		log:     i.log.Named("Watcher"),
		client:  clientv3.NewWatcher(i.client),
		indexer: i.indexer,
		filter:  i.filter,
	}
	return i.shared
}

func (i *Informer) Lister() informer.Lister {
	if i.lister != nil {
		return i.lister
	}
	i.lister = &lister{
		log:    i.log.Named("Lister"),
		client: clientv3.NewKV(i.client),
	}
	return i.lister
}

func (i *Informer) Recorder() informer.Recorder {
	if i.recorder != nil {
		return i.recorder
	}
	i.recorder = &recorder{
		log:    i.log.Named("Recorder"),
		client: clientv3.NewKV(i.client),
	}
	return i.recorder
}
```


#### 测试

基于etcdcli可以测试我们Informer功能是否正常:
+ 测试Lister的逻辑
+ 测试Watcher的逻辑
+ 测试Recorder的逻辑

## 服务注册

既然需要调度, 因此我们的Node节点需要注册到中央来，我们才能知道如何调度

### 服务接口

我们先关注:
+ 服务的接口规范
+ 服务的数据结构

我们把workflow的服务抽象成了接口:
```go
type Register interface {
	Debug(logger.Logger)
	Registe() error
	UnRegiste() error
}
```

定义workflow服务的类型:
```go
const (
	// API 提供API访问的服务
	APIType = Type("api")
	// Worker 后台作业服务
	NodeType = Type("node")
	// Scheduler 调度器
	SchedulerType = Type("scheduler")
)
```

Node数据结构 用于泛指一个服务
```go
// Node todo
type Node struct {
	Region          string            `json:"region,omitempty"`
	ResourceVersion int64             `json:"resourceVersion,omitempty"`
	InstanceName    string            `json:"instance_name,omitempty"`
	ServiceName     string            `json:"service_name,omitempty"`
	Type            Type              `json:"type,omitempty"`
	Address         string            `json:"address,omitempty"`
	Version         string            `json:"version,omitempty"`
	GitBranch       string            `json:"git_branch,omitempty"`
	GitCommit       string            `json:"git_commit,omitempty"`
	BuildEnv        string            `json:"build_env,omitempty"`
	BuildAt         string            `json:"build_at,omitempty"`
	Online          int64             `json:"online,omitempty"`
	Tag             map[string]string `json:"tag,omitempty"`

	Prefix   string        `json:"-"`
	Interval time.Duration `json:"-"`
	TTL      int64         `json:"-"`
}
```

### 基于Etcd的注册中心

我们通过etcd来实现注册器

etcd有个租约的概念, 我们可以通过租约来控制一个key 的TTL, 我们基于此来实现注册的心跳功能

+ 往etcd里面写一个服务的key/value，并通过租约设置TTL
+ 每隔一个心跳周期，就KeepAliveOnce 把改租约 续约一次, 也就是心跳机制
+ 最好服务停止时，主动注销服务

1. 初次注册

+ key结构: inforboard/workflow/service/scheduler/{name_name}
+ value: node结构体json数据

```go
func (e *etcd) addOnce() error {
	// 获取leaseID
	resp, err := e.client.Lease.Grant(context.TODO(), e.node.TTL)
	if err != nil {
		return fmt.Errorf("get etcd lease id error, %s", err)
	}
	e.leaseID = resp.ID

	// 写入key
	if _, err := e.client.Put(context.Background(), e.instanceKey, e.instanceValue, clientv3.WithLease(e.leaseID)); err != nil {
		return fmt.Errorf("registe service '%s' with ttl to etcd3 failed: %s", e.instanceKey, err.Error())
	}
	e.instanceKey = e.instanceKey
	return nil
}
```

2. 续约

```go
func (e *etcd) keepAlive(ctx context.Context) {
	// 不停的续约
	interval := e.node.TTL / 5
	e.Infof("keep alive lease interval is %d second", interval)
	tk := time.NewTicker(time.Duration(interval) * time.Second)
	defer tk.Stop()
	for {
		select {
		case <-ctx.Done():
			e.Infof("keepalive goroutine exit")
			return
		case <-tk.C:
			Opctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			_, err := e.client.Lease.KeepAliveOnce(Opctx, e.leaseID)
			if err != nil {
				if strings.Contains(err.Error(), "requested lease not found") {
					// 避免程序卡顿造成leaseID失效(比如mac 电脑休眠))
					if err := e.addOnce(); err != nil {
						e.Errorf("refresh registry error, %s", err)
					} else {
						e.Warn("refresh registry success")
					}
				}
				e.Errorf("lease keep alive error, %s", err)
			} else {
				e.Debugf("lease keep alive key: %s", e.instanceKey)
			}
		}
	}
}
```

3. 注销

+ 删除注册上去的服务实例信息
+ 删除租约
+ 停止KeepAlive续约Goroutine

```go
// UnRegiste delete nodeed service from etcd, if etcd is down
// unnode while timeout.
func (e *etcd) UnRegiste() error {
	if e.isStopped {
		return errors.New("the instance has unregisted")
	}
	// delete instance key
	e.stopInstance <- struct{}{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if resp, err := e.client.Delete(ctx, e.instanceKey); err != nil {
		e.Warnf("unregiste '%s' failed: connect to etcd server timeout, %s", e.instanceKey, err.Error())
	} else {
		if resp.Deleted == 0 {
			e.Infof("unregiste '%s' failed, the key not exist", e.instanceKey)
		} else {
			e.Infof("服务实例(%s)注销成功", e.instanceKey)
		}
	}
	// revoke lease
	_, err := e.client.Lease.Revoke(context.TODO(), e.leaseID)
	if err != nil {
		e.Warnf("revoke lease error, %s", err)
		return err
	}
	e.isStopped = true
	// 停止续约心态
	e.keepAliveStop()
	return nil
}
```

### 服务启动时注册

在服务启动时，调用etcd 注册器来将服务实例信息注册到 注册到etcd

cmd/start.go
```go
// 注册服务
r, err := etcd_register.NewEtcdRegister(svr.node)
if err != nil {
	svr.log.Warn(err)
}
r.Debug(zap.L().Named("Register"))
defer r.UnRegiste()
if err := r.Registe(); err != nil {
	return err
}
```


### 测试

+ 测试正常流程，使用etcdctl 查看 etcd里面改服务的实例是否存在
+ 测试TTL超时，不能完成续约的情况

## Pipeline Controller

Pipeline Controller的核心逻辑:
+ 挑选一个Node


### Picker接口

这里我们有这种资源需要被调度:
+ Pipeline: 多个调度器中挑选一个来处理Pipeline
+ Step: 多个Node中挑选一个执行Step

下面就是Node的分类:
```go
const (
	// API 提供API访问的服务
	APIType = Type("api")
	// Worker 后台作业服务
	NodeType = Type("node")
	// Scheduler 调度器
	SchedulerType = Type("scheduler")
)

type Type string

type Node struct {
	Region          string            `json:"region,omitempty"`
	ResourceVersion int64             `json:"resourceVersion,omitempty"`
	InstanceName    string            `json:"instance_name,omitempty"`
	ServiceName     string            `json:"service_name,omitempty"`
	Type            Type              `json:"type,omitempty"`
	Address         string            `json:"address,omitempty"`
	Version         string            `json:"version,omitempty"`
	GitBranch       string            `json:"git_branch,omitempty"`
	GitCommit       string            `json:"git_commit,omitempty"`
	BuildEnv        string            `json:"build_env,omitempty"`
	BuildAt         string            `json:"build_at,omitempty"`
	Online          int64             `json:"online,omitempty"`
	Tag             map[string]string `json:"tag,omitempty"`

	Prefix   string        `json:"-"`
	Interval time.Duration `json:"-"`
	TTL      int64         `json:"-"`
}
```

接口定义:
```go
// Picker 挑选一个合适的node 运行Step
type StepPicker interface {
	Pick(*pipeline.Step) (*node.Node, error)
}

type PipelinePicker interface {
	Pick(*pipeline.Pipeline) (*node.Node, error)
}
```

### Roundrobin Picker

Picker就是我们挑选Node的算法, 因此我们先实现一个最简单的算法: RR

step picker:
```go
type roundrobinPicker struct {
	mu    *sync.Mutex
	next  int
	store cache.Store
}

// NewStepPicker 实现分调度
func NewStepPicker(nodestore cache.Store) (algorithm.StepPicker, error) {
	base := &roundrobinPicker{
		store: nodestore,
		mu:    new(sync.Mutex),
		next:  0,
	}
	return &stepPicker{base}, nil
}

type stepPicker struct {
	*roundrobinPicker
}

func (p *stepPicker) Pick(step *pipeline.Step) (*node.Node, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	nodes := p.store.List()
	if len(nodes) == 0 {
		return nil, fmt.Errorf("has no available nodes")
	}

	n := nodes[p.next]

	// 修改状态
	p.next = (p.next + 1) % p.store.Len()

	return n.(*node.Node), nil
}
```

pipeline picker:
```go
// NewPipelinePicker 实现分调度
func NewPipelinePicker(nodestore cache.Store) (algorithm.PipelinePicker, error) {
	base := &roundrobinPicker{
		store: nodestore,
		mu:    new(sync.Mutex),
		next:  0,
	}
	return &pipelinePicker{base}, nil
}

type pipelinePicker struct {
	*roundrobinPicker
}

func (p *pipelinePicker) Pick(step *pipeline.Pipeline) (*node.Node, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	nodes := p.store.List()
	if len(nodes) == 0 {
		return nil, fmt.Errorf("has no available nodes")
	}

	n := nodes[p.next]

	// 修改状态
	p.next = (p.next + 1) % p.store.Len()

	return n.(*node.Node), nil
}
```


### Controller

controller 启动逻辑:
+ Watch:  在Controller启动之前， watcher已经启动
+ List: Controller 首先Sync一下，拉去当前所有Pipeline，判定是否有需要处理的，添加到工作队列(work queue)
+ Worker: 然后启动worker，处理工作队列里面的事件(worker queue)
+ Sgin: 等待Controller 结束


```go
// PipelineTaskScheduler 调度器控制器
type Controller struct {
	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue      workqueue.RateLimitingInterface
	informer       informer.Informer
	step           step.Informer
	log            logger.Logger
    // 启动多个worker来处理事件
	workerNums     int
    // 当前running中的worker
	runningWorkers map[string]bool
	wLock          sync.Mutex
    // 调度挑选算法
	picker         algorithm.PipelinePicker
    // 调度器的名称
	schedulerName  string
}
```

下面是起点逻辑:
```go
// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) run(ctx context.Context, async bool) error {
	// Start the informer factories to begin populating the informer caches
	c.log.Infof("starting pipeline control loop, schedule name: %s", c.schedulerName)

	// 获取所有的pipeline
	if err := c.sync(ctx); err != nil {
		return err
	}

	// 启动worker 处理来自Informer的事件
	for i := 0; i < c.workerNums; i++ {
		go c.runWorker(fmt.Sprintf("worker-%d", i))
	}

	if async {
		go c.waitDown(ctx)
	} else {
		c.waitDown(ctx)
	}

	return nil
}
```

下面是sync的逻辑：
```go

func (c *Controller) sync(ctx context.Context) error {
	// 获取所有的pipeline
	listCount := 0
	ps, err := c.informer.Lister().List(ctx, nil)
	if err != nil {
		return err
	}

	// 看看是否有需要调度的
	for i := range ps.Items {
		p := ps.Items[i]

        // 判定Pipeline是否已经执行完成, 已经完成的Pipeline无效处理
        // 由此可见，我们Etcd里面是不适合存储大量历史数据的
		if p.IsComplete() {
			c.log.Debugf("pipline %s is complete, skip schedule",
				p.ShortDescribe())
			continue
		}

        // 判定改Pipeline是否需要当前调度器处理
        // 这里是多个Controller竞争一个Pipeline调度
		if !p.MatchScheduler(c.schedulerName) {
			c.log.Debugf("pipeline %s scheduler %s is not match this scheduler %s",
				p.ShortDescribe(), p.ScheduledNodeName(), c.schedulerName)
			continue
		}

        // 如果都不是，这添加到工作队列，等待调度
		c.enqueueForAdd(p)
		listCount++
	}
	c.log.Infof("%d pipeline need schedule", listCount)
	return nil
}
```

我们继续看Run worker的逻辑:
```go
// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *Controller) runWorker(name string) {
	isRunning, ok := c.runningWorkers[name]
	if ok && isRunning {
		c.log.Warnf("worker %s has running", name)
		return
	}
	c.wLock.Lock()
	c.runningWorkers[name] = true
	c.log.Infof("start worker %s", name)
	c.wLock.Unlock()
	for c.processNextWorkItem() {
	}
	if isRunning, ok = c.runningWorkers[name]; ok {
		c.wLock.Lock()
		delete(c.runningWorkers, name)
		c.wLock.Unlock()
		c.log.Infof("worker %s has stopped", name)
	}
}
```