package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findian(stringToCheck string) string {
	stringToCheck = strings.ToLower(stringToCheck)
	stringToCheck = strings.TrimSpace(stringToCheck)
	switch {
	case strings.HasPrefix(stringToCheck, "i") && strings.Contains(stringToCheck, "a") && strings.HasSuffix(stringToCheck, "n"):
		return "Found!"
	default:
		return "Not Found!"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')
	fmt.Println(findian(inputString))
}
