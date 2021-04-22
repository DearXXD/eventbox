package event

import "github.com/infraboard/mcube/bus/event"

// NewSaveEventRequest todo
func NewSaveEventRequest() *SaveEventRequest {
	return &SaveEventRequest{}
}

// NewSaveReponse todo
func NewSaveReponse() *SaveReponse {
	return &SaveReponse{}
}

// NewOperateEventSet 实例
func NewOperateEventSet() *OperateEventSet {
	return &OperateEventSet{
		Items: []*event.OperateEvent{},
	}
}
