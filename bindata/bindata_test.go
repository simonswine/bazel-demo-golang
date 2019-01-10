package bindata

import (
	"testing"
)

func TestMessage(t *testing.T) {
	data, err := Asset("test")
	if err != nil {
		// Asset was not found.
		t.Errorf("Asset was not embeded: %s", err)
	}
	got := string(data)
	const expected = "Hello World\n"
	if got != expected {
		t.Errorf("Got embedded asset %q\nexpected %q", got, expected)
	}
}
