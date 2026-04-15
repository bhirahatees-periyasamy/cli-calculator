package parser

import (
	"fmt"
	"strings"
)

type Expression struct {
	Left      string
	Operation Operator
	Right     string
}

func (exp *Expression) Parse(input string) {
	parts := strings.Split(input, "")

	partSlice := []string{}

	for _, part := range parts {
		if strings.TrimSpace(part) != "" {
			partSlice = append(partSlice, part)
			fmt.Println(partSlice)
		}
	}

	if len(partSlice) != 3 {
		fmt.Println("Invalid expression. Use format: 2 + 3")
		return
	}

	op, err := ParseOperator(partSlice[1])
	if err != nil {
		fmt.Println("Invalid operator:", err)
		return
	}

	exp.Left = partSlice[0]
	exp.Operation = op
	exp.Right = partSlice[2]
}
