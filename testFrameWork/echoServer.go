package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type respStruct struct {
	Method string
	RawURL string
	Header *http.Header
	Body   *json.RawMessage
}

type errStuct struct {
	ErrMsg string
	respStruct
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if len(body) == 0 {
		body = []byte{'{', '}'}
	}
	if err != nil {
		errMsg, _ := json.MarshalIndent(&errStuct{
			ErrMsg: fmt.Sprint(err),
			respStruct: respStruct{
				Method: r.Method,
				RawURL: r.URL.String(),
				Header: &r.Header,
				Body:   (*json.RawMessage)(&body),
			},
		}, "", "\t")

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(errMsg))
		return
	}

	w.WriteHeader(http.StatusOK)
	okResp, err := json.MarshalIndent(&respStruct{
		Method: r.Method,
		RawURL: r.URL.String(),
		Header: &r.Header,
		Body:   (*json.RawMessage)(&body),
	}, "", "\t")
	if err != nil {
		fmt.Println("err is ", err)
		return
	}
	io.WriteString(w, string(okResp))
}

func startEchoServer() {
	err := http.ListenAndServe(":8000", http.HandlerFunc(echoHandler))
	if err != nil {
		panic(err)
	}
}
