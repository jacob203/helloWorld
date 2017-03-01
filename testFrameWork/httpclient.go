package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/myteksi/go/gothena/commons/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	client  *http.Client
	compare = &compareStruct{defaultRT: http.DefaultTransport}
)

func init() {
	// using suggested values from
	// https://httpd.apache.org/docs/2.4/mod/core.html#keepalivetimeout
	client = &http.Client{
		Transport: compare,
		Timeout:   10 * time.Second,
	}
}

type compareStruct struct {
	defaultRT http.RoundTripper
	IsSame    bool
}

func (c *compareStruct) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := c.defaultRT.RoundTrip(req)

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	if len(reqBody) == 0 {
		reqBody = []byte{'{', '}'}
	}
	reqBytes, err := json.MarshalIndent(&respStruct{
		Method: req.Method,
		RawURL: req.URL.RequestURI(),
		Header: &req.Header,
		Body:   (*json.RawMessage)(&reqBody),
	}, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("req:", (string)(reqBytes))
	req.Body = ioutil.NopCloser(bytes.NewReader(reqBody))

	var respBytes []byte
	if resp.StatusCode != http.StatusOK {
		respBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		printBytes, err := json.MarshalIndent(&struct {
			Status          string
			respStructBytes *json.RawMessage
		}{
			Status:          resp.Status,
			respStructBytes: (*json.RawMessage)(&respBytes),
		}, "", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Println((string)(printBytes))
		c.IsSame = false
	}

	respBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("resp:", (string)(respBytes))

	c.IsSame = (string)(reqBytes) == (string)(respBytes)
	return resp, err
}

type Args struct {
	Params interface{} //used for url query string
	Data   interface{} //request body
	Header http.Header
}

func newHttpRequest(method, path string, args *Args) *http.Request {
	if args == nil {
		panic(errors.New("args can't be nil"))
	}

	requestUrl, err := url.Parse("http://localhost:8000")
	if err != nil {
		panic(errors.New("err to parse url" + err.Error()))
	}
	requestUrl.Path = path
	if args.Params != nil {
		urlParam, err := query.Values(args.Params)
		if err != nil {
			panic(errors.New(fmt.Sprintf("query url encoding failed, %+v", args.Params)))
		}
		requestUrl.RawQuery = urlParam.Encode()
	}

	var bodyBytes []byte
	if args.Data != nil {
		bodyBytes, err = json.Marshal(args.Data)
		if err != nil {
			panic(err)
		}
	}
	body := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest(method, requestUrl.String(), body)
	if args.Header != nil {
		for k, v := range args.Header {
			for _, value := range v {
				req.Header.Add(k, value)
			}
		}
	}

	return req
}

func CompareRequest(method, path string, args *Args) bool {
	req := newHttpRequest(method, path, args)
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return compare.IsSame
}
