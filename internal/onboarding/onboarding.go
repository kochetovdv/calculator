// Package prints different information about application using
package onboarding

import "fmt"

func Usage() {
	fmt.Println("How to use calculator application")
	fmt.Println("Expression should be entered like a line and should be finished with Enter.")
	fmt.Println("Result is an integer.")
	fmt.Println("It is allowed:")
	fmt.Println("\tto use digits")
	fmt.Println("\tto use standart mathematics operations: +, -, *, /")
	fmt.Println("\tto group expressions with the brackets: (,  )")
	fmt.Println("\tto use empty brackets (). It is similar with zero value")
	fmt.Println("\tto use spaces and tabulation")
	fmt.Println("\tto start expression with operation. It is similar to use previous result like a first operand")

	fmt.Println("Next actions will terminied with zero result:")
	fmt.Println("\tusing another symbols")
	fmt.Println("\tdublicatint mathematical operations")
	fmt.Println("\tusing brackets )( without operator between them")
	fmt.Println("\tfinishing expression with operator")
	fmt.Print("\n")
}
