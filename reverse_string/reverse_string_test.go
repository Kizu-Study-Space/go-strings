package reverse_string

import "testing"

func TestReverseString (t *testing.T) {
	expected := ".sesicrexe eseht ekil I"
	got := ReverseString("I like these exercises.")

	if got != expected {
		t.Errorf("got %s, wanted %s", got, expected)
	}
}