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
