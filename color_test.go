package gocolor

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestColorize(t *testing.T) {
	var buf bytes.Buffer
	input := `=== RUN   TestSomething
--- PASS:   TestSomething
--- FAIL:
--- SKIP:
a
b
PASS
FAIL

`
	r := strings.NewReader(input)
	err := Colorize(&buf, r)
	t.Log("painting input")
	if !errors.Is(ErrTestFailed, err) {
		t.Log(buf.String())
		t.Error("expect error if contains a failure")
	}
}

func TestNothing(t *testing.T) {
	t.SkipNow()
}

func TestFail(t *testing.T) {
	//t.Fail()
}
