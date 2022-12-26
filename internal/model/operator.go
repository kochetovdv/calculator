package model

import (
	"calculator/pkg/mth"
	"fmt"
)

type OperatorList []*Operator

func (o OperatorList) Len() int {
	return len(o)
}

func (o OperatorList) Less(i, j int) bool {
	return o[i].Priority > o[j].Priority
}

func (o OperatorList) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

type Operator struct {
	ID       *ID
	Symbol   rune
	Priority int
	OpFunc   func(n1, n2 Number) *Result
	Result   Result
}

func NewOperator(r rune, priority int, opFunc func(n1, n2 Number) *Result) *Operator {
	o := new(Operator)
	o.Symbol = r
	o.Priority = priority
	o.OpFunc = opFunc
	return o
}

// TODO Can we make functions with (o *Operator) or not?
// TODO If we can`t, possible we should to replace this methods somewhere from Operator

func Sum(n1, n2 Number) *Result {
	result := new(Result)
	if n1.isOk() && n2.isOk() {
		result.Value, result.Error = mth.Sum(n1.Result.Value, n2.Result.Value)
	} else {
		result.Value = 0
		result.Error = fmt.Errorf("error during operation Sum")
	}
	return result
}

func Minus(n1, n2 Number) *Result {
	result := new(Result)
	if n1.isOk() && n2.isOk() {
		result.Value, result.Error = mth.Minus(n1.Result.Value, n2.Result.Value)
	} else {
		result.Value = 0
		result.Error = fmt.Errorf("error during operation Minus")
	}
	return result
}

func Multiply(n1, n2 Number) *Result {
	result := new(Result)
	if n1.isOk() && n2.isOk() {
		result.Value, result.Error = mth.Multiply(n1.Result.Value, n2.Result.Value)
	} else {
		result.Value = 0
		result.Error = fmt.Errorf("error during operation Multiply")
	}
	return result
}

func Division(n1, n2 Number) *Result {
	result := new(Result)
	if n1.isOk() && n2.isOk() {
		result.Value, result.Error = mth.Division(n1.Result.Value, n2.Result.Value)
	} else {
		result.Value = 0
		result.Error = fmt.Errorf("error during operation Division: ")
	}
	return result
}
