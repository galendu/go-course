package impl

import (
	"context"

	"github.com/rs/xid"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"

	"gitee.com/infraboard/go-course/day14/demo/api/pkg/host"
)

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
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`
	queryHostSQL = `SELECT * FROM resource as r LEFT JOIN host h ON r.id=h.resource_id`
)

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
	s.l.Debugf("sql: %s", querySQL)
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
			&ins.Category, &ins.Type, &ins.InstanceId, &ins.Name, &ins.Description,
			&ins.Status, &ins.UpdateAt, &ins.SyncAt, &ins.SyncAccount,
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
