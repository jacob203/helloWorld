package shortUrl

import (
	"errors"
	"fmt"
)

const tokenBase int64 = 10 + 26 + 26
const shortUrlLen = 8

func IdToShortUrl(shortUrlId int64) string {
	numToChar := func(id uint8) byte {
		switch {
		case id >= 0 && id <= 9:
			return '0' + id
		case id >= 10 && id <= 35:
			return 'a' + id - 10
		case id >= 36 && id <= 61:
			return 'A' + id - 36
		default:
			panic(errors.New(fmt.Sprint("its value is ", id, ", it must be in 0~61")))

		}
	}

	tokenSlice := make([]byte, shortUrlLen, shortUrlLen)
	for i := 0; i < len(tokenSlice); i++ {
		tokenSlice[i] = numToChar((uint8)(shortUrlId % tokenBase))
		shortUrlId /= tokenBase
	}

	return string(tokenSlice)
}

func shortUrlToId(shortUrl string) int64 {
	charToNum := func(c byte) uint8 {
		switch {
		case c >= '0' && c <= '9':
			return c - '0'
		case c >= 'a' && c <= 'z':
			return c - 'a' + 10
		case c >= 'A' && c <= 'Z':
			return c - 'A' + 36
		default:
			panic(errors.New(fmt.Sprint("its value is ", c, ", it must be [0~9, a~z, A~Z]")))
		}
	}

	var id int64 = 0
	for i := len(shortUrl) - 1; i >= 0; i-- {
		id = id*tokenBase + (int64)(charToNum(shortUrl[i]))
	}

	return id
}
