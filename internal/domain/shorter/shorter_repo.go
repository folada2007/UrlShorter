package shorter

import "ShorterAPI/internal/domain/shorter/vo"

type Repository interface {
	New(vo vo.AliasVO) error
	GetById(id int) UrlMapping
	GetByName(name string) UrlMapping
}
