package main

import (
	"fmt"
	"strings"

	"github.com/bhirahatees-periyasamy/internal/cli"
	"github.com/bhirahatees-periyasamy/internal/evaluator"
	"github.com/bhirahatees-periyasamy/internal/parser"
)

func main() {
	fmt.Println("Welcome to CLI Calculator App!")
	for {
		input := strings.TrimSpace(cli.GetUserInput(">>"))
		if input == "exit" || input == "quit" {
			fmt.Println("Bye!")
			return
		}

		exp := &parser.Expression{}
		exp.Parse(input)

		result, err := evaluator.Evaluate(exp)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(result)
	}
}
