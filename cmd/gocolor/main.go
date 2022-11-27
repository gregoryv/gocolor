package main

import (
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/gocolor"
)

func main() {
	var (
		cli    = cmdline.NewBasicParser()
		custom = cli.Option("-c, --custom").String("")
	)
	usage := cli.Usage()
	usage.Example("custom color by regexp",
		`$ gocolor -c "error.*:red info.*:green"`,
		"",
		"custom value is a space delimited list of REGEXP:COLOR",
		"",
	)
	cli.Parse()
	err := gocolor.Colorize(os.Stdout, os.Stdin, custom)
	if err != nil {
		os.Exit(1)
	}
}
