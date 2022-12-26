package prevalidator

import (
	"calculator/pkg/errorchecking"
	"calculator/pkg/transformation"
	"strconv"
)

type Prevalidator struct {
	Expression string
}

func NewPrevalidator() *Prevalidator {
	p := new(Prevalidator)
	return p
}

func (p *Prevalidator) Prevalidate(result int, expression string) (string, error) {
	// Base spaces and tabulations erasing. It is need for future parsing: 1  +2 -> 1+2
	p.Expression = expression
	p.spaceDeleting()

	// Result adding at expression if it starts with and operator: +1*2 -> result+1*2
	p.resultAdding(result)

	// Empty brackets () changing for zero values: *()+ -> *0+
	// If brackets are just deleting then operator dublicating case possible to occure: *()+ -> *+
	p.emptyDoubleBracketsReplace()

	return p.Expression, nil
}

// spaceDeleting deletes all spacy symbols: \s,\t,\n
func (p *Prevalidator) spaceDeleting() {
	pattern := `[\s]+`
	p.Expression = transformation.SymbolDeleting(pattern, p.Expression)
}

// resultAdding adds result in the start of expression if it starts with an operator
func (p *Prevalidator) resultAdding(numberToAdd int) {
	check := operatorWithoutOperandsAtStarttValidation(p.Expression)
	if check {
		p.Expression = transformation.StringsConcatination(strconv.Itoa(numberToAdd), p.Expression)
	}
}

// emptyDoubleBracketsReplace changes all empty brackets () to zero value
func (p *Prevalidator) emptyDoubleBracketsReplace() {
	check := doubleEmptyBracketsValidation(p.Expression)
	if check {
		pattern := `\(\)`
		p.Expression = transformation.ReplaceAllSymbols(pattern, p.Expression, "0")
	}
}

// operatorWithoutOperandsAtStarttValidation checks expression starting with an operator
func operatorWithoutOperandsAtStarttValidation(s string) bool {
	// Pattern is formed via operator places at the start.
	pattern := `^[\/\+\-\*][\d\)]*`
	check := errorchecking.RegexpPatternCheck(pattern, s)
	return check
}

// doubleEmptyBracketsValidation chekcs for emprty brackets ()
func doubleEmptyBracketsValidation(s string) bool {
	pattern := `(\(\))*`
	check := errorchecking.RegexpPatternCheck(pattern, s)
	return check
}
