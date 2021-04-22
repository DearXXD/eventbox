package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/eventbox/client"
	"github.com/infraboard/eventbox/pkg"
	"github.com/infraboard/eventbox/pkg/event"
)

var (
	api = &handler{}
)

type handler struct {
	service event.ServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("examples")

	r.BasePath("books")
	r.Handle("POST", "/", h.CreateBook)
	r.Handle("GET", "/", h.QueryBook)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.Event()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("event", api)
}