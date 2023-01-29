package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vladjong/test_yadro/internal/entity"
)

type parser struct {
	data  map[int]entity.RowString
	path  []string
	cells map[string]int
}

func New(data map[int]entity.RowString, path []string) *parser {
	return &parser{
		data:  data,
		path:  path,
		cells: make(map[string]int, len(path)),
	}
}

func (p *parser) GetTable() (map[int]entity.RowString, error) {
	out := make(map[int]entity.RowString, len(p.path))
	for _, v := range p.path {
		pattern, number, err := p.getRowAndColl(v)
		if err != nil {
			return nil, fmt.Errorf("[Parser.GetTable]:%v", err)
		}
		if pattern == "A" {
			num, _ := p.getCell(p.data[number].A)
			p.cells[v] = num
		} else if pattern == "B" {
			num, _ := p.getCell(p.data[number].B)
			p.cells[v] = num
		} else if pattern == "C" {
			num, _ := p.getCell(p.data[number].Cell)
			p.cells[v] = num
		}
	}
	for k := range p.data {
		model := entity.RowString{
			A:    fmt.Sprintf("%v", p.cells[fmt.Sprintf("A%d", k)]),
			B:    fmt.Sprintf("%v", p.cells[fmt.Sprintf("B%d", k)]),
			Cell: fmt.Sprintf("%v", p.cells[fmt.Sprintf("Cell%d", k)]),
		}
		out[k] = model
	}
	return out, nil
}

func (p *parser) getCell(in string) (int, error) {
	var slice []string
	number, err := strconv.Atoi(in)
	if err == nil {
		return number, nil
	}
	if string(in[0]) != "=" {
		return 0, fmt.Errorf("[Parser.getCell]:%v", err)
	}
	for i := 0; i < len(in); i++ {
		if string(in[i]) == "=" {
			continue
		} else if string(in[i]) == "+" || string(in[i]) == "-" || string(in[i]) == "*" || string(in[i]) == "/" {
			slice = append(slice, string(in[i]))
		} else if string(in[i]) == "A" || string(in[i]) == "B" || string(in[i]) == "C" {
			j := i
			for ; j < len(in); j++ {
				if string(in[j]) == "+" || string(in[j]) == "-" || string(in[j]) == "*" || string(in[j]) == "/" {
					slice = append(slice, fmt.Sprintf("%v", p.cells[string(in[i:j])]))
					i = j - 1
					break
				} else if j == len(in)-1 {
					slice = append(slice, fmt.Sprintf("%v", p.cells[string(in[i:j+1])]))
					i = j
					break
				}
			}
		} else {
			return 0, fmt.Errorf("[Parser.getCell]: invalid character: %v", string(in[i]))
		}
	}
	return calculate(strings.Join(slice, "")), nil
}

func (p *parser) getRowAndColl(in string) (string, int, error) {
	pattern := string(in[0])
	len := len(in)
	number := 0
	if pattern == "A" || pattern == "B" {
		val, err := strconv.Atoi(in[1:len])
		if err != nil {
			return "", 0, fmt.Errorf("[Parser.getRowAndColl]:%v", err)
		}
		number = val
	} else if pattern == "C" {
		val, err := strconv.Atoi(in[4:len])
		if err != nil {
			return "", 0, fmt.Errorf("[Parser.getRowAndColl]:%v", err)
		}
		number = val
	} else {
		return "", 0, fmt.Errorf("[Parser.getRowAndColl]: invalid character: %v", in)
	}
	return pattern, number, nil
}
