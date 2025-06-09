package shorter

import "ShorterAPI/internal/domain/shorter/vo"

type Repository interface {
	New(vo vo.AliasVO) error
	FindLongUrlByKey(shortUrl string) (string, error)
}
