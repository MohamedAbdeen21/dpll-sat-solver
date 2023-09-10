package formula

import (
	"solver/literal"
	"testing"
)

type testCase struct {
	input              [][]string
	expected           string
	expectedPures      []string
	expectedUnits      []string
	expectedEmptyCount int
	toBeRemoved        string
}

func TestNewFormula(t *testing.T) {
	tests := []testCase{
		{
			input:              [][]string{{"1", "2"}, {"3", "4"}, {"-3"}, {"-2"}},
			expected:           "(1 v 2) ^ (3 v 4) ^ (-3) ^ (-2)",
			expectedPures:      []string{"1", "4"},
			expectedUnits:      []string{"-3", "-2"},
			expectedEmptyCount: 0,
		},
		{
			input:              [][]string{{"x1"}},
			expected:           "(x1)",
			expectedPures:      []string{"x1"},
			expectedUnits:      []string{"x1"},
			expectedEmptyCount: 0,
		},
		{
			input:              [][]string{{"1", "2"}, {"-2", "3"}, {"-3", "2"}},
			expected:           "(1 v 2) ^ (-2 v 3) ^ (-3 v 2)",
			expectedPures:      []string{"1"},
			expectedUnits:      []string{},
			expectedEmptyCount: 0,
		},
		{
			input:              [][]string{{"-1", "2"}, {"-2", "3"}, {"-3", "2"}, {"1"}},
			expected:           "(-1 v 2) ^ (-2 v 3) ^ (-3 v 2) ^ (1)",
			expectedPures:      []string{},
			expectedUnits:      []string{"1"},
			expectedEmptyCount: 0,
		},
	}

	for i, test := range tests {
		f := NewFormula(test.input)
		testAll(t, i, f, test)
	}
}

func TestClauseRemoval(t *testing.T) {
	tests := []testCase{
		{
			input:              [][]string{{"1", "2"}, {"3", "4"}, {"-3"}, {"-2"}},
			expected:           "(-3) ^ (-2)",
			expectedPures:      []string{"3", "2"}, // new pures after removal of original pures
			expectedUnits:      []string{"-3", "-2"},
			expectedEmptyCount: 0,
		},
		{
			input:              [][]string{{"x1"}, {"-x1"}},
			expected:           "(x1) ^ (-x1)",
			expectedPures:      []string{},
			expectedUnits:      []string{"x1", "-x1"},
			expectedEmptyCount: 0,
		},
	}

	for i, test := range tests {
		f := NewFormula(test.input)
		for pure, polarity := range f.Pures {
			f.RemoveClausesContainingLiteral(pure, polarity)
		}
		testAll(t, i, f, test)
	}
}

func TestLiteralRemovalFromClause(t *testing.T) {
	tests := []testCase{
		{
			input:              [][]string{{"1", "2"}, {"3", "4"}, {"-3"}, {"-2"}},
			expected:           "(2) ^ (3 v 4) ^ (-3) ^ (-2)",
			expectedPures:      []string{"4"},
			expectedUnits:      []string{"2", "-3", "-2"},
			expectedEmptyCount: 0,
			toBeRemoved:        "-1",
		},
		{
			input:              [][]string{{"x1"}, {"-x1"}},
			expected:           "() ^ (-x1)",
			expectedPures:      []string{"x1"},
			expectedUnits:      []string{"-x1"},
			expectedEmptyCount: 1,
			toBeRemoved:        "-x1",
		},
	}

	for i, test := range tests {
		f := NewFormula(test.input)
		f.RemoveLiteralNegationFromClauses(test.toBeRemoved)
		testAll(t, i, f, test)
	}
}

func TestAddingClause(t *testing.T) {
	tests := []testCase{
		{
			input:              [][]string{{"1", "2"}},
			expected:           "(1 v 2) ^ (3)",
			expectedPures:      []string{"1", "2", "3"},
			expectedUnits:      []string{"3"},
			expectedEmptyCount: 0,
		},
	}

	for i, test := range tests {
		f := NewFormula(test.input)
		newF := f.Add(literal.NewLiteral('+', "3"))
		if f == newF {
			t.Errorf("case %d: Add() didn't create a copy of original struct", i+1)
		}
		testAll(t, i, newF, test)
	}
}

func areSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	seen := make(map[string]bool)
	for _, s := range a {
		seen[s] = true
	}

	for _, s := range b {
		if !seen[s] {
			return false
		}
		seen[s] = false
	}

	return true
}

func testFormula(t *testing.T, id int, f *Formula, expected string) {
	if f.String() != expected {
		t.Errorf(
			"case %d: Formula not parsed correctly, expected=%s, got=%s",
			id+1,
			expected,
			f.String(),
		)
	}
}

func testUnits(t *testing.T, id int, actual map[string]literal.Polarity, expected []string) {
	if len(actual) == 0 && len(expected) == 0 {
		return
	}

	var str []string

	for s := range actual {
		str = append(str, s)
	}

	if !areSlicesEqual(str, expected) {
		t.Errorf(
			"case %d: Units are not identified correctly, expected=%v, got=%v",
			id+1,
			expected,
			str,
		)
	}
}

func testPures(t *testing.T, id int, actual map[string]literal.Polarity, expected []string) {
	if len(actual) == 0 && len(expected) == 0 {
		return
	}

	var str []string

	for s := range actual {
		str = append(str, s)
	}

	if !areSlicesEqual(str, expected) {
		t.Errorf(
			"case %d: Pures are not identified correctly, expected=%v, got=%v",
			id+1,
			expected,
			str,
		)
	}
}

func testEmptyCount(t *testing.T, id int, actual, expected int) {
	if actual != expected {
		t.Errorf(
			"case %d: Empty are not identified correctly, expected=%d, got=%d",
			id+1,
			expected,
			actual,
		)
	}
}

func testAll(t *testing.T, i int, actual *Formula, test testCase) {
	testFormula(t, i, actual, test.expected)
	testPures(t, i, actual.Pures, test.expectedPures)
	testUnits(t, i, actual.Units, test.expectedUnits)
	testEmptyCount(t, i, len(actual.EmptyClauses), test.expectedEmptyCount)
}
