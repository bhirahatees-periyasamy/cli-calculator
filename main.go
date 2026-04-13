package main

import (
	"fmt"
	"strings"

	"github.com/bhirahatees-periyasamy/internal/cli"
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

		fmt.Println(input)
		parser = &parser.Expression{}

		parser.parser(input)
	}
}
