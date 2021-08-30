package impl

import (
	"context"
	"database/sql"

	"github.com/rs/xid"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"

	"gitee.com/infraboard/go-course/day14/demo/api/pkg/host"
)

const (
	insertResourceSQL = `INSERT INTO resource (
		id,vendor,region,zone,create_at,expire_at,category,type,instance_id,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type,describe_hash
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	insertHostSQL = `INSERT INTO host (
		resource_id,cpu,memory,gpu_amount,gpu_spec,os_type,os_name,
		serial_number,image_id,internet_max_bandwidth_out,
		internet_max_bandwidth_in,key_pair_name,security_groups
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`

	queryHostSQL      = `SELECT * FROM resource as r LEFT JOIN host h ON r.id=h.resource_id`
	deleteHostSQL     = `DELETE FROM host WHERE resource_id = ?;`
	deleteResourceSQL = `DELETE FROM resource WHERE id = ?;`
)

func (s *service) SaveHost(ctx context.Context, h *host.Host) (*host.Host, error) {
	var (
		stmt *sql.Stmt
		err  error
	)

	h.Id = xid.New().String()
	h.ResourceId = h.Id

	// 开启一个事物
	// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
	// 关于事物级别可以参考文章: https://zhuanlan.zhihu.com/p/117476959
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 执行结果提交或者回滚事务
	// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// 避免SQL注入, 请使用Prepare
	stmt, err = tx.Prepare(insertResourceSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		h.Id, h.Vendor, h.Region, h.Zone, h.CreateAt, h.ExpireAt, h.Category, h.Type, h.InstanceId,
		h.Name, h.Description, h.Status, h.UpdateAt, h.SyncAt, h.SyncAccount, h.PublicIP,
		h.PrivateIP, h.PayType, h.DescribeHash,
	)
	if err != nil {
		return nil, err
	}

	// 生成描写信息的Hash
	if err := h.GenHash(); err != nil {
		return nil, err
	}

	// 避免SQL注入, 请使用Prepare
	stmt, err = tx.Prepare(insertHostSQL)
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
		return nil, exception.NewInternalServerError("prepare query host error, %s", err.Error())
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
			&ins.PublicIP, &ins.PrivateIP, &ins.PayType, &ins.DescribeHash, &ins.ResourceId,
			&ins.CPU, &ins.Memory, &ins.GPUAmount, &ins.GPUSpec, &ins.OSType, &ins.OSName,
			&ins.SerialNumber, &ins.ImageID, &ins.InternetMaxBandwidthOut, &ins.InternetMaxBandwidthIn,
			&ins.KeyPairName, &ins.SecurityGroups,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query host error, %s", err.Error())
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

func (s *service) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (
	*host.Host, error) {
	query := sqlbuilder.NewQuery(queryHostSQL)
	querySQL, args := query.Where("id = ?", req.Id).BuildQuery()
	s.l.Debugf("sql: %s", querySQL)
	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query host error, %s", err.Error())
	}
	defer queryStmt.Close()

	ins := host.NewDefaultHost()
	err = queryStmt.QueryRow(args...).Scan(
		&ins.Id, &ins.Vendor, &ins.Region, &ins.Zone, &ins.CreateAt, &ins.ExpireAt,
		&ins.Category, &ins.Type, &ins.InstanceId, &ins.Name, &ins.Description,
		&ins.Status, &ins.UpdateAt, &ins.SyncAt, &ins.SyncAccount,
		&ins.PublicIP, &ins.PrivateIP, &ins.PayType, &ins.DescribeHash, &ins.ResourceId,
		&ins.CPU, &ins.Memory, &ins.GPUAmount, &ins.GPUSpec, &ins.OSType, &ins.OSName,
		&ins.SerialNumber, &ins.ImageID, &ins.InternetMaxBandwidthOut, &ins.InternetMaxBandwidthIn,
		&ins.KeyPairName, &ins.SecurityGroups,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFound("%#v not found", req)
		}
		return nil, exception.NewInternalServerError("describe host error, %s", err.Error())
	}

	return ins, nil
}

func (s *service) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (
	*host.Host, error) {
	var (
		stmt *sql.Stmt
		err  error
	)

	ins, err := s.DescribeHost(ctx, host.NewDescribeHostRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	// 开启一个事物
	// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
	// 关于事物级别可以参考文章: https://zhuanlan.zhihu.com/p/117476959
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 执行结果提交或者回滚事务
	// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	stmt, err = tx.Prepare(deleteHostSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Id)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	stmt, err = s.db.Prepare(deleteResourceSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Id)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	return ins, nil
}
