package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(text, "\n")
	newlines := make([]string, 0)
	for _, line := range lines {
		newlines = append(newlines, strings.TrimSpace(line))
	}
	final := strings.Join(newlines, " ")
	clipboard.WriteAll(final)
	fmt.Println("--------------------Onelined And Copied----------------------")
	fmt.Println(final)
}
