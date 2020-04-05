package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreeet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Max")

	got := buffer.String()
	want := "Hello, Max"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
