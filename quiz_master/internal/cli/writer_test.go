package cli

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// Testing Stdout inspiration and insight from: https://stackoverflow.com/a/47281683 (EnterSB)

func TestDisplayVersion(t *testing.T) {
	// Prepare the reader and writer from os pipe
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Actual func that print output
	DisplayVersion()

	// Close writer
	w.Close()
	// Read []byte data from reader
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if strings.TrimSpace(string(out)) != "Quiz Master version 1.0.0. By: Jidni Ilman" {
		t.Errorf("Expected %s, got %sa", "Quiz Master version 1.0.0. By: Jidni Ilman", out)
	}
}
