package validator

// TODO Make logging
// TODO Possible all validations make with one regexp or with another way

import (
	standart "calculator/pkg/errorchecking"
	"fmt"
)

/* TODO Make typical structure for validation like this:

type Validation struct {
	ValidationName    string
	RegularExpression string
	Error             string
}

func NewValidation(v, r, e string) *Validation {
	validat := new(Validation)
	validat.ValidationName = v
	validat.RegularExpression = r
	validat.Error = e
	return validat
}

func NewValidations(v ...Validation) []Validation {
	var val []Validation
	val = append(val, v...)
	return val
}*/

type Validator struct {
}

func NewValidator() *Validator {
	v := new(Validator)
	return v
}

func (v *Validator) Validate(expression string) (bool, error) {
	var err error

	// Checking for denied (now allowed) symbols
	err = DeniedSymbolsValidation(expression)
	if err != nil {
		return false, err
	}

	// Checking for dublicated operators, for example ++ or */
	err = DoubleOperatorsValidation(expression)
	if err != nil {
		return false, err
	}

	// Checking for denied combinations of operators and brackets, for example +) or (+
	err = OperatorWithGroupCombinationValidation(expression)
	if err != nil {
		return false, err
	}

	// Checking for operators in the end.
	// TODO Logic here is that we don`t know what is after operator. If operator is at the start of expression, we have previous result or zero value
	// But in the end there`s possible error be occuried: for example 1+1/ -> 1+1/0
	// Maybe we should cut the last operator. Then 1+1/ -> 1+1
	err = OperatorWithoutOperandsAtEndtValidation(expression)
	if err != nil {
		return false, err
	}

	// Checking for brackets close to digits, for example )42 or 42(
	err = ValueWithGroupCombinationValidation(expression)
	if err != nil {
		return false, err
	}

	// Checking for every open bracket has close bracket, count of ( is equal to count of )
	err = GroupCountValidation(expression)
	if err != nil {
		return false, err
	}

	// If the expression has no two digits and one operator then error occuried
	// TODO make that digit as a result of expression
	err = TwoNumbersAndOneOperatorValidation(expression)
	if err != nil {
		return false, err
	}

	return true, nil
}

// TODO Make typical validation. Fix TwoNumbersAndOneOperatorValidation and GroupCountValidation
// Checking for denied symbols
func DeniedSymbolsValidation(s string) error {
	// Pattern is formed with (^) - denies all non-allowed symbols
	pattern := `[^\d\s\/\+\-\*\(\)]+`
	check := standart.RegexpPatternCheck(pattern, s)
	if check {
		return fmt.Errorf("expression have denied symbols")
	}
	return nil
}

// Checking for dublicated operators, for example ++ or */
func DoubleOperatorsValidation(s string) error {
	// Pattern is formed by searching two or more operators close to each other
	pattern := `[\/\+\-\*]{2,}`
	check := standart.RegexpPatternCheck(pattern, s)
	if check {
		return fmt.Errorf("dublicating operators was found")
	}
	return nil
}

// Checking for denied combinations of operators and brackets, for example +) or (+
func OperatorWithGroupCombinationValidation(s string) error {
	// Pattern is formed by searching operator before ) or after (
	pattern := `[\/\+\-\*]\)|\([\/\+\-\*]`
	check := standart.RegexpPatternCheck(pattern, s)
	if check {
		return fmt.Errorf("denied combination of brackets and operators")
	}
	return nil
}

// Checking for operators in the end.
func OperatorWithoutOperandsAtEndtValidation(s string) error {
	// Pattern is formed by searching operator at the end of the expression
	pattern := `[\d\)]*[\/\+\-\*]$`
	check := standart.RegexpPatternCheck(pattern, s)
	if check {
		return fmt.Errorf("expression ends with operator")
	}
	return nil
}

// Checking for brackets close to digits, for example )42 or 42(
func ValueWithGroupCombinationValidation(s string) error {
	// Pattern is formed by searching digits after close bracket ) or before open bracket (
	pattern := `\d\(|\)\d`
	check := standart.RegexpPatternCheck(pattern, s)
	if check {
		return fmt.Errorf("combination of digits and bracket without operator between them")
	}
	return nil
}

// Checking for every open bracket has close bracket, count of ( is equal to count of )
func GroupCountValidation(s string) error {
	bracketsCount := 0 // There`s no brackets at the start of validation
	for _, el := range s {
		if el == '(' {
			bracketsCount++
		} else if el == ')' {
			bracketsCount--
		}
		if bracketsCount < 0 { // if count is less then 0 then error occurred
			break
		}
	}
	if bracketsCount < 0 {
		return fmt.Errorf("expression starts with ) or count of ) is more than count of (")
	} else if bracketsCount > 0 {
		return fmt.Errorf("count of ( is not equal with count of )")
	}
	return nil
}

// Checking for )( combination
func GroupCloseWithoutDataBetweenValidation(s string) error {
	// Pattern is formed by searching )( combination
	pattern := `\)\(` // there was check for () and )(, but after a while decided to replace ()->0. Previous pattern:=`\(\)|\)\(`
	check := standart.RegexpPatternCheck(pattern, s)
	if check {
		return fmt.Errorf("denied combination )(")
	}
	return nil
}

// If the expression has no two digits and one operator then error occuried
// TODO make that digit as a result of expression
func TwoNumbersAndOneOperatorValidation(s string) error {
	pattern := `\d(.*)[\/\+\-\*](.*)\d`
	check := standart.RegexpPatternCheck(pattern, s)
	if !check { // WANING! It is different with another methods where check==true
		return fmt.Errorf("expression has no two digits and one operator at least")
	}
	return nil
}
