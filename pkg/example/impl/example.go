package impl

import (
	"context"

	"github.com/infraboard/mcube/grpc/gcontext"

	"github.com/infraboard/eventbox/pkg"
	"github.com/infraboard/eventbox/pkg/example"
)

func (s *service) CreateBook(ctx context.Context, req *example.CreateBookRequest) (*example.Book, error) {
	in, err := gcontext.GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}
	tk := pkg.S().GetToken(in.GetRequestID())
	s.log.Debug(tk)
	return example.NewBook(req), nil
}

func (s *service) QueryBook(ctx context.Context, req *example.QueryBookRequest) (*example.BookSet, error) {
	return example.NewBookSet(), nil
}
