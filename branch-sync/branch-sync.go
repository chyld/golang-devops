package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	re := regexp.MustCompile("(ahead) ([0-9]+)|(behind) ([0-9]+)")
	matches := re.FindAllStringSubmatch(text, -1)
	var output string

	if len(matches) >= 1 {
		if matches[0][1] == "ahead" {
			output += "+" + matches[0][2]
		}
		if matches[0][3] == "behind" {
			output += "-" + matches[0][4]
		}
	}

	if len(matches) == 2 {
		output += "/-" + matches[1][4]
	}

	if len(output) > 0 {
		output += " "
	}

	fmt.Printf("%s", output)
}
