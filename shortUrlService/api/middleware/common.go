package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

const logTag = "common"

// HandlerFunc is the type of an endpoint handler.
type HandlerFunc func(interface{}, context.Context, *Request) *Response

// Handler is the type of an endpoint middleware.
type Handler func(Meta, HandlerFunc) HandlerFunc
type Request struct {
	*http.Request
	Body []byte
}

func (pReq *Request) MarshalJSON() ([]byte, error) {
	bodyBytes := pReq.Body
	if len(bodyBytes) == 0 {
		bodyBytes = []byte{'{', '}'}
	}
	body, err := json.Marshal(&struct {
		RequestURI string
		Header     http.Header
		Body       json.RawMessage
	}{
		RequestURI: pReq.RequestURI,
		Header:     pReq.Header,
		Body:       bodyBytes,
	})

	return body, err
}

type Response struct {
	Status      int
	ResponseDto interface{}
	Header      http.Header
}

type MiddleWareErrMsg struct {
	Req    *Request
	ErrMsg string `json:"ErrMsg", required:"true"`
}

func (pErr *MiddleWareErrMsg) MarshalJSON() ([]byte, error) {
	bodyBytes := pErr.Req.Body
	if len(bodyBytes) == 0 {
		bodyBytes = []byte{'{', '}'}
	}
	return json.Marshal(&struct {
		ErrMsg     string
		RequestURI string
		Header     http.Header
		Body       json.RawMessage
	}{
		ErrMsg:     pErr.ErrMsg,
		RequestURI: pErr.Req.RequestURI,
		Header:     pErr.Req.Header,
		Body:       bodyBytes,
	})
}

// Meta defines the route information (metadata) for a given endpoint.
type Meta struct {
	Service string
	Name    string

	Path           string
	Methods        []string
	HandlerFunc    HandlerFunc
	RequestDtoFunc func() interface{}

	Middleware []Handler
}

// BindToRouter attaches the given handler definitions to the router.
func BindToRouter(r *mux.Router, handlerMetas ...Meta) {
	for _, meta := range handlerMetas {

		httpHandlerFunc := BuildHandlerFunc(meta)
		r.HandleFunc(meta.Path, httpHandlerFunc).Methods(meta.Methods...)
	}
}

func writeResponse(rw http.ResponseWriter, req *http.Request, resp *Response) {
	// Set header
	for key, values := range resp.Header {
		for _, value := range values {
			rw.Header().Add(key, value)
		}
	}

	// Flush result, note the order cannot be changed
	rw.WriteHeader(resp.Status)

	if resp.ResponseDto != nil {
		body, err := json.MarshalIndent(resp.ResponseDto, "", "\t")
		if err == nil {
			_, err = rw.Write(body)
			if err != nil {
				fmt.Println("err is ", err)
			}
		} else {
			_, err = rw.Write([]byte(err.Error()))
		}
	}
}

// BuildHandlerFunc builds a http.HandlerFunc with meta
func BuildHandlerFunc(meta Meta) http.HandlerFunc {
	handlerFunc := meta.HandlerFunc

	// Because we are wrapping the final handlerMeta
	// the first middleware should be the last to be wrapped
	middleware := append(preMiddlewareList, meta.Middleware...)
	j := len(middleware)
	for i := range middleware {
		if middleware[j-i-1] == nil {
			continue
		}
		handlerFunc = middleware[j-i-1](meta, handlerFunc)
	}

	return func(rw http.ResponseWriter, httpReq *http.Request) {
		req := &Request{Request: httpReq}
		req.Body, _ = ioutil.ReadAll(req.Request.Body)

		defer func() {
			if r := recover(); r != nil {
				errMsg := &MiddleWareErrMsg{
					Req:    req,
					ErrMsg: fmt.Sprint(r),
				}
				resp := &Response{
					Status:      http.StatusInternalServerError,
					ResponseDto: errMsg,
				}
				writeResponse(rw, httpReq, resp)
			}
		}()

		//if err != nil {
		//	rw.WriteHeader(http.StatusInternalServerError)
		//	rw.Write([]byte(fmt.Sprint("Can't read content from request, its error is ", err)))
		//	return
		//}

		mashRes, err := json.MarshalIndent(req, "", "\t")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(fmt.Sprint("an error happen when mashing request, err is ", err)))
			return
		}
		fmt.Println("request:", (string)(mashRes))

		var requestDto interface{}
		if meta.RequestDtoFunc != nil {
			requestDto = meta.RequestDtoFunc()
		}
		resp := handlerFunc(requestDto, context.Background(), req)

		writeResponse(rw, httpReq, resp)
	}
}
