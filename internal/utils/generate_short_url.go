package utils

import (
	"math/rand"
	"time"
)

const shortURLMaxLength = 8
const shortURLMinLength = 4
const shortUrlItem = "1234567890abcdefghiklmnopqrstuvwxyzABCDEFGHIKLMNOPQRSTUVWXYZ"

var (
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GenerateShortUrl() string {
	shortUrl := make([]byte, rnd.Intn(shortURLMaxLength-shortURLMinLength+1)+shortURLMinLength)

	for i := range shortUrl {
		index := rnd.Intn(len(shortUrlItem))
		shortUrl[i] = shortUrlItem[index]
	}
	return string(shortUrl)
}
