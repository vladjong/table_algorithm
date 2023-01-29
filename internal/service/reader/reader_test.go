package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vladjong/test_yadro/internal/entity"
)

func TestReaderFirstLine(t *testing.T) {
	tests := []struct {
		name  string
		in    []string
		error bool
	}{
		{
			name:  "correct line",
			in:    []string{"", "A", "B", "Cell"},
			error: false,
		},
		{
			name:  "incorrect line",
			in:    []string{"A1", "B", "Cell"},
			error: true,
		},
	}

	for _, test := range tests {
		err := checkFirstLine(test.in)
		if test.error {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestReaderSizeLine(t *testing.T) {
	tests := []struct {
		name  string
		in    []string
		error bool
	}{
		{
			name:  "correct size line",
			in:    []string{"1", "2", "3", "4"},
			error: false,
		},
		{
			name:  "incorrect line",
			in:    []string{"A1", "B"},
			error: true,
		},
	}

	for _, test := range tests {
		err := checkSize(test.in)
		if test.error {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestReaderGetKeyAnValue(t *testing.T) {
	tests := []struct {
		name  string
		in    []string
		key   int
		value entity.RowString
		error bool
	}{
		{
			name:  "correct line",
			in:    []string{"", "1", "2", "3"},
			value: entity.RowString{A: "1", B: "2", Cell: "3"},
			key:   1,
			error: false,
		},
		{
			name:  "incorrect line key",
			in:    []string{"A1", "B"},
			error: true,
		},
	}

	for _, test := range tests {
		err := checkSize(test.in)
		if test.error {
			assert.NotNil(t, err)
		}
	}
}
