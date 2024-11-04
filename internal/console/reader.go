package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Listen() (string, error) {
	fmt.Print("\033[90m")
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	defer fmt.Print("\033[0m")
	// Remove trailing spaces and return
	return strings.TrimSpace(input), nil
}
