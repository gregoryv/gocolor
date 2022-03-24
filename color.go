package gocolor

import (
	"bufio"
	"bytes"
	"errors"
	"io"

	"github.com/gregoryv/vt100"
)

// Colorize go test output and returns ErrTestFailed if a test failure
// is found.
func Colorize(w io.Writer, r io.Reader) error {
	s := bufio.NewScanner(r)
	var err error
	for s.Scan() {
		line := s.Bytes()
		switch {

		case bytes.Contains(line, []byte("=== RUN")):
			writeColored(w, yellow, []byte("=== RUN"), line)

		case bytes.Contains(line, []byte("--- FAIL:")):
			writeColored(w, red, []byte("--- FAIL:"), line)
			err = ErrTestFailed

		case bytes.Contains(line, []byte("--- SKIP:")):
			writeColored(w, cyan, []byte("--- SKIP:"), line)

		case bytes.Contains(line, []byte("--- PASS:")):
			writeColored(w, green, []byte("--- PASS:"), line)

		case bytes.Contains(line, []byte("PASS")):
			writeColored(w, green, []byte("PASS"), line)

		case bytes.Contains(line, []byte("FAIL")):
			writeColored(w, red, []byte("FAIL"), line)
			err = ErrTestFailed

		default:
			w.Write([]byte(line))
		}
		w.Write(newLine)
	}
	return err
}

func writeColored(w io.Writer, color []byte, prefix, line []byte) {
	l := []byte(line)
	w.Write(color)
	i := bytes.Index(l, []byte(prefix)) + len(prefix)
	w.Write(l[:i])
	w.Write(reset)
	w.Write(l[i:])
}

var (
	ErrTestFailed = errors.New("failed")

	fg = vt100.ForegroundColors()

	yellow = fg.Yellow.Bytes()
	red    = fg.Red.Bytes()
	green  = fg.Green.Bytes()
	cyan   = fg.Cyan.Bytes()

	vt    = vt100.Attributes()
	reset = vt.Reset.Bytes()

	newLine = []byte("\n")
)
