package input

import (
	"calculator/pkg/input"

	"fmt"
)

type Input struct {
	UserInput string
	Error     error
}

func NewInput() *Input {
	i := new(Input)
	return i
}

func (i *Input) ReadInput(result int) (string, error) {
	//TODO Arguments for usege, help, close/exit
	fmt.Printf("%d >>> ", result)

	// User`s expression input
	i.UserInput, i.Error = input.InputString()
	if i.Error != nil {
		return "", fmt.Errorf("failed due input %w", i.Error)
	}

	// Empty input
	if i.UserInput == "" {
		return "", fmt.Errorf("your input is empty")
	}

	return i.UserInput, nil
}
