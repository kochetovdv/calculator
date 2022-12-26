package input

import (
	"bufio"
	"fmt"
	"os"
)

func InputString() (string, error) {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed due user input %w", err)
	}
	return str, nil
}
