package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gregoryv/gocolor"
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: go test . | gocolor
`)
	}
	flag.Parse()
	err := gocolor.Colorize(os.Stdout, os.Stdin)
	if err != nil {
		os.Exit(1)
	}
}
