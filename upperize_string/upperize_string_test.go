package upperize_string
import "testing"

func TestUpperizeString (t *testing.T) {
	expected := "I AM LEARNING GO."
	got := UpperizeString("I am learning Go.")

	if expected != got {
		t.Errorf("Expected %s, but got %s", expected, got)
	}
}
