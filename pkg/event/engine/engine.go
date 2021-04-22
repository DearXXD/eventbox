package engine

import (
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
	return e.b.Sub(event.Type_Operate.String(), e.Hanle)
}

func (e *Engine) Hanle(topic string, et *event.Event) error {
	_, err := e.s.SaveEvent(nil, nil)
	if err != nil {
		e.addToRetryBuf(et)
	}
	return nil
}

func (e *Engine) addToRetryBuf(et *event.Event) {
	e.l.Debugf("add event: %s to retry buffer", et.GetID())
	e.retryBuf = append(e.retryBuf, et)
}
