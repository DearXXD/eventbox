// Code generated by protoc-gen-go-http. DO NOT EDIT.

package example

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				GrpcPath:         "/eventbox.example.Service/CreateBook",
				FunctionName:     "CreateBook",
				Path:             "/books/",
				Method:           "POST",
				Resource:         "book",
				AuthEnable:       true,
				PermissionEnable: true,
				AuditLog:         false,
				Labels:           map[string]string{"action": "create"},
			},
			{
				GrpcPath:         "/eventbox.example.Service/QueryBook",
				FunctionName:     "QueryBook",
				Path:             "/books/",
				Method:           "GET",
				Resource:         "book",
				AuthEnable:       false,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{},
			},
		},
	}
	return set
}
