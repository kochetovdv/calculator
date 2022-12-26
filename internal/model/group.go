package model

// Group is brackets
type Group struct {
	Symbol   rune // bracket symbol
	Priority int  // priority increament or decrement in expression
}

func NewGroup(r rune, p int) *Group {
	g := new(Group)
	g.Symbol = r
	g.Priority = p
	return g
}
