package capitalize_string
import "testing"

func TestCapitalizeString(t *testing.T) {
	expected := "Ruby monstas for the win!"
	got := CapitalizeString("ruby monstas for the win!")

	if expected != got {
		t.Errorf("Expected \"%s\", but got \"%s\"", expected, got)
	}
}
