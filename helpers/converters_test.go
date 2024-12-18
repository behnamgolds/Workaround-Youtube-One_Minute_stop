package helpers_test

import (
	"testing"

	"yt-getembed/helpers"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			"https://www.youtube.com/watch?v=EngW7tLk6R8",
			"https://www.youtube.com/embed/EngW7tLk6R8",
		},
		{
			"https://www.youtube.com/watch?v=LMvSu8kdysk",
			"https://www.youtube.com/embed/LMvSu8kdysk",
		},
	}

	for _, cs := range cases {
		got, _ := helpers.Convert(cs.input)
		if got != cs.want {
			t.Errorf("Convert(%q): want %q, got %q", cs.input, cs.want, got)
		}
	}
}
