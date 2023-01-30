package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/vladjong/test_yadro/internal/service/dfs"
	"github.com/vladjong/test_yadro/internal/service/parser"
	"github.com/vladjong/test_yadro/internal/service/reader"
	"github.com/vladjong/test_yadro/internal/service/view"
)

func main() {
	filename, err := getFilename()
	if err != nil {
		log.Fatal(err)
	}
	reader := reader.New(filename)
	table, err := reader.ReadCsv()
	if err != nil {
		log.Fatal(err)
	}

	dfs := dfs.New(table.Data)
	path, err := dfs.GetPath()
	if err != nil {
		log.Fatal(err)
	}

	parser := parser.New(table.Data, path)
	tableMap, err := parser.GetTable()
	if err != nil {
		log.Fatal(err)
	}

	table.Data = tableMap

	view := view.New(table)

	view.PrintTable()
}

func getFilename() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("[getFilename]:empty filename")
	}
	filename := os.Args[1]
	ext := path.Ext(filename)
	if ext != ".csv" {
		return "", fmt.Errorf("[getFilename]:incorrect extension file:%v, need:.csv", ext)
	}
	return filename, nil
}
