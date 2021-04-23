package event

import "github.com/infraboard/mcube/bus/event"

// NewSaveEventRequest todo
func NewSaveEventRequest() *SaveEventRequest {
	return &SaveEventRequest{}
}

func (req *SaveEventRequest) Add(item *event.Event) {
	req.Items = append(req.Items, item)
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
