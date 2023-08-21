package capitalize_string
import "testing"

func TestCapitalizeString (t *testing.T) {
	expected := "I AM LEARNING GO."
	got := CapitalizeString("I am learning Go.")

	if expected != got {
		t.Errorf("Expected %s, but got %s", expected, got)
	}
}
