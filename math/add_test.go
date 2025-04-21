package math_test

import (
	"example/math"
	"testing"
)

type addTestCase struct {
	a, b, expected int
}

var testCases = []addTestCase{
	{1, 1, 2},
	{25, 25, 50},
	{2, 1, 3},
	{1, 10, 11},
}

func TestAdd(t *testing.T) {
	for _, tc := range testCases {
		got := math.Add(tc.a, tc.b)

		if got != tc.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.expected)
		}
	}
}