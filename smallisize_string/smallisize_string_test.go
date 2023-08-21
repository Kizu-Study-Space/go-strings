package smallisize_string
import "testing"

func TestSmallisizeString(t *testing.T) {
	expected := "let's code go!"
	got := SmallisizeString("Let's code GO!")

	if expected != got {
		t.Errorf("Expected: \"%s\", but got: \"%s\"", expected, got)
	}
}
