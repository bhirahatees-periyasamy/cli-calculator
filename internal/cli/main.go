package cli

import (
	"bufio"
	"fmt"
	"os"
)

func GetUserInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}
