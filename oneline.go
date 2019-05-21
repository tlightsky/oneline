package main

import (
	"fmt"
	"strings"
	"flag"

	"github.com/atotto/clipboard"
)

var (
	quietFlag bool
)

func init() {
	flag.BoolVar(&quietFlag, "q", false, "the quiet file")
}

func main() {
	flag.Parse()

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
	if !quietFlag {
		//fmt.Println("--------------------Onelined And Copied----------------------")
		fmt.Println(final)
	}
}
