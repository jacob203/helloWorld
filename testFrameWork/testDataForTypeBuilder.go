package main

import (
	"math"
	"math/rand"
	"reflect"
	"time"
)

var (
	randomCreater = rand.New(rand.NewSource(time.Now().UnixNano()))

	testString     = "helloworld"
	testDataBool   = []bool{false, true}
	testDataInt    = []int{math.MinInt32, 0, randomCreater.Int(), math.MaxInt32}
	testDataInt64  = []int64{math.MinInt32, 0, randomCreater.Int63(), math.MaxInt32}
	testDataString = []string{"", "abc"}

	TestDataBuilderList = map[string]interface{}{
		reflect.Int.String():          testDataInt,
		"*" + reflect.Int.String():    pointerOf(testDataInt),
		reflect.Int64.String():        testDataInt64,
		"*" + reflect.Int64.String():  pointerOf(testDataInt64),
		reflect.String.String():       testDataString,
		"*" + reflect.String.String(): pointerOf(testDataString),
		reflect.Bool.String():         testDataBool,
		"*" + reflect.Bool.String():   pointerOf(testDataBool),
	}
)

func pointerOf(srcSlice interface{}) interface{} {
	srcSliceValue := reflect.ValueOf(srcSlice)
	srcSliceLen := srcSliceValue.Len()
	//the last one is nil
	dstSliceValue := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(srcSliceValue.Index(0).Type())), srcSliceLen+1, srcSliceLen+1)
	for i := 0; i < srcSliceLen; i++ {
		dstSliceValue.Index(i).Set(srcSliceValue.Index(i).Addr())
	}

	return dstSliceValue.Interface()
}
