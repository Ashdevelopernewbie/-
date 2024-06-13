package bytes

import (
	"testing"
)

func TestCheckBytes(t *testing.T) {
	got := CheckBytes("../utils/test.txt")
	expected := 342190
	if got != expected {
		t.Errorf("\ngot 	 %v\nwant	 %v", got, expected)
	} else {
		t.Logf("\ngot 	 %v\nwant	 %v", got, expected)
	}
}
