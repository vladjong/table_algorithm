package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vladjong/test_yadro/internal/entity"
)

func TestGetTable(t *testing.T) {
	tests := []struct {
		name  string
		data  map[int]entity.RowString
		path  []string
		out   []entity.RowString
		error bool
	}{
		{
			name:  "correct",
			out:   []entity.RowString{{"1", "3", "4"}, {"5", "8", "13"}, {"1", "4", "4"}},
			data:  map[int]entity.RowString{1: {"1", "3", "4"}, 2: {"=B5+A1", "8", "=Cell5+B5+A2"}, 5: {"=A1", "=A5+B1", "=Cell1"}},
			path:  []string{"A1", "B2", "A5", "Cell1", "Cell5", "B1", "B5", "A2", "Cell2"},
			error: false,
		},
		{
			name:  "refers to itself",
			data:  map[int]entity.RowString{1: {"A1", "3", "4"}, 2: {"=B5+A1", "8", "=Cell5+B5+A2"}, 5: {"=A1", "=A5+B1", "=Cell1"}},
			path:  []string{"A1", "B2", "A5", "Cell1", "Cell5", "B1", "B5", "A2", "Cell2"},
			error: true,
		},
		{
			name:  "refers to dont exist cell",
			data:  map[int]entity.RowString{1: {"A15", "3", "4"}, 2: {"=B5+A1", "8", "=Cell5+B5+A2"}, 5: {"=A1", "=A5+B1", "=Cell1"}},
			path:  []string{"A1", "B2", "A5", "Cell1", "Cell5", "B1", "B5", "A2", "Cell2"},
			error: true,
		},
		{
			name:  "work to number",
			out:   []entity.RowString{{"15", "3", "4"}, {"33", "8", "55"}, {"15", "18", "4"}},
			data:  map[int]entity.RowString{1: {"=5+10", "3", "4"}, 2: {"=B5+A1", "8", "=Cell5+B5+A2"}, 5: {"=A1", "=A5+B1", "=Cell1"}},
			path:  []string{"A1", "B2", "A5", "Cell1", "Cell5", "B1", "B5", "A2", "Cell2"},
			error: false,
		},
	}

	for _, test := range tests {
		data, err := New(test.data, test.path).GetTable()
		if test.error {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, test.out[0], data[1])
		assert.Equal(t, test.out[1], data[2])
		assert.Equal(t, test.out[2], data[5])
	}
}
