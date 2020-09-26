package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/thuc201995/xoso/pkg/service"
)

// GetByDateRequest collects the request parameters for the GetByDate method.
type GetByDateRequest struct {
	Date string `json:"date"`
}

// GetByDateResponse collects the response parameters for the GetByDate method.
type GetByDateResponse struct {
	Rs  []service.XSMN `json:"rs"`
	Err error          `json:"err"`
}

// MakeGetByDateEndpoint returns an endpoint that invokes GetByDate on the service.
func MakeGetByDateEndpoint(s service.XosoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByDateRequest)
		rs, err := s.GetByDate(ctx, req.Date)

		return GetByDateResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByDateResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetByDate implements Service. Primarily useful in a client.
func (e Endpoints) GetByDate(ctx context.Context, date string) (rs []service.XSMN, err error) {
	request := GetByDateRequest{Date: date}
	response, err := e.GetByDateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByDateResponse).Rs, response.(GetByDateResponse).Err
}

// GetByProvinceRequest collects the request parameters for the GetByProvince method.
type GetByProvinceRequest struct {
	Date     string `json:"date"`
	Province string `json:"province"`
}

// GetByProvinceResponse collects the response parameters for the GetByProvince method.
type GetByProvinceResponse struct {
	Rs  service.XSMN `json:"rs"`
	Err error        `json:"err"`
}

// MakeGetByProvinceEndpoint returns an endpoint that invokes GetByProvince on the service.
func MakeGetByProvinceEndpoint(s service.XosoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByProvinceRequest)
		rs, err := s.GetByProvince(ctx, req.Date, req.Province)
		return GetByProvinceResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByProvinceResponse) Failed() error {
	return r.Err
}

// GetByProvince implements Service. Primarily useful in a client.
func (e Endpoints) GetByProvince(ctx context.Context, date string, province string) (rs service.XSMN, err error) {
	request := GetByProvinceRequest{
		Date:     date,
		Province: province,
	}
	response, err := e.GetByProvinceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByProvinceResponse).Rs, response.(GetByProvinceResponse).Err
}

// GetXSMTByDateRequest collects the request parameters for the GetXSMTByDate method.
type GetXSMTByDateRequest struct {
	Date string `json:"date"`
}

// GetXSMTByDateResponse collects the response parameters for the GetXSMTByDate method.
type GetXSMTByDateResponse struct {
	Rs  []service.XSMT `json:"rs"`
	Err error          `json:"err"`
}

// MakeGetXSMTByDateEndpoint returns an endpoint that invokes GetXSMTByDate on the service.
func MakeGetXSMTByDateEndpoint(s service.XosoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetXSMTByDateRequest)
		rs, err := s.GetXSMTByDate(ctx, req.Date)
		return GetXSMTByDateResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetXSMTByDateResponse) Failed() error {
	return r.Err
}

// GetXSMTByDate implements Service. Primarily useful in a client.
func (e Endpoints) GetXSMTByDate(ctx context.Context, date string) (rs []service.XSMT, err error) {
	request := GetXSMTByDateRequest{Date: date}
	response, err := e.GetXSMTByDateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetXSMTByDateResponse).Rs, response.(GetXSMTByDateResponse).Err
}

// GetMTByProvinceRequest collects the request parameters for the GetMTByProvince method.
type GetMTByProvinceRequest struct {
	Date     string `json:"date"`
	Province string `json:"province"`
}

// GetMTByProvinceResponse collects the response parameters for the GetMTByProvince method.
type GetMTByProvinceResponse struct {
	Rs  service.XSMT `json:"rs"`
	Err error        `json:"err"`
}

// MakeGetMTByProvinceEndpoint returns an endpoint that invokes GetMTByProvince on the service.
func MakeGetMTByProvinceEndpoint(s service.XosoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetMTByProvinceRequest)
		rs, err := s.GetMTByProvince(ctx, req.Date, req.Province)
		return GetMTByProvinceResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetMTByProvinceResponse) Failed() error {
	return r.Err
}

// GetMTByProvince implements Service. Primarily useful in a client.
func (e Endpoints) GetMTByProvince(ctx context.Context, date string, province string) (rs service.XSMT, err error) {
	request := GetMTByProvinceRequest{
		Date:     date,
		Province: province,
	}
	response, err := e.GetMTByProvinceEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetMTByProvinceResponse).Rs, response.(GetMTByProvinceResponse).Err
}
