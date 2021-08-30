package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"

	"gitee.com/infraboard/go-course/day14/demo/api/pkg/host"
)

func (h *handler) QueryHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := host.NewQueryHostRequestFromHTTP(r)
	set, err := h.service.QueryHost(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ins := host.NewDefaultHost()

	if err := request.GetDataFromRequest(r, ins); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.SaveHost(r.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DescribeHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := host.NewDescribeHostRequestWithID(ps.ByName("id"))
	set, err := h.service.DescribeHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := host.NewDeleteHostRequestWithID(ps.ByName("id"))
	set, err := h.service.DeleteHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
