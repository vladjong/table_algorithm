package main

import (
	"log"

	"github.com/vladjong/test_yadro/internal/service/composite"
	"github.com/vladjong/test_yadro/internal/service/flag"
)

func main() {
	filename, err := flag.GetFilename()
	if err != nil {
		log.Fatal(err)
	}

	if err := composite.GetTable(filename); err != nil {
		log.Fatal(err)
	}
}
