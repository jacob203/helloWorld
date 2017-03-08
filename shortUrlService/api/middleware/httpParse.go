package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"mime"
	"net/http"
)

const maxMemorySize = 1 << 20

func HttpParseHandler(meta Meta, nextFunc HandlerFunc) HandlerFunc {
	return func(reqDto interface{}, ctx context.Context, httpReq *Request) *Response {
		err := decode(reqDto, httpReq)
		if err != nil {
			errStr := fmt.Sprint("httpParse decode error:", err)
			return &Response{
				Status: http.StatusInternalServerError,
				ResponseDto: &MiddleWareErrMsg{
					Req:    httpReq,
					ErrMsg: errStr,
				},
			}
		}

		return nextFunc(reqDto, ctx, httpReq)
	}

}

func decode(requestDto interface{}, req *Request) error {
	var err error

	if requestDto == nil {
		return nil
	}

	ct := req.Header.Get("Content-Type")
	if ct != "" {
		ct, _, err = mime.ParseMediaType(ct)
		if err != nil {
			return err
		}
	}

	switch ct {
	case "application/json":
		// WARNING: never not remove this zero-length check.
		// some clients may give an empty body which is invalid json but must
		// still be processed.
		if req.ContentLength > 0 {
			if err = json.Unmarshal(req.Body, requestDto); err != nil {
				fmt.Println("fail to parse requestDto", err)
				return err
			}
		}
		// also call ParseForm so that url query params are parsed for us.
		// the body will be ignored by current implementation which checks
		// the Content-Type header.
		if err = req.ParseForm(); err != nil {
			return err
		}
	case "multipart/form-data":
		if err = req.ParseMultipartForm(maxMemorySize); err != nil {
			return err
		}
	default: // ParseForm handles all other common content types
		if err = req.ParseForm(); err != nil {
			return err
		}
	}

	// put mux variables into Form, so that they can be parsed together by schema
	vars := mux.Vars(req.Request)
	if vars != nil {
		for key, value := range vars {
			// use Set() here to overwrite any values in the request with values from the map
			req.Form.Set(key, value)
		}
	}

	decoder := schema.NewDecoder()
	// note that this causes url query params for application/json content-type
	// to have a higher precendence (which is not the case for other types due
	// to the implementation of `ParseForm` and `ParseMultipartForm`)
	err = decoder.Decode(requestDto, req.Form)
	if err != nil {
		return err
	}

	return nil
}
