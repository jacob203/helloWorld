package shortUrl

import (
	"MyProjects/helloWorld/shortUrlService/api/middleware"
	"context"
	"fmt"
	"net/http"
)

type shortUrlRequest struct {
	ShortUrl string
}

type longUrlResponse struct {
	LongUrl string
}

var QueryShortUrlMeta = middleware.Meta{
	Service: "short url",
	Name:    "query short url",
	Path:    "/v1/shortUrl",
	Methods: []string{http.MethodGet},
	RequestDtoFunc: func() interface{} {
		return &shortUrlRequest{}
	},
	HandlerFunc: queryShortUrlHandler,
	Middleware:  []middleware.Handler{},
}

func queryShortUrlHandler(requestDto interface{}, _ context.Context, pReq *middleware.Request) *middleware.Response {
	pReqDto := requestDto.(*shortUrlRequest)
	if len(pReqDto.ShortUrl) == 0 {
		return &middleware.Response{
			Status: http.StatusInternalServerError,
			ResponseDto: &middleware.MiddleWareErrMsg{
				Req:    pReq,
				ErrMsg: "query short url is empty",
			},
		}
	}

	fmt.Println("shortUrl is ", pReqDto.ShortUrl)
	id := shortUrlToId(pReqDto.ShortUrl)

	var longUrl string
	err := selectLongUrlByShortUrlId.QueryRow(id).Scan(&longUrl)
	if err != nil {
		return &middleware.Response{
			Status: http.StatusInternalServerError,
			ResponseDto: &middleware.MiddleWareErrMsg{
				Req:    pReq,
				ErrMsg: fmt.Sprint("query id is ", id, ", err is ", err),
			},
		}
	}

	resp := &middleware.Response{
		Status: http.StatusOK,
		ResponseDto: &longUrlResponse{
			LongUrl: longUrl,
		},
	}

	return resp
}
