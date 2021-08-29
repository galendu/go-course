# Demo后端

功能: CMDB主机信息录入与查询

涉及到的技能:

+ go http标准库
+ 第三方路由库: [httprouter](https://github.com/julienschmidt/httprouter)
+ go 操作mysql


## 项目骨架介绍

项目组织的核心思路是: 每个业务模块尽量独立, 方便后期扩展和迁移成独立的服务

```
conf  程序配置对象
pkg   业务领域包
  - host
	- model     业务需要的数据模型
	- interface 业务接口(领域方法)
	- impl      业务具体实现
  - mysql
  - lb
  - ...
main 程序入口文件
```

## 数据结构与接口定义

定义数据结构:
```go
package host

const (
	PrivateIDC Vendor = iota
	Tencent
	AliYun
	HuaWei
)

type HostSet struct {
	Items []*Host `json:"items"`
	Total int     `json:"total"`
}

type Host struct {
	*Resource
	*Describe
}

type Vendor int

type Resource struct {
	Id          string            `json:"id"`          // 全局唯一Id
	Vendor      Vendor            `json:"vendor"`      // 厂商
	Region      string            `json:"region"`      // 地域
	Zone        string            `json:"zone"`        // 区域
	CreateAt    int64             `json:"create_at"`   // 创建时间
	ExpireAt    int64             `json:"expire_at"`   // 过期时间
	Category    string            `json:"category"`    // 种类
	Type        string            `json:"type"`        // 规格
	InstanceId  string            `json:"instance_id"` // 实例ID
	Name        string            `json:"name"`        // 名称
	Description string            `json:"description"` // 描述
	Status      string            `json:"status"`      // 服务商中的状态
	Tags        map[string]string `json:"tags"`        // 标签
	UpdateAt    int64             `json:"update_at"`   // 更新时间
	SyncAt      int64             `json:"sync_at"`     // 同步时间
	SyncAccount string            `json:"sync_accout"` // 同步的账号
	PublicIP    string            `json:"public_ip"`   // 公网IP
	PrivateIP   string            `json:"private_ip"`  // 内网IP
	PayType     string            `json:"pay_type"`    // 实例付费方式
}

type Describe struct {
	ResourceId              string `json:"resource_id"`                // 关联Resource
	CPU                     int    `json:"cpu"`                        // 核数
	Memory                  int    `json:"memory"`                     // 内存
	GPUAmount               int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec                 string `json:"gpu_spec"`                   // GPU类型
	OSType                  string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName                  string `json:"os_name"`                    // 操作系统名称
	SerialNumber            string `json:"serial_number"`              // 序列号
	ImageID                 string `json:"image_id"`                   // 镜像ID
	InternetMaxBandwidthOut int    `json:"internet_max_bandwidth_out"` // 公网出带宽最大值，单位为 Mbps
	InternetMaxBandwidthIn  int    `json:"internet_max_bandwidth_in"`  // 公网入带宽最大值，单位为 Mbps
	KeyPairName             string `json:"key_pair_name"`              // 秘钥对名称
	SecurityGroups          string `json:"security_groups"`            // 安全组  采用逗号分隔
}
```

定义业务支持的方法
```go
package host

import (
	"context"
)

type Service interface {
	SaveHost(context.Context, *Host) (*Host, error)
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
}

type QueryHostRequest struct {
	PageSize   uint64 `json:"page_size,omitempty"`
	PageNumber uint64 `json:"page_number,omitempty"`
}
```

## 接口实例定义

接下里就需要实现我们定义的这个Host服务了

定义服务需要实现这个服务的实例: service

```go
package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"gitee.com/infraboard/go-course/day14/demo/api/conf"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	l  logger.Logger
}

func (s *service) Config() error {
	s.l = zap.L().Named("Policy")
	return nil
}
```

定义需要实现的方法:

```go
package impl

import (
	"context"

	"gitee.com/infraboard/go-course/day14/demo/api/pkg/host"
)

func (s *service) SaveHost(context.Context, *host.Host) (*host.Host, error) {
	return nil, nil
}

func (s *service) QueryHost(context.Context, *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}
```

由于我们使用MySQL的存储数据, 因此需要MySQL的配置, 我们把这个配置放置于 配置模块

## 如何管理项目配置

程序提供2中配置加载的方式:
+ 配置文件(toml格式): etc/demo.toml
```toml
[app]
name = "demo"
host = "0.0.0.0"
port = "8050"
key  = "this is your app key"

[mysql]
host = "127.0.0.1"
port = "3306"
username = "go_course"
password = "xxxx"
database = "go_course"

[log]
level = "debug"
path = "logs"
format = "text"
to = "stdout"
``` 
+ 环境变量: etc/demo.env
```sh
export MYSQL_HOST="127.0.0.1"
export MYSQL_PORT="3306"
export MYSQL_USERNAME="go_course"
export MYSQL_PASSWORD="xxx"
export MYSQL_DATABASE="go_course"
```

### 定义配置对象

如果将配置映射成程序里面的对象(比如 Config对象), 这里选用2个第三方库来解决:
+ [toml解析库](https://github.com/BurntSushi/toml)
+ [环境变量解析库](https://github.com/caarlos0/env)

定义我们程序需要的配置对象:
```go
// Config 应用配置
type Config struct {
	App   *app   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

type app struct {
	Name      string `toml:"name" env:"APP_NAME"`
	Host      string `toml:"host" env:"APP_HOST"`
	Port      string `toml:"port" env:"APP_PORT"`
	Key       string `toml:"key" env:"APP_KEY"`
	EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

// Log todo
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

// MySQL todo
type MySQL struct {
	Host        string `toml:"host" env:"D_MYSQL_HOST"`
	Port        string `toml:"port" env:"D_MYSQL_PORT"`
	UserName    string `toml:"username" env:"D_MYSQL_USERNAME"`
	Password    string `toml:"password" env:"D_MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"D_MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" env:"D_MYSQL_MAX_idle_TIME"`
	lock        sync.Mutex
}
```

单独定义下Log相关配置
```go
// LogFormat 日志格式
type LogFormat string

const (
	// TextFormat 文本格式
	TextFormat = LogFormat("text")
	// JSONFormat json格式
	JSONFormat = LogFormat("json")
)

// LogTo 日志记录到哪儿
type LogTo string

const (
	// ToFile 保存到文件
	ToFile = LogTo("file")
	// ToStdout 打印到标准输出
	ToStdout = LogTo("stdout")
)
```

为程序设置一些默认值:

```go
func newConfig() *Config {
	return &Config{
		App:   newDefaultAPP(),
		Log:   newDefaultLog(),
		MySQL: newDefaultMySQL(),
	}
}

func (a *app) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultAPP() *app {
	return &app{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8050",
		Key:  "default",
	}
}

// newDefaultLog todo
func newDefaultLog() *Log {
	return &Log{
		Level:   "debug",
		PathDir: "logs",
		Format:  "text",
		To:      "stdout",
	}
}

// newDefaultMySQL todo
func newDefaultMySQL() *MySQL {
	return &MySQL{
		Database:    "go_course",
		Host:        "127.0.0.1",
		Port:        "3306",
		MaxOpenConn: 200,
		MaxIdleConn: 50,
		MaxLifeTime: 1800,
		MaxIdleTime: 600,
	}
}
```

### 配置加载

+ 配置对象定义好了
+ 配置文件也准备好了

接下来就需要完成配置的加载, 分别为不同的配置提供不同的加载方法

```go
package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	global *Config
)

// C 全局配置对象
func C() *Config {
	if global == nil {
		panic("Load Config first")
	}
	return global
}

// LoadConfigFromToml 从toml中添加配置文件, 并初始化全局对象
func LoadConfigFromToml(filePath string) error {
	cfg := newConfig()
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return err
	}
	// 加载全局配置单例
	global = cfg
	return nil
}

// LoadConfigFromEnv 从环境变量中加载配置
func LoadConfigFromEnv() error {
	cfg := newConfig()
	if err := env.Parse(cfg); err != nil {
		return err
	}
	// 加载全局配置单例
	global = cfg
	return nil
}
```

### 其他需要全局单实例的配置

由于数据的链接 其他服务都要使用, 这里也做成了全局单例

```go
var (
	db *sql.DB
)

// getDBConn use to get db connection pool
func (m *MySQL) getDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}

// GetDB todo
func (m *MySQL) GetDB() (*sql.DB, error) {
	// 加载全局数据量单例
	m.lock.Lock()
	defer m.lock.Unlock()
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}
```

这样我们就可以通过Conf,获取全局单例, 比如获取数据库连接
```go
db, err := conf.C().MySQL.GetDB()
if err != nil {
	return err
}
```

## 基于MySQL存储实现实例

为我们service 补充mysql链接依赖:

```go
type service struct {
	db *sql.DB
	l  logger.Logger
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.l = zap.L().Named("Host")
	s.db = db
	return nil
}
```

### 数据库表设计

数据库我自己搭建了一套MySQL, 你们可以选择自己搭建，也可以使用我搭建的,到时候发群里面

连接数据库的工具这里使用 Navicat, (个人需求: 因为他也可以连mongo), 你们可以根据自己喜好选择趁手的工具

这里我们设计2张表, 一张用于存储Resource通用信息, 方便我们全局快速解索资源

设计注意事项:
+ 注意存储引擎的选择(InnoDB)
+ 主键选择 和 唯一键考虑清楚
+ 考虑数据类型与长度，选择合适的类型，避免空间浪费
+ 字符串注意确认字符集, 如果需要存入中文, 请选择utf8编码
+ 为过滤条件的字段 添加索引
+ 高频组合查询可以考虑 联合索引
+ 注意选择使用索引的方法: Hash Btree Normal

需要添加索引的字段
```
instance_id  Hash
name         Btree
status       Hash
private_ip   Btree
public_ip    Btree
```

最后resource表的创建SQL:
```sql
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` char(64) CHARACTER SET latin1 NOT NULL,
  `vendor` tinyint(1) NOT NULL,
  `region` varchar(64) CHARACTER SET latin1 NOT NULL,
  `zone` varchar(64) CHARACTER SET latin1 NOT NULL,
  `create_at` bigint(13) NOT NULL,
  `expire_at` bigint(13) DEFAULT NULL,
  `category` varchar(64) CHARACTER SET latin1 NOT NULL,
  `type` varchar(120) CHARACTER SET latin1 NOT NULL,
  `instance_id` varchar(120) CHARACTER SET latin1 NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status` varchar(255) CHARACTER SET latin1 NOT NULL,
  `update_at` bigint(13) DEFAULT NULL,
  `sync_at` bigint(13) DEFAULT NULL,
  `sync_accout` varchar(255) CHARACTER SET latin1 DEFAULT NULL,
  `public_ip` varchar(64) CHARACTER SET latin1 DEFAULT NULL,
  `private_ip` varchar(64) CHARACTER SET latin1 DEFAULT NULL,
  `pay_type` varchar(255) CHARACTER SET latin1 DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `instance_id` (`instance_id`) USING HASH,
  KEY `name` (`name`) USING BTREE,
  KEY `status` (`status`) USING HASH,
  KEY `private_ip` (`public_ip`) USING BTREE,
  KEY `public_ip` (`public_ip`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
```

host表SQL如下:
```sql
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for host
-- ----------------------------
DROP TABLE IF EXISTS `host`;
CREATE TABLE `host` (
  `resource_id` varchar(64) NOT NULL,
  `cpu` tinyint(4) NOT NULL,
  `mem` int(13) NOT NULL,
  `gpu_amount` tinyint(4) DEFAULT NULL,
  `gpu_spec` varchar(255) DEFAULT NULL,
  `os_type` varchar(255) DEFAULT NULL,
  `os_name` varchar(255) DEFAULT NULL,
  `serial_number` varchar(120) DEFAULT NULL,
  `image_id` char(64) DEFAULT NULL,
  `internet_max_bandwidth_out` int(10) DEFAULT NULL,
  `internet_max_bandwidth_in` int(10) DEFAULT NULL,
  `key_pair_name` varchar(255) DEFAULT NULL,
  `security_groups` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

SET FOREIGN_KEY_CHECKS = 1;
```

### 实现存储和查询接口

定义Insert和Select语句
```go
const (
	insertResourceSQL = `INSERT INTO resource (
		id,vendor,region,zone,create_at,expire_at,category,type,instance_id,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	insertHostSQL = `INSERT INTO host (
		resource_id,cpu,memory,gpu_amount,gpu_spec,os_type,os_name,
		serial_number,image_id,internet_max_bandwidth_out,
		internet_max_bandwidth_in,key_pair_name,security_groups
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?);`
	queryHostSQL = `SELECT * FROM resource as r LEFT JOIN host h ON r.id=h.resource_id;`
)
```

操作数据库过程:
+ prepare stmt
+ stmt.Exec
+ 查询时需要使用sqlbuilder(自己简单实现)

```go
func (s *service) SaveHost(ctx context.Context, h *host.Host) (*host.Host, error) {
	h.Id = xid.New().String()
	h.ResourceId = h.Id

	// 避免SQL注入, 请使用Prepare
	stmt, err := s.db.Prepare(insertResourceSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		h.Id, h.Vendor, h.Region, h.Zone, h.CreateAt, h.ExpireAt, h.Category, h.Type, h.InstanceId,
		h.Name, h.Description, h.Status, h.UpdateAt, h.SyncAt, h.SyncAccount, h.PublicIP,
		h.PrivateIP, h.PayType,
	)
	if err != nil {
		return nil, err
	}

	// 避免SQL注入, 请使用Prepare
	stmt, err = s.db.Prepare(insertHostSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		h.ResourceId, h.CPU, h.Memory, h.GPUAmount, h.GPUSpec, h.OSType, h.OSName,
		h.SerialNumber, h.ImageID, h.InternetMaxBandwidthOut,
		h.InternetMaxBandwidthIn, h.KeyPairName, h.SecurityGroups,
	)
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (s *service) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	query := sqlbuilder.NewQuery(queryHostSQL)
	querySQL, args := query.Order("create_at").Desc().Limit(req.OffSet(), uint(req.PageSize)).BuildQuery()
	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query job task error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := host.NewHostSet()
	for rows.Next() {
		ins := host.NewDefaultHost()
		err := rows.Scan(
			&ins.Id, &ins.Vendor, &ins.Region, &ins.Zone, &ins.CreateAt, &ins.ExpireAt,
			&ins.Category, &ins.Category, &ins.Type, &ins.InstanceId, &ins.Name,
			&ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt, &ins.SyncAccount,
			&ins.PublicIP, &ins.PrivateIP, &ins.PayType, &ins.ResourceId, &ins.CPU,
			&ins.Memory, &ins.GPUAmount, &ins.GPUSpec, &ins.OSType, &ins.OSName,
			&ins.SerialNumber, &ins.ImageID, &ins.InternetMaxBandwidthOut, &ins.InternetMaxBandwidthIn,
			&ins.KeyPairName, &ins.SecurityGroups,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query job task error, %s", err.Error())
		}
		set.Add(ins)
	}

	// 获取total
	countSQL, args := query.BuildCount()
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	return set, nil
}
```

到此我们实现了数据的存储与分页查询

## HTTP API暴露

我们以RestFull风格来设计我们的接口

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func QueryUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "query user!\n")
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "create user!\n")
}

func DescribeUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "describe, %s!\n", ps.ByName("name"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "delete, %s!\n", ps.ByName("name"))
}

var (
	addr = ":8080"
)

func main() {
	router := httprouter.New()
	router.GET("/users", QueryUser)
	router.POST("/users", CreateUser)
	router.GET("/users/:id", DescribeUser)
	router.DELETE("/users/:id", DeleteUser)

	log.Printf("listen on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
```





