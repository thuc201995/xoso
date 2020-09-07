package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/thuc201995/xoso/pkg/service"
)

// GetsRequest collects the request parameters for the Gets method.
type GetsRequest struct {
	Date string `json:"date"`
}

// GetsResponse collects the response parameters for the Gets method.
type GetsResponse struct {
	Rs  service.XSMN `json:"rs"`
	Err error        `json:"err"`
}

// MakeGetsEndpoint returns an endpoint that invokes Gets on the service.
func MakeGetsEndpoint(s service.XosoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetsRequest)
		rs, err := s.Gets(ctx, req.Date)
		return GetsResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetsResponse) Failed() error {
	return r.Err
}

// BarRequest collects the request parameters for the Bar method.
type BarRequest struct {
	S string `json:"s"`
}

// BarResponse collects the response parameters for the Bar method.
type BarResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeBarEndpoint returns an endpoint that invokes Bar on the service.
func MakeBarEndpoint(s service.XosoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BarRequest)
		rs, err := s.Bar(ctx, req.S)
		return BarResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r BarResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Gets implements Service. Primarily useful in a client.
func (e Endpoints) Gets(ctx context.Context, date string) (rs service.XSMN, err error) {
	request := GetsRequest{Date: date}
	response, err := e.GetsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetsResponse).Rs, response.(GetsResponse).Err
}

// Bar implements Service. Primarily useful in a client.
func (e Endpoints) Bar(ctx context.Context, s string) (rs string, err error) {
	request := BarRequest{S: s}
	response, err := e.BarEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(BarResponse).Rs, response.(BarResponse).Err
}
