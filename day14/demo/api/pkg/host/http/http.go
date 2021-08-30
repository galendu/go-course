package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"gitee.com/infraboard/go-course/day14/demo/api/pkg"
	"gitee.com/infraboard/go-course/day14/demo/api/pkg/host"
)

var (
	api = &handler{}
)

type handler struct {
	service host.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Host")
	h.service = pkg.Host
	return nil
}

func (h *handler) QueryUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := host.NewQueryHostRequest()
	set, err := h.service.QueryHost(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func RegistAPI(r *httprouter.Router) {
	r.GET("/users", api.QueryUser)
	r.POST("/users", api.CreateUser)
}
