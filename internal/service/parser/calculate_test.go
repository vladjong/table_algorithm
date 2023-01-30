package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	tests := []struct {
		name string
		in   rune
		out  bool
	}{
		{
			name: "number",
			in:   '2',
			out:  true,
		},
		{
			name: "not number",
			in:   'a',
			out:  false,
		},
	}

	for _, test := range tests {
		ch := isNumber(test.in)
		assert.Equal(t, test.out, ch)
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		out   int
		error bool
	}{
		{
			name:  "sum",
			in:    "2+3+5",
			out:   10,
			error: false,
		},
		{
			name:  "sub",
			in:    "2-3-5",
			out:   -6,
			error: false,
		},
		{
			name:  "div",
			in:    "10/5",
			out:   2,
			error: false,
		},
		{
			name:  "mul",
			in:    "10*5",
			out:   50,
			error: false,
		},
		{
			name:  "hard case",
			in:    "100-10*5/10+1",
			out:   96,
			error: false,
		},
		{
			name:  "incorrect operator",
			in:    "10>5",
			out:   0,
			error: true,
		},
		{
			name:  "div by zero",
			in:    "10/0",
			out:   0,
			error: true,
		},
	}

	for _, test := range tests {
		in, err := calculate(test.in)
		if test.error {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, test.out, in)
	}
}
