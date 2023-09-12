package literal

type Polarity bool

const (
	Positive Polarity = true
	Negative Polarity = false
)

// Literals have a name and a polarity, call .String() for a string
// reprsentation of literal with polarity
type Literal struct {
	Name     string
	Polarity Polarity
}

// return a new Literal with the polarity
// indicated by sign and the given name
func NewLiteral(sign byte, name string) *Literal {
	l := &Literal{
		Name: name,
	}

	l.setPolarity(sign)
	return l
}

// helper to set the polarity attribute of the struct
func (l *Literal) setPolarity(sign byte) {
	if sign == '-' {
		l.Polarity = Negative
		return
	}

	// 0 is byte's equivalent to nil
	if sign == '+' || sign == 0 {
		l.Polarity = Positive
		return
	}
}

// return a string reprsentation of the literal,
// mainly for testing and debugging
func (l *Literal) String() string {
	if l.Polarity {
		return l.Name
	} else {
		return "-" + l.Name
	}
}

// return a copy of the literal with negated polarity,
// Used in the core DPLL algorithm
func (l Literal) Negate() *Literal {
	lit := &Literal{
		Name:     l.Name,
		Polarity: !l.Polarity,
	}
	return lit
}
