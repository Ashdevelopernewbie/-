package bytes

import (
	"testing"
)

func TestCheckBytes(t *testing.T) {
	got := CheckBytes("/utils/test.txt")
	expected := 342190
	if got != expected {
		t.Errorf("got %v, want %v", got, expected)
	}
}

func TestTry(t *testing.T) {
	got := Try("helo")
	expected := "helo"
	if got != expected {
		t.Errorf("got %v, want %v", got, expected)
	}
}
