package mth // mathematics
import "fmt"

func Sum(a, b int) (int, error) {
	return a + b, nil
}

func Minus(a, b int) (int, error) {
	return a - b, nil
}

func Multiply(a, b int) (int, error) {
	return a * b, nil
}

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
