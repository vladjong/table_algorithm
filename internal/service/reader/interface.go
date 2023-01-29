package reader

import "github.com/vladjong/test_yadro/internal/entity"

type Reader interface {
	ReadCsv() (map[int]entity.RowString, error)
}
