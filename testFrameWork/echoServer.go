package main

import (
	"bytes"
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
	Body   string
}

type errStuct struct {
	ErrMsg string
	respStruct
}

func Marshal(v interface{}) ([]byte, error) {
	var bodyBytes bytes.Buffer
	pEncoder := json.NewEncoder(&bodyBytes)
	pEncoder.SetEscapeHTML(false)
	pEncoder.SetIndent("", "\t")
	err := pEncoder.Encode(v)
	return bodyBytes.Bytes(), err
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errMsg, _ := Marshal(&errStuct{
			ErrMsg: fmt.Sprint(err),
			respStruct: respStruct{
				Method: r.Method,
				RawURL: r.URL.RequestURI(),
				Header: &r.Header,
				Body:   string(body),
			}})
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(errMsg))
		return
	}

	w.WriteHeader(http.StatusOK)
	okResp, err := Marshal(&respStruct{
		Method: r.Method,
		RawURL: r.URL.RequestURI(),
		Header: &r.Header,
		Body:   string(body),
	})
	if err != nil {
		panic(err)
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
