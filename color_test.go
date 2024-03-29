package gocolor

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestColorize_custom(t *testing.T) {
	var buf bytes.Buffer
	const output = `ttsub out CONNECT ---- -------- MQTT5 ttsub 0s 20 bytes
b something`
	r := strings.NewReader(output)
	_ = Colorize(&buf, r, NewCustom("ttsub:cyan", "bytes:green"))
	t.Log(buf.String())

}

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
	err := Colorize(&buf, r, NewCustom(""))
	if !errors.Is(ErrTestFailed, err) {
		t.Log(buf.String())
		t.Error("expect error if contains a failure")
	}
}

func TestX(t *testing.T) {
	t.Run("", func(t *testing.T) {
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
		Colorize(ioutil.Discard, r, NewCustom(""))
		r.Reset(data)
	}
}

func TestAFailure(t *testing.T) {
	//t.Fail()
}
