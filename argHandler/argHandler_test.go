package arghandler

import "testing"

func TestArgHandler(t *testing.T) {
	t.Run("expected argument", func(t *testing.T) {
		got := ArgHandler("-c")
		expected := "checking size of the file"
		if got != expected {
			t.Errorf("\ngot 	 %v\nwant	 %v", got, expected)
		} else {
			t.Logf("\ngot 	 %v\nwant	 %v", got, expected)
		}
	})
	t.Run("empty argument", func(t *testing.T) {
		got := ArgHandler("")
		expected := "invalid argument"
		if got != expected {
			t.Errorf("\ngot 	 %v\nwant	 %v", got, expected)
		} else {
			t.Logf("\ngot 	 %v\nwant	 %v", got, expected)
		}
	})
	t.Run("wrong argument", func(t *testing.T) {
		got := ArgHandler("-d")
		expected := "invalid argument"
		if got != expected {
			t.Errorf("\ngot 	 %v\nwant	 %v", got, expected)
		} else {
			t.Logf("\ngot 	 %v\nwant	 %v", got, expected)
		}

	})
}
