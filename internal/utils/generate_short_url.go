package utils

import (
	"math/rand"
	"time"
)

const shortURLLength = 8
const shortUrlItem = "1234567890abcdefghiklmnopqrstuvwxyzABCDEFGHIKLMNOPQRSTUVWXYZ"

var (
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GenerateShortUrl() string {
	shortUrl := make([]byte, shortURLLength)

	for v := range shortUrl {
		index := rnd.Intn(len(shortUrlItem))
		shortUrl[v] = shortUrlItem[index]
	}
	return string(shortUrl)
}
