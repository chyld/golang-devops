package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "regexp"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  text = strings.TrimSpace(text)

  re := regexp.MustCompile("(ahead) ([0-9]+)|(behind) ([0-9]+)")
  matches := re.FindAllStringSubmatch(text, -1)
  output := ""

  if len(matches) >= 1{
    if matches[0][1] == "ahead" {
      output += "+" + matches[0][2]
    }
    if matches[0][3] == "behind" {
      output += "-" + matches[0][4]
    }
  }
  if len(matches) >= 2{
    if matches[1][3] == "behind" {
      output += "/-" + matches[1][4]
    }
  }
  if len(matches) >= 1{
    output += " "
  }
  fmt.Printf("%s", output)
}

