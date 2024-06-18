package logical

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDecodeWaysNumber(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "",
			output: 0,
		},
		{
			input:  "06",
			output: 0,
		},
		{
			input:  "9",
			output: 1,
		},
		{
			input:  "99",
			output: 1,
		},
		{
			input:  "00",
			output: 0,
		},
		{
			input:  "1201234",
			output: 3,
		},
		{
			input:  "1123",
			output: 5,
		},
		{
			input:  "10011",
			output: 0,
		},
		{
			input:  "12",
			output: 2,
		},
		{
			input:  "102",
			output: 1,
		},
		{
			input:  "6032",
			output: 0,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("input is \"%s\"", test.input), func(t *testing.T) {
			assert.Equal(t, test.output, GetDecodeWaysNumber(test.input))
		})
	}
}
