package event

import "github.com/infraboard/mcube/http/request"

// NewQueryEventkRequest 查询book列表
func NewQueryEventkRequest(page *request.PageRequest) *QueryEventRequest {
	return &QueryEventRequest{
		Page: &page.PageRequest,
	}
}
