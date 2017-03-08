package server

import (
	"MyProjects/helloWorld/shortUrlService/api"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const logTag = "shortUrl"

func Start() {
	r := initRouter()
	if err := http.ListenAndServe(":8090", r); err != nil {
		fmt.Println("http server fails to start, err is ", err)
	}
}

func initRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	api.BuildRouter(r)

	return r
}

//func initServer(r *mux.Router) *graceful.Server {
//	// Use graceful to do proper clean up before server exit
//	return &graceful.Server{
//		Server: &http.Server{
//			Addr:           "0.0.0.0:8088",
//			Handler:        r,
//			ReadTimeout:    10 * time.Second,
//			WriteTimeout:   10 * time.Second,
//			MaxHeaderBytes: 1 << 20,
//		},
//		Timeout: time.Second * 10,
//	}
//}
