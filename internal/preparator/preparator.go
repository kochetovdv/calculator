package preparator

import (
	"calculator/pkg/transformation"
	"strings"
)

// TODO: make structure for automatically generated allowed symbols
var (
//	layoutDigits   = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
//	layoutOperands = [...]string{"+", "-", "*", "/"}
//	layoutGroup    = [...]string{"(", ")"}

// layoutSpaces   = []string{" ", "\t", "\n"}
)

type Preparator struct {
	AllowedSymbols    []rune
	AllowedOperators  []rune
	AllowedGroupOpen  []rune
	AllowedGroupClose []rune
	AllowedDigits     []rune
}

func NewPreparator(digits, operators, groupOpen, groupClose []rune) *Preparator {
	p := new(Preparator)
	p.AllowedOperators = operators
	p.AllowedGroupOpen = groupOpen
	p.AllowedGroupClose = groupClose
	p.AllowedDigits = digits
	p.AllowedSymbols = append(p.AllowedSymbols, p.AllowedOperators...)
	p.AllowedSymbols = append(p.AllowedSymbols, p.AllowedGroupOpen...)
	p.AllowedSymbols = append(p.AllowedSymbols, p.AllowedGroupClose...)
	p.AllowedSymbols = append(p.AllowedSymbols, p.AllowedDigits...)
	return p
}

func (p *Preparator) Preparate(expression string) ([]string, error) {

	expression = transformation.StringsConcatination("(", expression)
	expression = transformation.StringsConcatination(expression, ")")

	var slice []rune
	for _, el := range expression {
		slice = append(slice, el)
	}
	expression = transformation.AddSymbolBeforeAndAfterAnotherSymbol(slice, ' ', p.AllowedOperators...)
	slice = nil
	for _, el := range expression {
		slice = append(slice, el)
	}
	expression = transformation.AddSymbolAfterAnotherSymbol(slice, ' ', p.AllowedGroupOpen...)

	slice = nil
	for _, el := range expression {
		slice = append(slice, el)
	}
	expression = transformation.AddSymbolBeforeAnotherSymbol(slice, ' ', p.AllowedGroupClose...)

	expression = strings.Replace(expression, "  ", " ", -1) // erasing duplicate spaces, for exmple, with +(
	expressionArr := transformation.Split(expression, " ")

	return expressionArr, nil
}
