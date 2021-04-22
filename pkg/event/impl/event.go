package impl

import (
	"context"

	"github.com/infraboard/mcube/grpc/gcontext"

	"github.com/infraboard/eventbox/pkg"
	"github.com/infraboard/eventbox/pkg/event"
)

func (s *service) SaveEvent(ctx context.Context, req *event.SaveEventRequest) (*event.SaveReponse, error) {
	in, err := gcontext.GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}
	tk := pkg.S().GetToken(in.GetRequestID())
	s.log.Debug(tk)
	return event.NewSaveReponse(), nil
}

func (s *service) QueryEvent(ctx context.Context, req *event.QueryEventRequest) (*event.OperateEventSet, error) {
	return event.NewOperateEventSet(), nil
}
