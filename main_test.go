package main

import (
	"testing"
)

func Test_GridPress(t *testing.T) {
	testcases := []struct {
		name       string
		pos        int
		startState string
		endState   string
	}{
		{"Press 0", 0, "000000000", "111100100"},
		{"Press 1", 1, "000000000", "111010010"},
		{"Press 2", 2, "000000000", "111001001"},
		{"Press 3", 3, "000000000", "100111100"},
		{"Press 4", 4, "000000000", "010111010"},
		{"Press 5", 5, "000000000", "001111001"},
		{"Press 6", 6, "000000000", "100100111"},
		{"Press 7", 7, "000000000", "010010111"},
		{"Press 8", 8, "000000000", "001001111"},
	}

	for _, tc := range testcases {
		grid := grid([]byte(tc.startState))
		grid.Press(tc.pos)
		if grid.String() != tc.endState {
			t.Errorf("expected: %s, got: %s", tc.endState, grid)
		}
	}

}
