package flag

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlag(t *testing.T) {
	tests := []struct {
		name  string
		in    []string
		out   string
		error bool
	}{
		{
			name:  "correct filename",
			in:    []string{"exe", "test.csv"},
			out:   "test.csv",
			error: false,
		},
		{
			name:  "incorrect filename",
			in:    []string{"exe", "test.txt"},
			out:   "",
			error: true,
		},
		{
			name:  "empty file",
			in:    []string{"exe"},
			out:   "",
			error: true,
		},
	}

	for _, test := range tests {
		os.Args = test.in
		name, err := GetFilename()
		if test.error {
			assert.NotNil(t, err)
			continue
		}
		assert.Equal(t, test.out, name)
	}
}
