package letter_in_string
import "testing"

func TestLetterInString(t *testing.T){
	expected := true
	got := LetterInString("Come on, let’s all learn Go!", "m")

	if expected != got {
		t.Errorf("got %t, wanted %t", got, expected)
	}
}