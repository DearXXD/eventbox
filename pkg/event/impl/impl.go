package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"

	"github.com/infraboard/eventbox/pkg"
	"github.com/infraboard/eventbox/pkg/event"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	event.UnimplementedServiceServer

	log logger.Logger
}

func (s *service) Config() error {
	// get global config with here
	s.log = zap.L().Named("Event")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return event.HttpEntry()
}

func init() {
	pkg.RegistryService("event", Service)
}
