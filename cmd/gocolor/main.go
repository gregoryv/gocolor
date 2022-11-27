package main

import (
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/gocolor"
)

func main() {
	var (
		cli = cmdline.NewBasicParser()
	)
	cli.Parse()
	err := gocolor.Colorize(os.Stdout, os.Stdin)
	if err != nil {
		os.Exit(1)
	}
}
