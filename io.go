package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ProjectIO interface {
	say(string)
	ask(string) string
}

type ConsoleIO struct{}

func (io ConsoleIO) say(txt string) {
	println(txt)
}

func (io ConsoleIO) ask(prompt string) string {
	r := bufio.NewReader(os.Stdin)

	fmt.Print(prompt + " ")
	response, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.TrimSpace(response)
}
