package main

import (
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/gocolor"
)

func main() {
	var (
		cli  = cmdline.NewBasicParser()
		expr = cli.NamedArg("COLOREXP...").Strings("")
	)
	usage := cli.Usage()
	usage.Example("custom color by regexp",
		`$ gocolor "error.*:red" "info.*:green"`,
		"",
		"custom value is a space delimited list of REGEXP:COLOR",
		"",
	)
	cli.Parse()
	custom := gocolor.NewCustom(expr...)
	err := gocolor.Colorize(os.Stdout, os.Stdin, custom)
	if err != nil {
		os.Exit(1)
	}
}
