package main

import (
	"log"

	"github.com/vladjong/test_yadro/internal/service/dfs"
	"github.com/vladjong/test_yadro/internal/service/parser"
	"github.com/vladjong/test_yadro/internal/service/reader"
	"github.com/vladjong/test_yadro/internal/service/view"
)

const (
	FILENAME = "test.csv"
)

func main() {
	reader := reader.New(FILENAME)
	table, err := reader.ReadCsv()
	if err != nil {
		log.Fatal(err)
	}

	dfs := dfs.New(table.Data)
	path := dfs.GetPath()

	parser := parser.New(table.Data, path)
	tableMap, err := parser.GetTable()
	if err != nil {
		log.Fatal(err)
	}

	table.Data = tableMap

	view := view.New(table)

	view.PrintTable()
}
