package main

import ()
import (
	"math"
	"reflect"
)

type testStruct struct {
	method      string
	path        string
	requestType reflect.Type
}

var testStructlist = []testStruct{
	{
		method:      "Get",
		path:        "/grabpay/credit/topup/methods/",
		requestType: reflect.TypeOf(GetCreditTopupMethodsRequest{}),
	},
}

func buildTestRequests(requestType reflect.Type) interface{} {
	var maxTypeCandidateNumsF float64 = 0.0
	for i := 0; i < requestType.NumField(); i++ {
		fieldTypeStr := requestType.Field(i).Type.String()
		maxTypeCandidateNumsF = math.Max(maxTypeCandidateNumsF, float64(reflect.ValueOf(TestDataBuilderList[fieldTypeStr]).Len()))
	}

	testRequestCount := (int)(maxTypeCandidateNumsF)
	resSlice := reflect.MakeSlice(reflect.SliceOf(requestType), testRequestCount, testRequestCount)
	for i := 0; i < resSlice.Len(); i++ {
		elem := resSlice.Index(i)
		for j := 0; j < requestType.NumField(); j++ {
			fieldTypeStr := requestType.Field(j).Type.String()
			fieldTypeLen := reflect.ValueOf(TestDataBuilderList[fieldTypeStr]).Len()
			elem.Field(j).Set(reflect.ValueOf(TestDataBuilderList[fieldTypeStr]).Index(i % fieldTypeLen))
		}
	}

	return resSlice.Interface()
}
