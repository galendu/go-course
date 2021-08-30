package host

import (
	"context"
	"net/http"
	"strconv"
)

type Service interface {
	SaveHost(context.Context, *Host) (*Host, error)
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
}

const (
	// DefaultPageSize 默认分页大小
	DefaultPageSize = 20
	// DefaultPageNumber 默认页号
	DefaultPageNumber = 1
)

func NewQueryHostRequestFromHTTP(r *http.Request) *QueryHostRequest {
	qs := r.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)

	if psUint64 == 0 {
		psUint64 = DefaultPageSize
	}
	if pnUint64 == 0 {
		pnUint64 = DefaultPageNumber
	}
	return &QueryHostRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
	}
}

type QueryHostRequest struct {
	PageSize   uint64 `json:"page_size,omitempty"`
	PageNumber uint64 `json:"page_number,omitempty"`
}

func (q *QueryHostRequest) OffSet() int64 {
	return int64((q.PageNumber - 1) * q.PageSize)
}
