package count_string

import "testing"

func TestCountString(t *testing.T) {
	got := CountString("I really think I am getting the hang of Go now.")
	expected := 47

	if got != expected {
		t.Errorf("got %d, wanted %d", got, expected)
	}
}