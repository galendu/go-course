package host

type HostSet struct {
	Items []*Host `json:"items"`
	Total int     `json:"total"`
}

type Host struct {
	*Resource
	*Describe
}

type Resource struct {
	Vendor      string            `json:"vendor"`      // 厂商
	Region      string            `json:"region"`      // 地域
	Zone        string            `json:"zone"`        // 区域
	CreatedAt   string            `json:"created_at"`  // 创建时间
	ExpiredAt   string            `json:"expired_at"`  // 过期时间
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
}

type Describe struct {
	CPU                     int    `json:"cpu"`                        // 核数
	Memory                  int    `json:"Memory"`                     // 内存
	GPUAmount               int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec                 string `json:"gpu_spec"`                   // GPU类型
	OSType                  string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName                  string `json:"os_name"`                    // 操作系统名称
	SerialNumber            string `json:"serial_number"`              // 序列号
	PayType                 string `json:"pay_type"`                   // 实例付费方式
	ImageID                 string `json:"image_id"`                   // 镜像ID
	PublicIP                string `json:"public_ip"`                  // 公网IP
	PrivateIP               string `json:"private_ip"`                 // 内网IP
	InternetMaxBandwidthOut int    `json:"internet_max_bandwidth_out"` // 公网出带宽最大值，单位为 Mbps
	InternetMaxBandwidthIn  int    `json:"internet_max_bandwidth_in"`  // 公网入带宽最大值，单位为 Mbps
	KeyPairName             string `json:"key_pair_name"`              // 秘钥对名称
	SecurityGroups          string `json:"security_groups"`            // 安全组  采用逗号分隔
}
