package formula

import (
	"solver/clause"
	"solver/literal"
	"strings"
)

// a formula is an AND between Clauses. Also called CNF or Product of Sums
type Formula struct {
	Clauses      []*clause.Clause
	EmptyClauses []*clause.Clause
	Units        map[string]literal.Polarity
	Pures        map[string]literal.Polarity
	Assignments  map[string]bool
	DontCare     map[string]bool
}

// Create a new Formula from multiple lines of inputs (clause per line)
func NewFormula(input [][]string) *Formula {
	f := &Formula{
		Units:       make(map[string]literal.Polarity),
		Pures:       make(map[string]literal.Polarity),
		Assignments: make(map[string]bool),
		DontCare:    make(map[string]bool),
	}

	for _, c := range input {
		newClause := clause.NewClause(c)
		f.Clauses = append(f.Clauses, newClause)
	}

	f.setPures()
	f.setUnits()

	return f
}

// Remove all clauses where the Literal name appears, polarity sensitive (i.e.
// calling this with literal +1 doesn't remove clauses with -1)
func (f *Formula) RemoveClausesContainingLiteral(name string, polarity literal.Polarity) {
	// make sure it wasn't removed by a previous call
	f.assign(name, bool(polarity))

	if polarity == literal.Negative {
		name = negate(name)
	}

	for _, clause := range f.Clauses {
		if clause.Contains(name) {
			f.markRestAsDontCare(clause, name)
			f.removeClause(clause)
		}
	}

	f.setPures()
	// don't need to call f.setEmpty or f.setUnits, because we only remove clauses not
	// literals
}

// Remove the negation of the given literal from all clauses. If target clause
// is a unit clause, leaves behind an empty clause. Meaning formula not valid
func (f *Formula) RemoveLiteralNegationFromClauses(name string) {
	name = negate(name)

	for _, clause := range f.Clauses {
		clause.RemoveLiteral(name)
	}

	f.setUnits()
	f.setEmpty()
	f.setPures()
}

// A string representation of the formula, mainly for testing and debugging
func (f *Formula) String() string {
	var formulaString []string
	for _, c := range f.Clauses {
		formulaString = append(formulaString, c.String())
	}

	return strings.Join(formulaString, " ^ ")
}

func (f Formula) Add(l *literal.Literal) *Formula {
	fCopy := f.copy()
	newClause := &clause.Clause{Literals: []*literal.Literal{l}, Unit: true}
	fCopy.Clauses = append(fCopy.Clauses, newClause)
	fCopy.setUnits()
	fCopy.setPures()

	return fCopy
}

// creates a copy of the formula (notice the reciever is not a pointer),
// helps with recursive calls to solve() function
func (f Formula) copy() *Formula {
	return &f
}

// remove a clause from the formula, re-setting pures is handled after
// call in RemoveClausesContainingLiteral()
func (f *Formula) removeClause(c *clause.Clause) {
	var newClauses []*clause.Clause
	for _, clause := range f.Clauses {
		if clause != c {
			newClauses = append(newClauses, clause)
		}
	}

	f.Clauses = newClauses
}

// re-set the Units in the Formula struct, called after removing literals from clauses
func (f *Formula) setUnits() {
	units := make(map[string]literal.Polarity)
	for _, clause := range f.Clauses {
		if clause.Unit {
			units[clause.Literals[0].String()] = clause.Literals[0].Polarity
		}
	}

	f.Units = units
}

// re-set the Empty in the Formula struct, called after removing literals from
// clauses
func (f *Formula) setEmpty() {
	var empty []*clause.Clause
	for _, clause := range f.Clauses {
		if clause.Empty {
			empty = append(empty, clause)
		}
	}

	f.EmptyClauses = empty
}

// re-set the Pures in the formula struct, called after removing clauses or
// literals from clauses
func (f *Formula) setPures() {
	literalPolarity := make(map[string]literal.Polarity)
	removed := make(map[string]bool)
	for _, clause := range f.Clauses {
		for _, literal := range clause.Literals {
			if currPolarity, exists := literalPolarity[literal.Name]; exists {
				if currPolarity != literal.Polarity { // impure
					delete(literalPolarity, literal.Name)
					removed[literal.Name] = true
				}
			} else if _, wasRemoved := removed[literal.Name]; !wasRemoved {
				literalPolarity[literal.Name] = literal.Polarity
			}
		}
	}

	pures := make(map[string]literal.Polarity)
	for literal, polarity := range literalPolarity {
		pures[literal] = polarity
	}
	f.Pures = pures
}

// Mark the rest of the literals as Don't Care, called after removing a clause
// that contains a pure or a unit
func (f *Formula) markRestAsDontCare(c *clause.Clause, name string) {
	for _, literal := range c.Literals {
		if literal.String() != name && f.isUnique(literal) {
			f.DontCare[literal.Name] = true
		}
	}
}

// Need to know if a literal is unique before assigning it as a don't care
// if not unique, this means that it gets a value later
func (f *Formula) isUnique(lit *literal.Literal) bool {
	var wasSeen bool = false
	for _, clause := range f.Clauses {
		for _, literal := range clause.Literals {
			if literal.Name == lit.Name {
				if !wasSeen {
					wasSeen = true
				} else {
					return false
				}
			}
		}
	}

	return true
}

// assign the literal a value (true/false)
func (f *Formula) assign(lit string, value bool) {
	// quick and dirty workaround: instead of assigning negation to true, assign
	// literals to false
	if lit[0] == '-' {
		f.Assignments[lit[1:]] = false
	} else {
		f.Assignments[lit] = value
	}
}

// helper function to negate the given string
func negate(name string) string {
	if name[0] == '-' {
		return name[1:]
	} else {
		return "-" + name
	}
}
