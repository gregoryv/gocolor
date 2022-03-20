package gocolor

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// Colorize paint go test output and returns ErrTestFailed if a test
// failure is found.
func Colorize(w io.Writer, r io.Reader) error {
	s := bufio.NewScanner(r)
	var err error
	for s.Scan() {
		line := s.Text()
		var color string
		var prefix string
		switch {

		case strings.HasPrefix(line, "=== RUN"):
			color = YELLOW
			prefix = "=== RUN"

		case strings.HasPrefix(line, "--- FAIL"):
			color = RED
			prefix = "--- FAIL"
			err = ErrTestFailed

		case strings.HasPrefix(line, "--- PASS"):
			color = GREEN
			prefix = "--- PASS"

		case line == "PASS":
			color = GREEN
			prefix = "PASS"

		case line == "FAIL":
			color = RED
			prefix = "FAIL"
			err = ErrTestFailed
		}
		// paint any values
		if color != "" {
			w.Write([]byte(color))
			w.Write([]byte(prefix))
			w.Write([]byte(RESET))
			w.Write([]byte(line[len(prefix):]))
		} else {
			w.Write([]byte(line))
		}
		w.Write([]byte("\n"))
	}
	return err
}

var ErrTestFailed = errors.New("failed")

const (
	RED      = "\033[31m"
	GREEN    = "\033[32m"
	YELLOW   = "\033[33m"
	WHITE    = "\033[37m"
	BG_GREEN = "\033[42m"
	RESET    = "\033[0m"
)

type VTCode uint

/*https://www2.ccs.neu.edu/research/gpc/VonaUtils/vona/terminal/vtansi.htm#colors
0	Reset all attributes
1	Bright
2	Dim
4	Underscore
5	Blink
7	Reverse
8	Hidden

	Foreground Colours
30	Black
31	Red
32	Green
33	Yellow
34	Blue
35	Magenta
36	Cyan
37	White

	Background Colours
40	Black
41	Red
42	Green
43	Yellow
44	Blue
45	Magenta
46	Cyan
47	White
*/
