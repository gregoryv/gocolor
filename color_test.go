package gocolor

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestColorize(t *testing.T) {
	var buf bytes.Buffer
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
	data, err := os.ReadFile("testdata/test.output")
	if err != nil {
		b.Fatal(err)
	}
	r := bytes.NewReader(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Colorize(ioutil.Discard, r)
		r.Reset(data)
	}
}

func xTestAFailure(t *testing.T) {
	t.Fail()
}
