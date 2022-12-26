package model

type Result struct {
	Value int
	Error error
}

func NewResult() *Result {
	r := new(Result)
	return r
}
