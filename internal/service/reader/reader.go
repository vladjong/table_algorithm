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
		return fmt.Errorf("[Reader.checkFirstLine]: incorrect fist line, should be \"\", not \"%v\"", in[0])
	} else if in[1] != "A" {
		return fmt.Errorf("[Reader.checkFirstLine]: incorrect fist line, should be \"A\", not \"%v\"", in[1])
	} else if in[2] != "B" {
		return fmt.Errorf("[Reader.checkFirstLine]: incorrect fist line, should be \"B\", not \"%v\"", in[2])
	} else if in[3] != "Cell" {
		return fmt.Errorf("[Reader.checkFirstLine]: incorrect fist line, should be \"Cell\", not \"%v\"", in[3])
	}
	return nil
}

func checkSize(in []string) error {
	if len(in) != 4 {
		return fmt.Errorf("[Reader.checkSize]: incorrect size row, should be 3, not %v", len(in))
	}
	return nil
}

func (r *reader) ReadCsv() (map[int]entity.RowString, error) {
	m := make(map[int]entity.RowString)
	f, err := os.Open(r.filename)
	if err != nil {
		return nil, fmt.Errorf("[Reader.ReadCsv]:%v", err)
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
			return nil, fmt.Errorf("[Reader.ReadCsv]:%v", err)
		}
		if checkSize(rec) != nil {
			return nil, fmt.Errorf("[Reader.ReadCsv]: %v", len(rec))
		}
		if isFirstRow {
			if err := checkFirstLine(rec); err != nil {
				return nil, fmt.Errorf("[Reader.ReadCsv]:%v", err)
			}
			isFirstRow = false
			continue
		}
		key, row, err := getKeyAndValue(rec)
		if err != nil {
			return nil, fmt.Errorf("[Reader.ReadCsv]:%v", err)
		}
		m[key] = row
	}
	return m, nil
}
