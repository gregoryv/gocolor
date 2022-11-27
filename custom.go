package gocolor

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/gregoryv/vt100"
)

func NewCustom(expr ...string) *Custom {
	c := &Custom{}
	for _, v := range expr {
		if expr, err := ParseExpr(v); err != nil {
			log.Print(err)
		} else {
			c.exp = append(c.exp, expr)
		}
	}
	return c
}

type Custom struct {
	exp []*Expr
}

func (c *Custom) Colorize(w io.Writer, line []byte) bool {
	var done bool
	for _, exp := range c.exp {
		if !done && !exp.Regexp.Match(line) {
			continue
		}
		done = true
		line = exp.Regexp.ReplaceAll(line, exp.attributes)
	}
	if done {
		w.Write(line)
	}
	return done
}

func ParseExpr(v string) (*Expr, error) {
	if v == "" {
		return nil, fmt.Errorf("Expr: empty")
	}
	parts := strings.Split(v, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Expr: missing ':'")
	}

	re, err := regexp.Compile("(" + parts[0] + ")")
	if err != nil {
		return nil, err
	}
	// create colored replacement
	attr, err := vt100.ParseCodeBytes(parts[1])
	if err != nil {
		return nil, err
	}
	repl := append(attr, []byte("$1")...)
	repl = append(repl, reset...)

	return &Expr{
		Regexp:     re,
		attributes: repl,
	}, nil
}

type Expr struct {
	*regexp.Regexp
	attributes []byte
}
