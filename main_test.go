package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNumerals(t *testing.T) {
	type Tests struct {
		val      uint
		expected string
	}
	tests := []Tests{
		{
			val:      1,
			expected: "I",
		},
		{
			val:      2,
			expected: "II",
		},
		{
			val:      3,
			expected: "III",
		},
		{
			val:      4,
			expected: "IV",
		},
		{
			val:      5,
			expected: "V",
		},
	}

	for _, ctest := range tests {
		t.Run(ctest.expected, func(t *testing.T) {
			assert.Equal(t, ToRomanNumeral(ctest.val), ctest.expected)
		})
	}
}
