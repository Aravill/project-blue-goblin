package console

import (
	"bufio"
	"os"
	"strings"
)

func Listen() (string, error) {
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	// Remove trailing spaces and return
	return strings.TrimSpace(input), nil
}
