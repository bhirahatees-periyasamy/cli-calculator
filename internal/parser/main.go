package parser

import (
	"fmt"
	"strings"
)

type Expression struct {
	left      string
	operation Operator
	right     string
}

func (exp *Expression) Parse(input string) {
	parts := strings.Fields(input)

	partSlice = new([]string)

	if len(partSlice) != 3 {
		fmt.Println("Invalid expression. Use format: 2 + 3")
		return
	}

	exp.left = parts[0]
	exp.operation = parts[1]
	exp.right = parts[2]

	fmt.Println("Left:", exp.left)
	fmt.Println("Operator:", exp.operation)
	fmt.Println("Right:", exp.right)
}
