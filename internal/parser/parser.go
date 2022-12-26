package parser

import (
	"calculator/internal/model"
	"fmt"
	"unicode"
)

type Parser struct {
	Operators []model.Operator
	Groups    []model.Group
	Numbers   []rune
}

func NewParser(operators []model.Operator, groups []model.Group) *Parser {
	p := new(Parser)
	p.Operators = operators
	p.Groups = groups
	return p
}

var (
	IDChain []*model.ID
	id      *model.ID
)

func (p *Parser) Parse(expressionsArr []string) (model.Expression, error) {
	var currentExpressionPriority int // variable is used for expressions in brackets prioritization

	operators := make([]*model.Operator, 0) // array for parsed operators in expression
	numbers := make([]*model.Number, 0)     // array for parsed deigits in expression
	var errParsing error
	// Parsing elements from string array, checking are the elements operators, digits or brackets
	for _, exp := range expressionsArr {
		switch {
		case p.validateNumber(exp):
			num := model.NewNumber(exp)
			num.ID = updateID()
			numbers = append(numbers, num)
		case p.validateOperator(exp):
			op, err := p.GetOperator(exp) // Getting operator by symbol
			if err != nil {
				errParsing = err
			}
			op.ID = updateID()
			// Operator priority is equal current part of expression priority plused default priority of operator
			op.Priority = currentExpressionPriority + op.Priority
			operators = append(operators, &op)
		case p.validateGroup(exp): 
			gr, err := p.GetGroup(exp) // Getting bracket by symbol
			if err != nil {
				errParsing = err
			}
			// Bracket changes current part of expression priority. 
			currentExpressionPriority = currentExpressionPriority + gr.Priority
		}
	}
	var expressions model.Expression
	expressions.Numbers = numbers
	expressions.Operators = operators

	for _, el := range expressions.Numbers {
		el.Convert()
	}
	return expressions, errParsing
}

// Method validates is the symbol an operator
func (p *Parser) validateOperator(s string) bool {
	for _, op := range p.Operators {
		if s == string(op.Symbol) {
			return true
		}
	}
	return false
}

// Method returns an operator by symbol
func (p *Parser) GetOperator(s string) (model.Operator, error) {
	for _, op := range p.Operators {
		if s == string(op.Symbol) {
			return op, nil
		}
	}
	return model.Operator{}, fmt.Errorf("unknown operator")
}

// Method validates is the symbol a bracket
func (p *Parser) validateGroup(s string) bool {
	for _, gr := range p.Groups {
		if s == string(gr.Symbol) {
			return true
		}
	}
	return false
}

// Method returns a bracket by symbol
func (p *Parser) GetGroup(s string) (model.Group, error) {
	for _, gr := range p.Groups {
		if s == string(gr.Symbol) {
			return gr, nil
		}
	}
	return model.Group{}, fmt.Errorf("unknown bracket")
}

// Method validates is the string a digit
// TODO Remake for a float
func (p *Parser) validateNumber(s string) bool {
	for _, c := range s {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

// Method makes new ID using global IDChain array
func updateID() *model.ID {
	i := len(IDChain)
	if i == 0 {
		id = model.NewID(i)
	} else {
		id = model.NewIDChain(IDChain[i-1])
	}
	IDChain = append(IDChain, id)
	return id
}
