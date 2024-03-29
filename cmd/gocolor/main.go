package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/gocolor"
)

func main() {
	var (
		cli  = cmdline.NewBasicParser()
		expr = cli.NamedArg("COLOREXP...").Strings("")
		pal  = cli.Option("-c, --color-palette", "shows color names").Bool()
	)
	usage := cli.Usage()
	usage.Example("custom color by regexp",
		`$ gocolor "error.*:red" "info.*:green"`,
	)
	cli.Parse()

	if pal {
		os.Stdout.Write(colorPalette())
		os.Exit(0)
	}

	// look for any .gocolor file here or in home directory
	filename := filepath.Join(os.Getenv("HOME"), ".gocolor")
	home := loadExpr(filename)
	local := loadExpr(".gocolor")
	all := append(expr, local...)
	all = append(all, home...)
	custom := gocolor.NewCustom(all...)
	err := gocolor.Colorize(os.Stdout, os.Stdin, custom)
	if err != nil {
		os.Exit(1)
	}
}

func loadExpr(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return strings.Split(string(data), "\n")
}

func colorPalette() []byte {
	var buf bytes.Buffer
	colors := []string{
		"black",
		"red",
		"green",
		"yellow",
		"blue",
		"magenta",
		"cyan",
		"white",
		"bgblack",
		"bgred",
		"bggreen",
		"bgyellow",
		"bgblue",
		"bgmagenta",
		"bgcyan",
		"bgwhite",
	}

	attributes := []string{
		"",
		"dim",
		"bright",
		"blink",
		"underscore",
		"reverse",
	}

	for _, color := range colors {
		for _, attr := range attributes {
			switch attr {
			case "":
				custom := gocolor.NewCustom(".*:" + color)
				custom.Colorize(&buf, []byte(color))
				fmt.Fprintf(&buf, " %s\n", color)
			default:
				expr := ".*:" + color + ";" + attr
				custom := gocolor.NewCustom(expr)
				custom.Colorize(&buf, []byte(color))
				fmt.Fprintf(&buf, " %s;%s\n", color, attr)
			}
		}

	}
	return buf.Bytes()
}
