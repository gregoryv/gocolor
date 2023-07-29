package gocolor

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/gregoryv/vt100"
)

// Colorize go test output and returns ErrTestFailed if a test failure
// is found.
func Colorize(w io.Writer, r io.Reader, custom *Custom) error {
	s := bufio.NewScanner(r)
	var err error
	var firstline sync.Once
	for s.Scan() {
		line := s.Bytes()
		firstline.Do(func() {
			if strings.HasPrefix(string(line), "package ") {
				w = NewGodocColors(w)
			}
		})
		switch {
		case custom.Colorize(w, line):
		case writeColoredWorkingDirPath(w, line):
		case writeColoredPath(w, line):
		case writeColored(w, line, yellow, "=== RUN"):
		case writeColored(w, line, green, "--- PASS:"):
		case writeColored(w, line, green, "PASS"):
		case writeColored(w, line, red, "--- FAIL:"):
		case writeColored(w, line, cyan, "--- SKIP:"):
		case writeColoredCoverage(w, line):
		case writeColored(w, line, red, "FAIL"):
			err = ErrTestFailed

		default:
			w.Write([]byte(line))
		}
		w.Write(newLine)
	}
	return err
}

var pathPattern = `\/(?:[\w.-]+\/)*[\w.-]+\.go`
var pathExpr = regexp.MustCompile(pathPattern)

// only match files in the working directory
func writeColoredWorkingDirPath(w io.Writer, line []byte) bool {
	if !pathExpr.Match(line) {
		return false
	}
	dir, err := os.Getwd()
	if err != nil {
		return false
	}
	if !bytes.Contains(line, []byte(dir)) {
		return false
	}

	expr, _ := ParseExpr(pathPattern + ":cyan") // brighter cyan
	custom := Custom{[]*Expr{expr}}
	return custom.Colorize(w, line)
}

func writeColoredPath(w io.Writer, line []byte) bool {
	if !pathExpr.Match(line) {
		return false
	}
	expr, _ := ParseExpr(pathPattern + ":cyan;dim") // dimmer cyan
	custom := Custom{[]*Expr{expr}}
	return custom.Colorize(w, line)
}

func writeColored(w io.Writer, line []byte, color []byte, prefix string) bool {
	i := bytes.Index(line, []byte(prefix))
	if i == -1 {
		return false
	}
	i += len(prefix)
	w.Write(color)
	w.Write(line[:i])
	w.Write(reset)
	w.Write(line[i:])
	return true
}

func writeColoredCoverage(w io.Writer, line []byte) bool {
	i := bytes.Index(line, []byte("coverage:"))
	if i == -1 {
		return false
	}
	var percent float64
	_, err := fmt.Sscanf(string(line[i:]), "coverage: %f%% of statements", &percent)
	if err != nil {
		return false
	}
	var color []byte
	switch {
	case percent < 100.0:
		color = magenta // todo orange
	default:
		color = green
	}

	w.Write(line[:i])
	w.Write([]byte("coverage: "))
	w.Write(color)
	fmt.Fprintf(w, "%v%%", percent)
	w.Write(reset)
	w.Write([]byte(" of statements"))
	return true
}

var (
	ErrTestFailed = errors.New("failed")

	fg      = vt100.ForegroundColors()
	yellow  = fg.Yellow.Bytes()
	red     = fg.Red.Bytes()
	green   = fg.Green.Bytes()
	cyan    = fg.Cyan.Bytes()
	magenta = fg.Magenta.Bytes()

	vt    = vt100.Attributes()
	reset = vt.Reset.Bytes()

	newLine = []byte("\n")
)
