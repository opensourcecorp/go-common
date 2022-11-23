package osc

import "testing"

func TestDedent(t *testing.T) {
	t.Run("dedent a source-indented string", func(t *testing.T) {
		s := `
		abc
			easy as
				one two three
		`

		want := `abc
	easy as
		one two three`
		got := Dedent(s)

		if want != got {
			t.Errorf("\nwant: %q\ngot:  %q", want, got)
		}
	})

	t.Run("return the same string if already source-dedented", func(t *testing.T) {
		s := `abc
	easy as
		one two three`
		want := s
		got := Dedent(s)

		if want != got {
			t.Errorf("\nwant: %q\ngot:  %q", want, got)
		}
	})
}

func TestCountLeadingWhitespace(t *testing.T) {
	t.Run("string has leading whitespace", func(t *testing.T) {
		s := "  abc"

		want := 2
		got := countLeadingWhitespace(s)

		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("string has no leading whitespace", func(t *testing.T) {
		s := "abc"

		want := 0
		got := countLeadingWhitespace(s)

		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("string has leading whitespace and more than one word", func(t *testing.T) {
		s := "  x y z"

		want := 2
		got := countLeadingWhitespace(s)

		if want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})
}
