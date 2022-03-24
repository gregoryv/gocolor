package gocolor

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"testing"
)

func TestColorize(t *testing.T) {
	var buf bytes.Buffer

	r := strings.NewReader(testOutput)
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

func BenchmarkColorize(b *testing.B) {
	r := strings.NewReader(testOutput)
	for i := 0; i < b.N; i++ {
		Colorize(ioutil.Discard, r)
		r.Reset(testOutput)
	}
}

const testOutput = `=== RUN   TestSomething
--- PASS:   TestSomething
    --- PASS:
--- FAIL:
--- SKIP:

a
b
PASS
FAIL

`
