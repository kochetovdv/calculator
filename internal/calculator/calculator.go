package calculator

//TODO refactor Print
//TODO make config file with validations patterns, errors texts 
//TODO make errors more consistent, typical

import (
	"calculator/internal/computer"
	"calculator/internal/input"
	"calculator/internal/model"
	"calculator/internal/parser"
	"calculator/internal/preparator"
	"calculator/internal/prevalidator"
	"calculator/internal/printer"
	"calculator/internal/validator"
	"fmt"
)

type expInput interface {
	ReadInput(result int) (string, error)
}

type expPrevalidator interface {
	Prevalidate(result int, expression string) (string, error)
}

type expValidator interface {
	Validate(expression string) (bool, error)
}

type expPreparator interface {
	Preparate(expression string) ([]string, error)
}

type expParser interface {
	Parse(expressionsArr []string) (model.Expression, error)
}

type expComputer interface {
	Calculate(expressions model.Expression) (int, error)
}

type expPrinter interface {
	Print(operations []*model.Expression)
}

type Calculator struct {
	input        expInput
	prevalidator expPrevalidator
	validator    expValidator
	preparator   expPreparator
	parser       expParser
	computer     expComputer
	printer      expPrinter
}

func NewCalculator() *Calculator {
	c := new(Calculator)
	c.printer = printer.NewPrinter()
	c.input = input.NewInput()
	c.prevalidator = prevalidator.NewPrevalidator()
	c.validator = validator.NewValidator()
	c.preparator = preparator.NewPreparator(
		[]rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'},
		[]rune{'+', '-', '*', '/'},
		[]rune{'('},
		[]rune{')'},
	)

	var operatorsArr []model.Operator
	var priorityPlusMinus = 1
	var priorityMultiplyDivision = 2
	operatorsArr = append(operatorsArr,
		*model.NewOperator('+', priorityPlusMinus, model.Sum),
		*model.NewOperator('-', priorityPlusMinus, model.Minus),
		*model.NewOperator('*', priorityMultiplyDivision, model.Multiply),
		*model.NewOperator('/', priorityMultiplyDivision, model.Division),
	)

	var groupArr []model.Group
	var priorityGroupOpens = 3
	var priorityGroupEnds = -3
	groupArr = append(groupArr,
		*model.NewGroup('(', priorityGroupOpens),
		*model.NewGroup(')', priorityGroupEnds),
	)

	c.parser = parser.NewParser(operatorsArr, groupArr)
	c.computer = computer.NewComputer()

	return c
}

func (c *Calculator) Run(defaulResult int) {
	var operationsArr []string
	var result int

	for {
		inputString, err := c.input.ReadInput(result)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			result = 0
			continue
		}

		inputString, err = c.prevalidator.Prevalidate(result, inputString)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			result = 0
			continue
		}

		check, err := c.validator.Validate(inputString)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			result = 0
			continue
		}

		if check {
			operationsArr, err = c.preparator.Preparate(inputString)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				result = 0
				continue
			}
		}

		expression, err := c.parser.Parse(operationsArr)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			result = 0
			continue
		}
		
		result, err = c.computer.Calculate(expression)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			result = 0
			continue
		} else {
			fmt.Printf("Result: %d\n", result)

		}
	}
}
