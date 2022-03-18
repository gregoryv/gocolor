package gocolor

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Colorize(w io.Writer, r io.Reader) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		switch {
		case strings.HasPrefix(line, "=== RUN"):
			fmt.Fprintf(w, "%s%s%s%s\n", YELLOW, "=== RUN", RESET, line[8:])

		case strings.HasPrefix(line, "--- FAIL"):
			fmt.Fprintf(w, "%s%s%s%s\n", RED, "--- FAIL", RESET, line[8:])

		case strings.HasPrefix(line, "=== PASS"):
			fmt.Fprintf(w, "%s%s%s%s\n", GREEN, "=== PASS", RESET, line[8:])

		default:
			w.Write([]byte(line))
		}
	}
	return s.Err()
}

const (
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	RESET  = "\033[0m"
)

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
