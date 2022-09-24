package common

import (
	"bufio"
	"os"
	"strings"
)

func getInput() string {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
