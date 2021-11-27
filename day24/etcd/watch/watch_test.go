package watch_test

import (
	"testing"

	"gitee.com/infraboard/go-course/day24/etcd/watch"
)

func TestCreate(t *testing.T) {
	watch.UpdateConfig("cmdb v3")
}

func TestDelete(t *testing.T) {
	watch.DeleteConfig()
}
