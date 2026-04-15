package evaluator

import (
	"fmt"
	"strconv"

	"github.com/bhirahatees-periyasamy/internal/parser"
)

func Evaluate(exp *parser.Expression) (float64, error) {
	left, err := strconv.ParseFloat(exp.Left, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid left operand %q: %w", exp.Left, err)
	}

	right, err := strconv.ParseFloat(exp.Right, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid right operand %q: %w", exp.Right, err)
	}

	switch exp.Operation {
	case parser.Add:
		return left + right, nil
	case parser.Subract:
		return left - right, nil
	case parser.Multiply:
		return left * right, nil
	case parser.Divide:
		if right == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("unknown operator")
	}
}
