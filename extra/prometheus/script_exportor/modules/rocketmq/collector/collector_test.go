package collector_test

import (
	"bytes"
	"fmt"
	"testing"

	"gitee.com/infraboard/go-course/extra/prometheus/script_exportor/modules/rocketmq/collector"
	"gitee.com/infraboard/go-course/extra/prometheus/script_exportor/modules/rocketmq/conf"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
)

var (
	retgistry = prometheus.NewRegistry()
)

func TestCollect(t *testing.T) {
	c := collector.NewCollector()
	c.Conf.FileConfig.Path = "../sample/data.txt"

	if err := retgistry.Register(c); err != nil {
		t.Fatal(err)
	}

	mf, err := retgistry.Gather()
	if err != nil {
		t.Fatal(err)
	}

	// 编码输出
	b := bytes.NewBuffer([]byte{})
	enc := expfmt.NewEncoder(b, expfmt.FmtText)
	for i := range mf {
		enc.Encode(mf[i])
	}

	fmt.Println(b.String())
}

func init() {
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
}
