package service

import "context"

// XosoService describes the service.
type XosoService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicXosoService struct{}

func (b *basicXosoService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicXosoService returns a naive, stateless implementation of XosoService.
func NewBasicXosoService() XosoService {
	return &basicXosoService{}
}

// New returns a XosoService with all of the expected middleware wired in.
func New(middleware []Middleware) XosoService {
	var svc XosoService = NewBasicXosoService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
