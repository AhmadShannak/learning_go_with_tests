package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Known Word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		AssertString(t, got, want)
	})

	t.Run("Not Know Word", func(t *testing.T) {
		d := Dictionary{"test": "this is just a test"}
		_, er := d.Search("meow")
		want := NotFoundError
		if er == None {
			t.Fatal("expected to get error")
		}
		AssertError(t, er, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Adding New", func(t *testing.T) {
		d := Dictionary{"test": "this is just a test"}
		d.Add("test2", "this is also another test")

		value, _ := d.Search("test2")
		want := "this is also another test"
		AssertString(t, value, want)
	})

	t.Run("Adding what already exist", func(t *testing.T) {
		d := Dictionary{"test": "this is just a test"}
		er := d.Add("test", "this is just a test")

		want := AlreadyExistError
		AssertError(t, er, want)
	})
}

func TestUpdate(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}
	d.Update("test", "no longer a normal test")

	got, _ := d.Search("test")
	want := "no longer a normal test"
	AssertString(t, got, want)
}

func TestDelete(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}
	d.Delete("test")

	_, er := d.Search("test")
	want := NotFoundError

	AssertError(t, er, want)
}

func AssertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertError(t *testing.T, got, want DictionaryError) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
