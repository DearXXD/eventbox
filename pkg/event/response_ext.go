package event

import "github.com/infraboard/mcube/bus/event"

// NewSaveEventRequest todo
func NewSaveEventRequest() *SaveEventRequest {
	return &SaveEventRequest{}
}

func (req *SaveEventRequest) Add(item *event.Event) {
	req.Items = append(req.Items, item)
}

func (req *SaveEventRequest) Ids() []string {
	ids := make([]string, 0, len(req.Items))
	for i := range req.Items {
		ids = append(ids, req.Items[i].GetID())
	}

	return ids
}

func (req *SaveEventRequest) ParseEvent() ([]interface{}, error) {
	docs := make([]interface{}, 0, len(req.Items))
	for i := range req.Items {
		switch req.Items[i].Header.Type {
		case event.Type_Operate:
			data := &event.OperateEventData{}
			err := req.Items[i].ParseBoby(data)
			if err != nil {
				return nil, err
			}
			oe := &event.OperateEvent{
				Header: req.Items[i].Header,
				Body:   data,
			}
			if err != nil {
				return nil, err
			}
			docs = append(docs, oe)
		}
	}

	return docs, nil
}

// NewSaveReponse todo
func NewSaveReponse() *SaveReponse {
	return &SaveReponse{}
}

func (resp *SaveReponse) AddSuccess(ids ...string) {
	for i := range ids {
		resp.Success = append(resp.Success, ids[i])
	}
}

func (resp *SaveReponse) AddFailed(ids ...string) {
	for i := range ids {
		resp.Failed = append(resp.Failed, ids[i])
	}
}

// NewOperateEventSet 实例
func NewOperateEventSet() *OperateEventSet {
	return &OperateEventSet{
		Items: []*event.OperateEvent{},
	}
}
