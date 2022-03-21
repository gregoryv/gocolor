package gocolor

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/gregoryv/vt100"
)

var (
	fg = vt100.ForegroundColors()
	bg = vt100.BackgroundColors()
	vt = vt100.Attributes()
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

		case strings.HasPrefix(line, "=== RUN"):
			color = fg.Yellow
			prefix = "=== RUN"

		case strings.HasPrefix(line, "--- FAIL:"):
			color = fg.Red
			prefix = "--- FAIL "
			err = ErrTestFailed

		case strings.HasPrefix(line, "--- SKIP:"):
			color = fg.Cyan
			prefix = "--- SKIP "

		case strings.HasPrefix(line, "--- PASS:"):
			color = fg.Green
			prefix = "--- PASS "

		case line == "PASS":
			color = fg.Green
			prefix = "PASS"

		case strings.HasPrefix(line, "FAIL"):
			color = fg.Red
			prefix = "FAIL"
			err = ErrTestFailed
		}
		// paint any values
		if color >= 30 {
			w.Write(color.Bytes())
			w.Write([]byte(prefix))
			w.Write(vt.Reset.Bytes())
			w.Write([]byte(line[len(prefix):]))
		} else {
			w.Write([]byte(line))
		}
		w.Write([]byte("\n"))
	}
	return err
}

var ErrTestFailed = errors.New("failed")
