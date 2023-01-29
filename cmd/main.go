package main

import (
	"fmt"
	"log"

	"github.com/vladjong/test_yadro/internal/service/dfs"
	"github.com/vladjong/test_yadro/internal/service/reader"
)

const (
	FILENAME = "data.csv"
)

func main() {
	reader := reader.New(FILENAME)
	data, err := reader.ReadCsv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	dfs := dfs.New(data)
	path := dfs.GetPath()
	fmt.Println(path)
}
