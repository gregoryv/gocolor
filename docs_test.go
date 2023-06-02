package gocolor

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/gregoryv/golden"
)

func TestGoDocColor(t *testing.T) {
	in, err := os.Open("testdata/godoc.txt")
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	s := bufio.NewScanner(in)
	w := NewGodocColors(&buf)
	for s.Scan() {
		w.Write(s.Bytes())
		w.Write([]byte("\n"))
	}
	golden.Assert(t, buf.String())
}
