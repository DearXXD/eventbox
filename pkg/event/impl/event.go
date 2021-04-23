package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"

	"github.com/infraboard/eventbox/pkg/event"
)

func (s *service) SaveEvent(ctx context.Context, req *event.SaveEventRequest) (*event.SaveReponse, error) {
	ins, err := req.ParseEvent()
	if err != nil {
		return nil, err
	}

	if _, err := s.col.InsertMany(context.TODO(), ins); err != nil {
		return nil, exception.NewInternalServerError("inserted policy(%s) document error, %s",
			req.Ids(), err)
	}

	resp := event.NewSaveReponse()
	resp.AddSuccess(req.Ids()...)
	return resp, nil
}

func (s *service) QueryEvent(ctx context.Context, req *event.QueryEventRequest) (*event.OperateEventSet, error) {
	// in, err := gcontext.GetGrpcInCtx(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// tk := pkg.S().GetToken(in.GetRequestID())
	// s.log.Debug(tk)
	return event.NewOperateEventSet(), nil
}
