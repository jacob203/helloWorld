package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/myteksi/go/gothena/commons/errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

var (
	client               *http.Client
	ignoredHeaderKeyList = []string{"Accept-Encoding", "User-Agent"}
)

func init() {
	// using suggested values from
	// https://httpd.apache.org/docs/2.4/mod/core.html#keepalivetimeout
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 50,
			Dial: (&net.Dialer{
				Timeout:   2 * time.Second,
				KeepAlive: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second},
		Timeout: 10 * time.Second,
	}
}

type Args struct {
	Params interface{} //used for url query string
	Data   interface{} //request body
	Header http.Header
}

func MakeHttpCall(method, path string, args *Args) (*http.Response, error) {
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

	fmt.Println("requestUrl:", requestUrl.String())
	req, err := http.NewRequest(method, requestUrl.String(), body)
	if args.Header != nil {
		for k, v := range args.Header {
			for _, value := range v {
				req.Header.Add(k, value)
			}
		}
	}

	return client.Do(req)
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
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	reqBytes, err := Marshal(&respStruct{
		Method: req.Method,
		RawURL: req.URL.RequestURI(),
		Header: &req.Header,
		Body:   string(reqBody),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("req:", (string)(reqBytes))
	req.Body = ioutil.NopCloser(bytes.NewReader(reqBody))

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var respBodyBytes []byte
	if resp.StatusCode != http.StatusOK {
		respBodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		printBytes, err := Marshal(&struct {
			Status          string
			respStructBytes *json.RawMessage
		}{
			Status:          resp.Status,
			respStructBytes: (*json.RawMessage)(&respBodyBytes),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println((string)(printBytes))
		return false
	}

	respBodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	respBody := &respStruct{}
	err = json.Unmarshal(respBodyBytes, respBody)
	if err != nil {
		panic(err)
	}
	if respBody.Header != nil {
		for _, deletingKey := range ignoredHeaderKeyList {
			respBody.Header.Del(deletingKey)
		}
	}
	newRespBodyBytes, err := Marshal(respBody)
	if err != nil {
		panic(err)
	}

	fmt.Println("resp:", (string)(newRespBodyBytes))

	return (string)(reqBytes) == (string)(newRespBodyBytes)
}
