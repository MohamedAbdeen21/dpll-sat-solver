package clause

import (
	"strings"

	"solver/literal"
)

type Clause struct {
	Literals []*literal.Literal
	Unit     bool
	Empty    bool
}

// Create a new Clause from a line of inputs (array of literals)
func NewClause(clauseLiterals []string) *Clause {
	var lit []*literal.Literal
	for _, l := range clauseLiterals {
		if l[0] == '-' || l[0] == '+' {
			lit = append(lit, literal.NewLiteral(l[0], l[1:]))
		} else {
			lit = append(lit, literal.NewLiteral('+', l))
		}
	}

	c := &Clause{
		Literals: lit,
	}

	c.setUnit()
	c.setEmpty()

	return c
}

// A string representation of the clause, mainly for testing and debugging
func (c *Clause) String() string {
	var LiteralsNames []string
	for _, literal := range c.Literals {
		LiteralsNames = append(LiteralsNames, literal.String())
	}
	return "(" + strings.Join(LiteralsNames, " v ") + ")"
}

// Remove a literal from the clause and re-set the Unit
// and Empty flags, this is polarity-sensitive
func (c *Clause) RemoveLiteral(name string) {
	var newLiterals []*literal.Literal
	for _, lit := range c.Literals {
		if lit.String() != name {
			newLiterals = append(newLiterals, lit)
		}
	}

	c.Literals = newLiterals
	c.setUnit()
	c.setEmpty()
}

// helper to determine if clause contains a literal
func (c *Clause) Contains(name string) bool {
	for _, lit := range c.Literals {
		if lit.String() == name {
			return true
		}
	}

	return false
}

// helper to re-set the Unit flag
func (c *Clause) setUnit() {
	c.Unit = len(c.Literals) == 1
}

// helper to re-set the Empty flag
func (c *Clause) setEmpty() {
	c.Empty = len(c.Literals) == 0
}
