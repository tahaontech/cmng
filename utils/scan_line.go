package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ScanLn(text string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(text)
	name, _ := reader.ReadString('\n') // Reads until a newline character
	return name
}
