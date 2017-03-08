package shortUrl

import (
	"MyProjects/helloWorld/shortUrlService/api/middleware"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/myteksi/go/gothena/commons/errors"
	"net/http"
	"net/url"
	"strings"
)

type createShortUrlRequestDto struct {
	LongUrl string `json:"longurl, required"`
}

type createShortUrlResponseDto struct {
	ShortUrl string `json:"shorturl" required:"true"`
}

var CreateShortUrlMeta = middleware.Meta{
	Service: "short url",
	Name:    "create short url",
	Path:    "/v1/shortUrl",
	Methods: []string{http.MethodPost},
	RequestDtoFunc: func() interface{} {
		return &createShortUrlRequestDto{}
	},
	HandlerFunc: createShortUrlHandler,
	Middleware:  []middleware.Handler{},
}

func validUrl(urlStr string) (*url.URL, error) {
	longUrl, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	isSupported := false
	supportedSchemeList := []string{
		"ftp",
		"http",
		"https",
	}
	for _, v := range supportedSchemeList {
		if strings.EqualFold(longUrl.Scheme, v) {
			isSupported = true
			break
		}
	}
	if !isSupported {
		return nil, errors.New(fmt.Sprint("Scheme ", longUrl.Scheme, " only schemes ", supportedSchemeList, "is supported"))
	}

	return longUrl, nil
}

func createShortUrlHandler(requestDto interface{}, _ context.Context, pReq *middleware.Request) *middleware.Response {
	reqDto := requestDto.(*createShortUrlRequestDto)
	_, err := validUrl(reqDto.LongUrl)
	if err != nil {
		return &middleware.Response{
			Status: http.StatusInternalServerError,
			ResponseDto: &middleware.MiddleWareErrMsg{
				Req:    pReq,
				ErrMsg: fmt.Sprint(err),
			},
		}
	}

	res, err := insertLongUrl.Exec(reqDto.LongUrl)
	if err != nil {
		return &middleware.Response{
			Status: http.StatusInternalServerError,
			ResponseDto: &middleware.MiddleWareErrMsg{
				Req:    pReq,
				ErrMsg: fmt.Sprint(err),
			},
		}
	}

	shortUrlId, err := res.LastInsertId()
	if err != nil {
		return &middleware.Response{
			Status: http.StatusInternalServerError,
			ResponseDto: &middleware.MiddleWareErrMsg{
				Req:    pReq,
				ErrMsg: fmt.Sprint(err),
			},
		}
	}

	resp := &middleware.Response{
		Status: http.StatusOK,
		ResponseDto: &createShortUrlResponseDto{
			ShortUrl: IdToShortUrl(shortUrlId),
		},
	}

	return resp
}
