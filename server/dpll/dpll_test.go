package dpll

import (
	"testing"

	"solver/formula"
)

type possibleAnswer struct {
	expectedAssignment map[string]bool
	expectedDontCare   []string
}

type testCase struct {
	input         [][]string
	expectedValid bool
	// one form can have multiple answers, I try my best to cover all of them
	possibleAnswers []possibleAnswer
}

func TestSolve(t *testing.T) {
	tests := []testCase{
		{
			input:         [][]string{{"1", "-1"}},
			expectedValid: true,
			possibleAnswers: []possibleAnswer{{
				expectedAssignment: map[string]bool{"1": false},
				expectedDontCare:   []string{},
			},
				{
					expectedAssignment: map[string]bool{"1": true},
					expectedDontCare:   []string{},
				},
			},
		},
		{
			input:         [][]string{{"-1", "2"}, {"-1"}},
			expectedValid: true,
			possibleAnswers: []possibleAnswer{
				{
					expectedAssignment: map[string]bool{"1": false},
					expectedDontCare:   []string{"2"},
				},
			},
		},
		{
			input:         [][]string{{"1"}, {"-1"}},
			expectedValid: false,
		},
		{
			input:         [][]string{{"1", "2"}, {"-1", "-2", "3"}},
			expectedValid: true,
			possibleAnswers: []possibleAnswer{
				{
					expectedAssignment: map[string]bool{"2": true, "3": true},
					expectedDontCare:   []string{"1"},
				},
			},
		},
		{
			input: [][]string{
				{"-x1", "-x2"},
				{"-x1", "x2"},
				{"x1", "-x2"},
				{"x2", "-x3"},
				{"x1", "x3"},
			},
			expectedValid: false,
		},
		{
			input: [][]string{
				{"-x1", "-x2"},
				{"x1", "-x2"},
				{"-x1", "-x3"},
			},
			expectedValid: true,
			possibleAnswers: []possibleAnswer{
				{
					expectedAssignment: map[string]bool{"x2": false, "x3": false},
					expectedDontCare:   []string{"x1"},
				},
			},
		},
		{
			input: [][]string{
				{"x2", "x1"},
				{"-x1"},
				{"-x2"},
				{"-x2", "-x3"},
				{"x3", "x1"},
			},
			expectedValid: false,
		},
		{
			input: [][]string{
				{"-x2", "-x1"},
				{"-x1"},
				{"-x3", "-x4"},
			},
			expectedValid: true,
			possibleAnswers: []possibleAnswer{
				{
					expectedAssignment: map[string]bool{"x1": false, "x3": false},
					expectedDontCare:   []string{"x2", "x4"},
				},
				{
					expectedAssignment: map[string]bool{"x1": false, "x4": false},
					expectedDontCare:   []string{"x2", "x3"},
				},
			},
		},
	}

	for i, test := range tests {
		f := formula.NewFormula(test.input)
		testAll(t, i, f, test)
	}
}

func testAnswers(t *testing.T, i int, f *formula.Formula, expectedAnswers []possibleAnswer) {

	var strDontCare []string
	for s := range f.DontCare {
		strDontCare = append(strDontCare, s)
	}

	got := possibleAnswer{
		expectedAssignment: f.Assignments,
		expectedDontCare:   strDontCare,
	}

	for _, answer := range expectedAnswers {
		if testAssignments(f.Assignments, answer.expectedAssignment) && testDontCare(strDontCare, answer.expectedDontCare) {
			return
		}
	}

	t.Fatalf(
		"case %d: wrong answer, expected one of=%v, got=%v",
		i+1,
		expectedAnswers,
		got,
	)
}

func testAssignments(assign map[string]bool, expected map[string]bool) bool {
	if len(expected) != len(assign) {
		return false
	}

	for str, value := range assign {
		if expectedValue, exists := expected[str]; !exists || expectedValue != value {
			return false
		}
	}
	return true
}

func testDontCare(dc []string, expected []string) bool {
	if len(expected) != len(dc) {
		return false
	}
	if !areSlicesEqual(dc, expected) {
		return false
	}

	return true
}

func testAll(t *testing.T, i int, f *formula.Formula, test testCase) {
	valid, _, _ := Solve(f)
	if valid != test.expectedValid {
		t.Errorf(
			"case %d: validity expected=%t, got=%t",
			i+1,
			test.expectedValid,
			valid,
		)
	}
	if valid {
		testAnswers(t, i, f, test.possibleAnswers)
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
