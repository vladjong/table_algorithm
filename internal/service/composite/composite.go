package composite

import (
	"fmt"

	"github.com/vladjong/test_yadro/internal/service/dfs"
	"github.com/vladjong/test_yadro/internal/service/parser"
	"github.com/vladjong/test_yadro/internal/service/reader"
	"github.com/vladjong/test_yadro/internal/service/view"
)

func GetTable(filename string) error {
	reader := reader.New(filename)
	table, err := reader.ReadCsv()
	if err != nil {
		return fmt.Errorf("[Composite.GetTable]:%v", err)
	}

	dfs := dfs.New(table.Data)
	path, err := dfs.GetPath()
	if err != nil {
		return fmt.Errorf("[Composite.GetTable]:%v", err)
	}

	parser := parser.New(table.Data, path)
	tableMap, err := parser.GetTable()
	if err != nil {
		return fmt.Errorf("[Composite.GetTable]:%v", err)
	}
	table.Data = tableMap

	view := view.New(table)
	view.PrintTable()

	return nil
}
