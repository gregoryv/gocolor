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
		var color vt100.Code
		var prefix string
		switch {

		case strings.Contains(line, RUN):
			color = fg.Yellow
			prefix = RUN

		case strings.Contains(line, "--- FAIL:"):
			color = fg.Red
			prefix = "--- FAIL:"
			err = ErrTestFailed

		case strings.Contains(line, "--- SKIP:"):
			color = fg.Cyan
			prefix = "--- SKIP:"

		case strings.Contains(line, "--- PASS:"):
			color = fg.Green
			prefix = "--- PASS:"

		case line == "PASS":
			color = fg.Green
			prefix = "PASS"

		case strings.HasPrefix(line, "FAIL"):
			color = fg.Red
			prefix = "FAIL"
			err = ErrTestFailed
		default:
			w.Write([]byte(line))
		}
		// paint any values
		if color >= 30 {
			writeColored(w, color, prefix, line)
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

const (
	RUN = "=== RUN"
)

var (
	ErrTestFailed = errors.New("failed")

	fg = vt100.ForegroundColors()
	bg = vt100.BackgroundColors()
	vt = vt100.Attributes()

	newLine = []byte("\n")
)
