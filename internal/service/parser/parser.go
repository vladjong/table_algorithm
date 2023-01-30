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
	p.fillCels()
	out := make(map[int]entity.RowString, len(p.path))
	for _, v := range p.path {
		pattern, number, err := p.getRowAndColl(v)
		if err != nil {
			return nil, fmt.Errorf("[Parser.GetTable]:%v", err)
		}
		if pattern == "A" {
			num, err := p.getCell(p.data[number].A, v)
			if err != nil {
				return nil, fmt.Errorf("[Parser.GetTable]:%v", err)
			}
			p.cells[v] = num
		} else if pattern == "B" {
			num, err := p.getCell(p.data[number].B, v)
			if err != nil {
				return nil, fmt.Errorf("[Parser.GetTable]:%v", err)
			}
			p.cells[v] = num
		} else if pattern == "C" {
			num, err := p.getCell(p.data[number].Cell, v)
			if err != nil {
				return nil, fmt.Errorf("[Parser.GetTable]:%v", err)
			}
			p.cells[v] = num
		} else if _, err := strconv.Atoi(pattern); err == nil {
			if _, err := p.getCell(p.data[number].Cell, v); err != nil {
				return nil, fmt.Errorf("[Parser.GetTable]:%v", err)
			}

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

func (p *parser) fillCels() {
	for _, v := range p.path {
		p.cells[v] = 0
	}
}

func (p *parser) getCell(in, path string) (int, error) {
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
		} else if string(in[i]) == "A" || string(in[i]) == "B" || string(in[i]) == "C" || isNumber(rune(in[i])) {
			j := i
			for ; j < len(in); j++ {
				if string(in[j]) == "+" || string(in[j]) == "-" || string(in[j]) == "*" || string(in[j]) == "/" {
					if path == string(in[i:j]) {
						return 0, fmt.Errorf("[Parser.getCell]:incorrect cell:%v", string(in[i:j]))
					}
					if _, err := strconv.Atoi(string(in[i:j])); err != nil {
						if _, ok := p.cells[string(in[i:j])]; !ok {
							return 0, fmt.Errorf("[Parser.getCell]:incorrect cell:%v", string(in[i:j]))
						}
						slice = append(slice, fmt.Sprintf("%v", p.cells[string(in[i:j])]))
					} else {
						slice = append(slice, string(in[i:j]))
					}
					i = j - 1
					break
				} else if j == len(in)-1 {
					if path == string(in[i:j+1]) {
						return 0, fmt.Errorf("[Parser.getCell]:incorrect cell:%v", string(in[i:j+1]))
					}
					if _, err := strconv.Atoi(string(in[i : j+1])); err != nil {
						if _, ok := p.cells[string(in[i:j+1])]; !ok {
							return 0, fmt.Errorf("[Parser.getCell]:incorrect cell:%v", string(in[i:j+1]))
						}
						slice = append(slice, fmt.Sprintf("%v", p.cells[string(in[i:j+1])]))
					} else {
						slice = append(slice, string(in[i:j+1]))
					}
					i = j
					break
				}
			}
		}
	}
	answer, err := calculate(strings.Join(slice, ""))
	if err != nil {
		return 0, fmt.Errorf("[Parser.getCell]:%v", err)
	}
	return answer, nil
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
	}
	return pattern, number, nil
}
