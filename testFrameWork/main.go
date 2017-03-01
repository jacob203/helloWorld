package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	go startEchoServer()
	for _, testStruct := range testStructlist {
		fmt.Println(testStruct.path)
		requestsValue := reflect.ValueOf(buildTestRequests(testStruct.requestType))
		for i := 0; i < requestsValue.Len(); i++ {
			requestByte, _ := json.MarshalIndent(requestsValue.Index(i).Interface(), "", "\t")
			fmt.Println(i, ":", string(requestByte))
			bRes := CompareRequest(testStruct.method, testStruct.path, &Args{Params: requestsValue.Index(i).Interface()})
			fmt.Println(i, ":", bRes)
			fmt.Println("=============================================")
		}
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("test end")
}
