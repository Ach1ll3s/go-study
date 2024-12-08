package practice1

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestHelloOutput(t *testing.T) {
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	HelloWorld()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	want := "Hello, world!\n"

	if got := string(out); got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
