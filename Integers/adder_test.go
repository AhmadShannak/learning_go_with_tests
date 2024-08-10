package integers

import "testing"

func TestAdd(t *testing.T) {
	got := Add(3, 2)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
