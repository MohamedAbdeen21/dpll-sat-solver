package dpll

import (
	"solver/formula"
	"solver/literal"
)

type Solution struct {
	Solved bool     `json:"solved"`
	Trues  []string `json:"trues"`
	Falses []string `json:"falses"`
	Dcs    []string `json:"dcs"`
}

// selects a literal to add to the formula as a Unit clause
func selectLiteral(f *formula.Formula) *literal.Literal {
	var literal *literal.Literal
	for _, clause := range f.Clauses {
		for _, lit := range clause.Literals {
			literal = lit
		}
	}
	return literal
}

// DPLL algorithm, returns valid boolean, map of assignments and set
// of don't care literals
func Solve(f *formula.Formula) (bool, map[string]bool, map[string]bool) {
	for unit := range f.Units {

		if _, exists := f.DontCare[unit]; exists {
			continue
		}

		f.RemoveClausesContainingLiteral(unit, true)
		f.RemoveLiteralNegationFromClauses(unit)
	}

	if len(f.EmptyClauses) != 0 {
		return false, nil, nil
	}

	for pure, polarity := range f.Pures {
		if _, exists := f.DontCare[pure]; exists {
			continue
		}
		f.RemoveClausesContainingLiteral(pure, polarity)
	}

	if len(f.Clauses) == 0 {
		return true, f.Assignments, f.DontCare
	} else {
		literal := selectLiteral(f)
		negated := literal.Negate()
		if valid, assignment, dc := Solve(f.Add(literal)); valid {
			return true, assignment, dc
		} else if valid, assignment, dc := Solve(f.Add(negated)); valid {
			return true, assignment, dc
		} else {
			return false, nil, nil
		}
	}
}

// Entry point to the program
func RunDPLL(input [][]string) Solution {
	sol := Solution{
		Solved: false,
		Trues:  []string{},
		Falses: []string{},
		Dcs:    []string{},
	}

	f := formula.NewFormula(input)
	valid, assignments, dc := Solve(f)

	if valid {
		sol.Solved = true

		for s := range dc {
			sol.Dcs = append(sol.Dcs, s)
		}

		for literal, assignedValue := range assignments {
			if assignedValue {
				sol.Trues = append(sol.Trues, literal)
			} else {
				sol.Falses = append(sol.Falses, literal)
			}
		}
	}

	return sol
}
