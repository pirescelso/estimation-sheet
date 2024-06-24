package common_test

import (
	"fmt"
	"testing"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestRoundToTwoDecimals(t *testing.T) {
	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "zero",
			input:    0.0,
			expected: 0.00,
		},
		{
			name:     "positive number down",
			input:    1.234,
			expected: 1.23,
		},
		{
			name:     "positive number up",
			input:    1.235,
			expected: 1.24,
		},

		{
			name:     "negative number down",
			input:    -1.234,
			expected: -1.23,
		},
		{
			name:     "negative number up",
			input:    -1.235,
			expected: -1.24,
		},

		{
			name:     "number with zero decimal",
			input:    1.00,
			expected: 1.00,
		},
		{
			name:     "number with zero decimal and negative",
			input:    -1.0,
			expected: -1.0,
		},
		{
			name:     "number 2.235",
			input:    2.235,
			expected: 2.24,
		},
		{
			name:     "number 4337.025",
			input:    4337.025,
			expected: 4337.03,
		},
		{
			name:     "number 1.12445",
			input:    1.12445,
			expected: 1.12,
		},

		{
			name:     "number 3.2349999999999994",
			input:    3.2349999999999994,
			expected: 3.23,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := common.RoundToTwoDecimals(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func FuzzRoundToTwoDecimals(f *testing.F) {
	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "zero",
			input:    0.00,
			expected: 0.00,
		},
		{
			name:     "positive number down",
			input:    1.234,
			expected: 1.23,
		},
		{
			name:     "positive number up",
			input:    1.235,
			expected: 1.24,
		},

		{
			name:     "negative number down",
			input:    -1.234,
			expected: -1.23,
		},
		{
			name:     "negative number up",
			input:    -1.235,
			expected: -1.24,
		},

		{
			name:     "number with zero decimal",
			input:    1.00,
			expected: 1.00,
		},
		{
			name:     "number with zero decimal and negative",
			input:    -1.0,
			expected: -1.0,
		},
	}

	for _, tc := range testCases {
		f.Add(tc.input)
	}
	f.Fuzz(func(t *testing.T, number float64) {
		result := common.RoundToTwoDecimals(number)
		if number > 0.00 {
			assert.Equal(t, fmt.Sprintf("%.2f", number), fmt.Sprintf("%.2f", result))
		}

	})
}
