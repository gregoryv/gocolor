package gocolor

import (
	"bufio"
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
			writeColored(w, fg.Yellow, "=== RUN", line)

		case strings.Contains(line, "--- FAIL:"):
			writeColored(w, fg.Red, "--- FAIL:", line)
			err = ErrTestFailed

		case strings.Contains(line, "--- SKIP:"):
			writeColored(w, fg.Cyan, "--- SKIP:", line)

		case strings.Contains(line, "--- PASS:"):
			writeColored(w, fg.Green, "--- PASS:", line)

		case line == "PASS":
			writeColored(w, fg.Green, "PASS", line)

		case strings.HasPrefix(line, "FAIL"):
			writeColored(w, fg.Red, "FAIL", line)
			err = ErrTestFailed

		default:
			w.Write([]byte(line))
		}
		w.Write(newLine)
	}
	return err
}

func writeColored(w io.Writer, color vt100.Code, prefix, line string) {
	w.Write(color.Bytes())
	i := strings.Index(line, prefix) + len(prefix)
	w.Write([]byte(line[:i]))
	w.Write(vt.Reset.Bytes())
	w.Write([]byte(line[i:]))
}

var (
	ErrTestFailed = errors.New("failed")

	fg = vt100.ForegroundColors()
	bg = vt100.BackgroundColors()
	vt = vt100.Attributes()

	newLine = []byte("\n")
)
