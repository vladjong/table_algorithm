package dfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vladjong/test_yadro/internal/entity"
)

func TestGetPath(t *testing.T) {
	tests := []struct {
		name  string
		in    map[int]entity.RowString
		out   int
		error bool
	}{
		{
			name:  "correct",
			out:   9,
			in:    map[int]entity.RowString{1: {"1", "3", "4"}, 2: {"=B5+A1", "8", "=Cell5+B5+A2"}, 5: {"=A1", "=A5+B1", "=Cell1"}},
			error: false,
		},
		{
			name:  "circle table",
			out:   0,
			in:    map[int]entity.RowString{1: {"=A2", "3", "4"}, 2: {"=A1", "8", "=Cell5+B5+A2"}, 5: {"=A1", "=A5+B1", "=Cell1"}},
			error: true,
		},
	}

	for _, test := range tests {

		data, err := New(test.in).GetPath()
		if test.error {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, test.out, len(data))
	}
}
