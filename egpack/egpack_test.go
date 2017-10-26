package egpack

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestReverse_EvenNumberOfRunes(t *testing.T) {
	if Reverse("abcd") != "dcba" {
		t.Errorf("Reverse(%q) == %q, want %q", "abcd", Reverse("abcd"), "dcba")
	}
}

func TestReverse_OddNumberOfRunes(t *testing.T) {
	if Reverse("abc") != "cba" {
		t.Errorf("Reverse(%q) == %q, want %q", "abc", Reverse("abc"), "cba")
	}
}
