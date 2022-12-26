package computer

import (
	"calculator/internal/model"
	"sort"
)

type Computer struct {
	Result model.Result
}

func NewComputer() *Computer {
	c := new(Computer)
	return c
}

func (c *Computer) Calculate(expressions model.Expression) (int, error) {
	// for debugging
	// expressions.Print()

	// Sorting operators by priority
	sort.Sort(model.OperatorList(expressions.Operators))
	c.Result = expressions.Calculate()
	if c.Result.Error != nil {
		return 0, c.Result.Error
	}

	return c.Result.Value, nil
}
