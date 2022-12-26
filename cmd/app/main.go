package main

import (
	"calculator/internal/calculator"
	"calculator/internal/onboarding"
)

const (
	calculateFailedExitCode = 20
	defaulResult            = 0
)

// TODO Configurator with exit codes, allowe symbols
// TODO launch JSON
// TODO arguments in expression for usage, help, close/exit
// TODO add logging
// TODO float instead of int
// TODO accuracy for float with default accuracy value and changing by argument entering
func main() {
	onboarding.Usage()

	calc := calculator.NewCalculator()
	calc.Run(defaulResult)
}
