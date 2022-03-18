package main

import (
	"os"

	"github.com/gregoryv/gocolor"
)

func main() {
	gocolor.Colorize(os.Stdout, os.Stdin)
}
