package engine

import (
	"fmt"

	"github.com/infraboard/mcube/bus"
	"github.com/infraboard/mcube/bus/event"
	"github.com/infraboard/mcube/grpc/gcontext"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	store "github.com/infraboard/eventbox/pkg/event"
)

var (
	DefaultRetryBufSize = 64
)

func NewEngine(b bus.Subscriber, s store.ServiceServer) *Engine {
	return &Engine{
		b:            b,
		s:            s,
		retryBuf:     make([]*event.Event, 0, DefaultRetryBufSize),
		retryBufSize: DefaultRetryBufSize,
		l:            zap.L().Named("Event Engine"),
	}
}

type Engine struct {
	b            bus.Subscriber
	s            store.ServiceServer
	retryBuf     []*event.Event
	retryBufSize int
	l            logger.Logger
}

func (e *Engine) SetRetryBufSize(s int) {
	e.retryBufSize = s
}

func (e *Engine) Start() error {
	if e.b == nil {
		return fmt.Errorf("global publisher not set")
	}

	e.l.Info("start engine ...")
	subT := event.Type_Operate.String()
	if err := e.b.Sub(subT, e.Hanle); err != nil {
		return err
	}
	e.l.Infof("ok! start sub topic: %s", subT)
	return nil
}

func (e *Engine) Hanle(topic string, et *event.Event) error {
	req := store.NewSaveEventRequest()
	req.Add(et)
	in := gcontext.NewGrpcInCtx()
	_, err := e.s.SaveEvent(in.Context(), req)
	if err != nil {
		e.l.Errorf("save event error, %s", err)
		e.addToRetryBuf(et)
	}
	return nil
}

func (e *Engine) addToRetryBuf(et *event.Event) {
	e.l.Debugf("add event: %s to retry buffer", et.Id)
	e.retryBuf = append(e.retryBuf, et)
}
