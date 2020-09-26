package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	http1 "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/thuc201995/xoso/pkg/config"
	endpoint "github.com/thuc201995/xoso/pkg/endpoint"
)

// makeGetByDateHandler creates the handler logic
func makeGetByDateHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Methods("GET").Path("/xsmn/gets/date/{date}").Handler(http1.NewServer(endpoints.GetByDateEndpoint, decodeGetByDateRequest, encodeGetByDateResponse, options...))
}

// decodeGetByDateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByDateRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	date := mux.Vars(r)["date"]

	if date == "" {
		err = config.NoBodyData
	}
	req = endpoint.GetByDateRequest{Date: date}

	return req, err
}

// encodeGetByDateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByDateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	switch err {
	case config.DateTimeError, config.NoBodyData:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeGetByProvinceHandler creates the handler logic
func makeGetByProvinceHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Methods("GET").Path("/get/province/{province}/date/{date}").Handler(http1.NewServer(endpoints.GetByProvinceEndpoint, decodeGetByProvinceRequest, encodeGetByProvinceResponse, options...))
}

// decodeGetByProvinceRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetByProvinceRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	vars := mux.Vars(r)
	date := vars["date"]
	province := vars["province"]
	if date == "" || province == "" {
		err = config.NoBodyData
	}
	req = endpoint.GetByProvinceRequest{Date: date, Province: province}

	return req, err
}

// encodeGetByProvinceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetByProvinceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetXSMTByDateHandler creates the handler logic
func makeGetXSMTByDateHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Methods("GET").Path("/xsmt/gets/date/{date}").Handler(http1.NewServer(endpoints.GetXSMTByDateEndpoint, decodeGetXSMTByDateRequest, encodeGetXSMTByDateResponse, options...))
}

// decodeGetXSMTByDateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetXSMTByDateRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	date := mux.Vars(r)["date"]
	if date == "" {
		err = config.NoBodyData
	}
	req = endpoint.GetXSMTByDateRequest{Date: date}

	return req, err
}

// encodeGetXSMTByDateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetXSMTByDateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetMTByProvinceHandler creates the handler logic
func makeGetMTByProvinceHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Methods("GET").Path("/xsmt/get/province/{province}/date/{date}").Handler(http1.NewServer(endpoints.GetMTByProvinceEndpoint, decodeGetMTByProvinceRequest, encodeGetMTByProvinceResponse, options...))
}

// decodeGetMTByProvinceRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetMTByProvinceRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	vars := mux.Vars(r)
	date := vars["date"]
	province := vars["province"]
	if date == "" || province == "" {
		err = config.NoBodyData
	}

	req = endpoint.GetMTByProvinceRequest{Date: date, Province: province}
	return req, err
}

// encodeGetMTByProvinceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetMTByProvinceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
