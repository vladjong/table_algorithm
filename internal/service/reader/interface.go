package reader

import "github.com/vladjong/test_yadro/internal/entity"

type Reader interface {
	ReadCsv() (entity.Table, error)
}
