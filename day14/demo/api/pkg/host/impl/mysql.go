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
	db *sql.DB
	l  logger.Logger
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.l = zap.L().Named("Policy")
	s.db = db
	return nil
}
