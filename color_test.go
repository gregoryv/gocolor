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
    --- PASS:
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

func TestX(t *testing.T) {
	t.Run("", func(t *testing.T) {
		t.Run("x", func(t *testing.T) {
			//t.Fail()
		})
		t.Run("", func(t *testing.T) {})
	})
}

func TestNothing(t *testing.T) {
	t.SkipNow()
}
