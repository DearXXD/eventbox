package engine

import (
	"context"

	"github.com/infraboard/mcube/bus"
	"github.com/infraboard/mcube/bus/event"
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
	_, err := e.s.SaveEvent(context.Background(), req)
	if err != nil {
		e.addToRetryBuf(et)
	}
	return nil
}

func (e *Engine) addToRetryBuf(et *event.Event) {
	e.l.Debugf("add event: %s to retry buffer", et.GetID())
	e.retryBuf = append(e.retryBuf, et)
}
