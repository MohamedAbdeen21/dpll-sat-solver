package literal

import "testing"

type testCase struct {
	inputName string
	inputSign byte
	expected  string
}

func TestNewLiteral(t *testing.T) {
	tests := []testCase{
		{
			inputName: "x",
			inputSign: '-',
			expected:  "-x",
		},
		{
			inputName: "z",
			inputSign: 0,
			expected:  "z",
		},
		{
			inputName: "20",
			inputSign: '+',
			expected:  "20",
		},
	}

	for i, test := range tests {
		lit := NewLiteral(test.inputSign, test.inputName)
		if lit.String() != test.expected {
			t.Errorf("case %d: literal expected=%s, got=%s", i+1, lit.String(), test.expected)
		}
	}
}

func TestNegation(t *testing.T) {
	tests := []testCase{
		{
			inputName: "x",
			inputSign: '-',
			expected:  "x",
		},
		{
			inputName: "z",
			inputSign: 0,
			expected:  "-z",
		},
		{
			inputName: "20",
			inputSign: '+',
			expected:  "-20",
		},
	}

	for i, test := range tests {
		lit := NewLiteral(test.inputSign, test.inputName).Negate()
		if lit.String() != test.expected {
			t.Errorf("case %d: literal expected=%s, got=%s", i+1, lit.String(), test.expected)
		}
	}
}
