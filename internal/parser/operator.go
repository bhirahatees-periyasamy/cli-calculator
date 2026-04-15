package parser

import "fmt"

type Operator int

const (
	Add Operator = iota
	Subract
	Multiply
	Divide
)

func ParseOperator(s string) (Operator, error) {
	switch s {
	case "+":
		return Add, nil
	case "-":
		return Subract, nil
	case "*":
		return Multiply, nil
	case "/":
		return Divide, nil
	default:
		return 0, fmt.Errorf("unknown operator: %q", s)
	}
}
