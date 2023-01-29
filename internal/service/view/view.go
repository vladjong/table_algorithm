package view

import (
	"fmt"

	"github.com/vladjong/test_yadro/internal/entity"
)

type view struct {
	table entity.Table
}

func New(table entity.Table) *view {
	return &view{
		table: table,
	}
}

func (v *view) PrintTable() {
	fmt.Printf(",A,B,Cell\n")
	for _, val := range v.table.Row {
		data := v.table.Data[val]
		fmt.Printf("%d,%v,%v,%v\n", val, data.A, data.B, data.Cell)
	}
}
