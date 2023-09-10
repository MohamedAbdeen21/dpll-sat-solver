package clause

import "testing"

type testCase struct {
	input            []string
	expected         string
	expectedUnits    bool
	expectedEmpty    bool
	expectedContains bool
	query            string
}

func TestNewClause(t *testing.T) {
	tests := []testCase{
		{
			input:         []string{"1", "2", "-2", "3"},
			expected:      "(1 v 2 v -2 v 3)",
			expectedUnits: false,
			expectedEmpty: false,
		},
		{
			input:         []string{"x1"},
			expected:      "(x1)",
			expectedUnits: true,
			expectedEmpty: false,
		},
	}

	for i, test := range tests {
		clause := NewClause(test.input)
		testAll(t, i, clause, test)
	}
}

func TestLiteralRemoval(t *testing.T) {
	tests := []testCase{
		{
			input:         []string{"1", "2", "-2", "3"},
			expected:      "(1 v -2 v 3)",
			expectedUnits: false,
			expectedEmpty: false,
			query:         "2",
		},
		{
			input:         []string{"x1"},
			expected:      "()",
			expectedUnits: false,
			expectedEmpty: true,
			query:         "x1",
		},
		{
			input:         []string{"x1", "x2"},
			expected:      "(x2)",
			expectedUnits: true,
			expectedEmpty: false,
			query:         "x1",
		},
		{
			input:         []string{"x1", "x2"},
			expected:      "(x1 v x2)",
			expectedUnits: false,
			expectedEmpty: false,
			query:         "x3",
		},
	}

	for i, test := range tests {
		clause := NewClause(test.input)
		clause.RemoveLiteral(test.query)
		testAll(t, i, clause, test)
	}
}

func TestContain(t *testing.T) {
	tests := []testCase{
		{
			input:            []string{"1", "-2", "3"},
			query:            "2",
			expectedContains: false,
		},
		{
			input:            []string{"1", "-2", "2"},
			query:            "2",
			expectedContains: true,
		},
		{
			input:            []string{"x1"},
			query:            "x2",
			expectedContains: false,
		},
		{
			input:            []string{"x1", "x2"},
			query:            "x1",
			expectedContains: true,
		},
	}

	for i, test := range tests {
		clause := NewClause(test.input)
		if clause.Contains(test.query) != test.expectedContains {
			t.Errorf(
				"case %d: clause %v contain %s expected=%t, got=%t",
				i+1,
				clause.Literals,
				test.query,
				test.expectedContains,
				clause.Contains(test.query),
			)
		}
	}
}

func testAll(t *testing.T, i int, clause *Clause, test testCase) {
	if clause.Unit != test.expectedUnits {
		t.Errorf(
			"case %d: clause %v unit expected=%t, got=%t",
			i+1,
			clause,
			clause.Unit,
			test.expectedUnits,
		)
	}
	if clause.Empty != test.expectedEmpty {
		t.Errorf(
			"case %d: clause %v empty expected=%t, got=%t",
			i+1,
			clause,
			clause.Empty,
			test.expectedEmpty,
		)
	}

	if clause.String() != test.expected {
		t.Errorf(
			"case %d: wrong clause parse expected=%s, got=%s",
			i+1,
			clause.String(),
			test.expected,
		)
	}
}
