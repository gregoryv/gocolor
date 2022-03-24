package gocolor

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/gregoryv/vt100"
)

// Colorize go test output and returns ErrTestFailed if a test failure
// is found.
func Colorize(w io.Writer, r io.Reader) error {
	s := bufio.NewScanner(r)
	var err error
	for s.Scan() {
		line := s.Text()

		switch {

		case strings.Contains(line, "=== RUN"):
			writeColored(w, yellow, "=== RUN", line)

		case strings.Contains(line, "--- FAIL:"):
			writeColored(w, red, "--- FAIL:", line)
			err = ErrTestFailed

		case strings.Contains(line, "--- SKIP:"):
			writeColored(w, cyan, "--- SKIP:", line)

		case strings.Contains(line, "--- PASS:"):
			writeColored(w, green, "--- PASS:", line)

		case line == "PASS":
			writeColored(w, green, "PASS", line)

		case strings.HasPrefix(line, "FAIL"):
			writeColored(w, red, "FAIL", line)
			err = ErrTestFailed

		default:
			w.Write([]byte(line))
		}
		w.Write(newLine)
	}
	return err
}

func writeColored(w io.Writer, color []byte, prefix, line string) {
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
