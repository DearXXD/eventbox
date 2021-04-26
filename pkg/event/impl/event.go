package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/grpc/gcontext"

	"github.com/infraboard/eventbox/pkg"
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
	in, err := gcontext.GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}
	tk := pkg.S().GetToken(in.GetRequestID())

	r := newQueryEventRequest(req)
	resp, err := s.col.Find(context.TODO(), r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find event error, error is %s", err)
	}

	set := event.NewOperateEventSet()
	// 循环
	for resp.Next(context.TODO()) {
		d := event.NewOperateEvent()
		if err := resp.Decode(d); err != nil {
			return nil, exception.NewInternalServerError("decode event error, error is %s", err)
		}

		set.Add(d)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get event count error, error is %s", err)
	}
	set.Total = count

	s.log.Debug(tk)
	return event.NewOperateEventSet(), nil
}
