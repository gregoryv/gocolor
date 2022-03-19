package gocolor

import (
	"bytes"
	"strings"
	"testing"
)

func TestColorize(t *testing.T) {
	var buf bytes.Buffer
	input := `=== RUN   TestSomething
--- PASS   TestSomething
--- FAIL
a
b
PASS
`
	r := strings.NewReader(input)
	if err := Colorize(&buf, r); err != nil {
		t.Log(buf.String())
		t.Error(err)
	}
}
