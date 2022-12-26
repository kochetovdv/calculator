package model

import (
	"fmt"
	"strconv"
)

type Number struct {
	ID          *ID
	StringValue string
	Result      *Result
}

func NewNumber(stringValue string) *Number {
	n := new(Number)
	n.StringValue = stringValue
	n.Result = NewResult()
	return n
}

func (n *Number) Convert() (bool, error) {
	n.Result.Value, n.Result.Error = strconv.Atoi(n.StringValue)
	if n.Result.Error != nil {
		return false, fmt.Errorf("failed converting in integer due to error %w", n.Result.Error)
	}
	return true, nil
}

func (n *Number) GetValue() int {
	return n.Result.Value
}

func (n *Number) isOk() bool {
	return n.Result.Error == nil
}
