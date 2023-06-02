package gocolor

import (
	"fmt"
	"io"
	"regexp"
)

func NewGodocColors(dst io.Writer) *GodocColors {
	return &GodocColors{dst}
}

type GodocColors struct {
	dst io.Writer
}

// Write colorizes the line, assuming it's an output from go doc.
func (g *GodocColors) Write(p []byte) (int, error) {
	v := highlightGoDoc(string(p))
	return g.dst.Write([]byte(v))
}

var (
	kwrepl = fmt.Sprintf("$1%s$2%s$3", fg.Magenta.String(), reset)
	tyrepl = fmt.Sprintf("$1%s$2%s$3", fg.Cyan.String(), reset)

	types       = regexp.MustCompile(`(\W)((?i:\w*\.)?\w+)(\)|\n|,)`)
	alias       = regexp.MustCompile(`(\s)(\w+\.\w+)( \.\.\.)`)
	docKeywords = regexp.MustCompile(`(\W?)(^package|^\s\s\s\sfunc|^func|^var|^const|^type|struct)(\{?\W)`)
)

// highlightGoDoc output
func highlightGoDoc(v string) string {

	if !docKeywords.MatchString(v) {
		return v
	}
	v = docKeywords.ReplaceAllString(v, kwrepl)
	v = types.ReplaceAllString(v, tyrepl)

	return v
}
