package parser

import (
	"fmt"
	"strings"

	"github.com/vladjong/test_yadro/internal/constant"
	"github.com/vladjong/test_yadro/internal/entity"
)

type parser struct {
	data  map[int]entity.RowString
	graph map[string][]string
}

func New(data map[int]entity.RowString) *parser {
	p := &parser{
		data:  data,
		graph: make(map[string][]string, len(data)*constant.SIZE_COL),
	}
	return p
}

func (p *parser) GetAdjacencyGraph() map[string][]string {
	for k1 := range p.data {
		var arrA []string
		var arrB []string
		var arrCell []string
		keyA := fmt.Sprintf("A%d", k1)
		keyB := fmt.Sprintf("B%d", k1)
		keyCell := fmt.Sprintf("Cell%d", k1)
		for k2, v2 := range p.data {
			arrA = append(arrA, isExist(keyA, k2, v2)...)
			arrB = append(arrB, isExist(keyB, k2, v2)...)
			arrCell = append(arrCell, isExist(keyCell, k2, v2)...)
		}
		if len(arrA) != 0 {
			p.graph[keyA] = arrA
		}
		if len(arrB) != 0 {
			p.graph[keyB] = arrB
		}
		if len(arrCell) != 0 {
			p.graph[keyCell] = arrCell
		}
	}
	return p.graph
}

func isExist(pattern string, key int, data entity.RowString) []string {
	var arr []string
	if strings.Contains(data.A, pattern) {
		arr = append(arr, fmt.Sprintf("A%d", key))
	}
	if strings.Contains(data.B, pattern) {
		arr = append(arr, fmt.Sprintf("B%d", key))
	}
	if strings.Contains(data.Cell, pattern) {
		arr = append(arr, fmt.Sprintf("Cell%d", key))
	}
	return arr
}
