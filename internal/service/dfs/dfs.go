package dfs

import (
	"fmt"
	"strings"

	"github.com/vladjong/test_yadro/internal/constant"
	"github.com/vladjong/test_yadro/internal/entity"
)

type dfs struct {
	path   []string
	graph  map[string][]string
	data   map[int]entity.RowString
	colors map[string]int
}

func New(data map[int]entity.RowString) *dfs {
	return &dfs{
		path:   make([]string, 0, len(data)*constant.SIZE_COL),
		graph:  getAdjacencyGraph(data),
		data:   data,
		colors: getColors(data),
	}
}

func (d *dfs) GetPath() ([]string, error) {
	for k, v := range d.colors {
		if v == constant.WHITE {
			if err := d.dfs(k, d.colors, d.graph); err != nil {
				return nil, fmt.Errorf("[DFS.GetPath]:%v", err)
			}
		}
	}
	reverseSlice(d.path)
	return d.path, nil
}

func (d *dfs) dfs(k string, colors map[string]int, graph map[string][]string) error {
	colors[k] = constant.GRAY
	for _, v := range graph[k] {
		if colors[v] == constant.WHITE {
			d.dfs(v, colors, graph)
		}
		if colors[v] == constant.GRAY {
			return fmt.Errorf("[DFS.dfs]:file has a cycle")
		}
	}
	colors[k] = constant.BLACK
	d.path = append(d.path, k)
	return nil
}

func getColors(data map[int]entity.RowString) map[string]int {
	colors := make(map[string]int, len(data)*constant.SIZE_COL)
	for k := range data {
		colors[fmt.Sprintf("A%d", k)] = constant.WHITE
		colors[fmt.Sprintf("B%d", k)] = constant.WHITE
		colors[fmt.Sprintf("Cell%d", k)] = constant.WHITE
	}
	return colors
}

func reverseSlice(path []string) {
	for left, right := 0, len(path)-1; left < right; left, right = left+1, right-1 {
		path[left], path[right] = path[right], path[left]
	}
}

func getAdjacencyGraph(data map[int]entity.RowString) map[string][]string {
	graph := make(map[string][]string, len(data)*constant.SIZE_COL)
	for k1 := range data {
		var arrA []string
		var arrB []string
		var arrCell []string
		keyA := fmt.Sprintf("A%d", k1)
		keyB := fmt.Sprintf("B%d", k1)
		keyCell := fmt.Sprintf("Cell%d", k1)
		for k2, v2 := range data {
			arrA = append(arrA, isExist(keyA, k2, v2)...)
			arrB = append(arrB, isExist(keyB, k2, v2)...)
			arrCell = append(arrCell, isExist(keyCell, k2, v2)...)
		}
		if len(arrA) != 0 {
			graph[keyA] = arrA
		}
		if len(arrB) != 0 {
			graph[keyB] = arrB
		}
		if len(arrCell) != 0 {
			graph[keyCell] = arrCell
		}
	}
	return graph
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
