package vo

type AliasVO struct {
	LongUrl  string
	ShortUrl string
}

func NewUrlAliasVO(longUrl string, shortUrl string) AliasVO {
	return AliasVO{
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	}
}
