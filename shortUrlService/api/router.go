package api

import (
	"MyProjects/helloWorld/shortUrlService/api/middleware"
	"MyProjects/helloWorld/shortUrlService/api/shortUrl"
	"github.com/gorilla/mux"
)

func BuildRouter(r *mux.Router) {
	handlers := []middleware.Meta{
		shortUrl.CreateShortUrlMeta,
		shortUrl.QueryShortUrlMeta,
	}

	middleware.BindToRouter(r, handlers...)
}
