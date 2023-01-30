package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/vladjong/test_yadro/internal/entity"
)

type reader struct {
	filename string
}

func New(filename string) *reader {
	return &reader{
		filename: filename,
	}
}

func getKeyAndValue(in []string) (int, entity.RowString, error) {
	key, err := strconv.Atoi(in[0])
	if err != nil {
		return 0, entity.RowString{}, fmt.Errorf("[Reader.getKeyAndValue]: %v", err)
	}
	return key, entity.RowString{
		A:    in[1],
		B:    in[2],
		Cell: in[3],
	}, nil
}

func checkFirstLine(in []string) error {
	if in[0] != "" {
		return fmt.Errorf("[Reader.checkFirstLine]:incorrect fist line, should be \"\", not \"%v\"", in[0])
	} else if in[1] != "A" {
		return fmt.Errorf("[Reader.checkFirstLine]:incorrect fist line, should be \"A\", not \"%v\"", in[1])
	} else if in[2] != "B" {
		return fmt.Errorf("[Reader.checkFirstLine]:incorrect fist line, should be \"B\", not \"%v\"", in[2])
	} else if in[3] != "Cell" {
		return fmt.Errorf("[Reader.checkFirstLine]:incorrect fist line, should be \"Cell\", not \"%v\"", in[3])
	}
	return nil
}

func checkSell(in []string) error {
	for _, v := range in {
		_, err := strconv.Atoi(v)
		if err == nil {
			continue
		}
		if len(in) < 2 {
			return fmt.Errorf("[Reader.checkSell]:incorrect cell: \"%v\"", v)
		}
		if string(v[0]) != "=" {
			return fmt.Errorf("[Reader.checkSell]:incorrect cell, should be \"=\", not \"%v\"", v)
		} else if string(v[1]) != "A" && string(v[1]) != "B" && string(v[1]) != "C" {
			if _, err := strconv.Atoi(string(v[1])); err == nil {
				continue
			}
			return fmt.Errorf("[Reader.checkSell]:incorrect cell, should be \"A, B, Cell or number\", not \"%v\"", v)
		}
	}
	return nil
}

func checkSize(in []string) error {
	if len(in) != 4 {
		return fmt.Errorf("[Reader.checkSize]:incorrect size row, should be 3, not %v", len(in))
	}
	return nil
}

func (r *reader) ReadCsv() (entity.Table, error) {
	m := make(map[int]entity.RowString)
	var rows []int
	f, err := os.Open(r.filename)
	if err != nil {
		return entity.Table{}, fmt.Errorf("[Reader.ReadCsv]:%v", err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	isFirstRow := true
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return entity.Table{}, fmt.Errorf("[Reader.ReadCsv]:%v", err)
		}
		if checkSize(rec) != nil {
			return entity.Table{}, fmt.Errorf("[Reader.ReadCsv]:%v", len(rec))
		}
		if isFirstRow {
			if err := checkFirstLine(rec); err != nil {
				return entity.Table{}, fmt.Errorf("[Reader.ReadCsv]:%v", err)
			}
			isFirstRow = false
			continue
		}
		if err := checkSell(rec); err != nil {
			return entity.Table{}, fmt.Errorf("[Reader.ReadCsv]:%v", err)
		}
		key, row, err := getKeyAndValue(rec)
		if err != nil {
			return entity.Table{}, fmt.Errorf("[Reader.ReadCsv]:%v", err)
		}
		rows = append(rows, key)
		m[key] = row
	}
	return entity.Table{
		Data: m,
		Row:  rows,
	}, nil
}
