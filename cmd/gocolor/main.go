package main

import (
	"os"

	"github.com/gregoryv/gocolor"
)

func main() {
	err := gocolor.Colorize(os.Stdout, os.Stdin)
	if err != nil {
		os.Exit(1)
	}
}
