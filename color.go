package gocolor

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Colorize(w io.Writer, r io.Reader) error {
	s := bufio.NewScanner(r)
	painted := GoTestOutput()
	for s.Scan() {
		line := s.Text()
		// paint any values
		for k, v := range painted {
			line = strings.ReplaceAll(line, k, v)
		}
		w.Write([]byte(line))
		w.Write([]byte("\n"))
	}
	return s.Err()
}

func GoTestOutput() map[string]string {
	painted := map[string]string{
		"=== RUN":  YELLOW,
		"--- FAIL": RED,
		"--- PASS": GREEN,
	}
	// color the values first
	for k, v := range painted {
		painted[k] = fmt.Sprintf("%s%s%s", v, k, RESET)
	}
	return painted
}

const (
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	RESET  = "\033[0m"
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
