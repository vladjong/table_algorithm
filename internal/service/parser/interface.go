package parser

import "github.com/vladjong/test_yadro/internal/entity"

type Parser interface {
	GetTable() (map[int]entity.RowString, error)
}
