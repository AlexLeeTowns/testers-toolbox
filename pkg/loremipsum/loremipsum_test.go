package loremipsum_test

import (
	"testing"
	"testing/fstest"

	l "github.com/AlexLeeTowns/testers-toolbox/pkg/loremipsum"
)

func TestReadLoremIpsumFile(t *testing.T) {
	t.Run("ReadLoremByCharacterCount reads file by character", func(t *testing.T) {
		fs := fstest.MapFS{
			"fake-file.txt": {Data: []byte("This is a text")},
		}

		got, err := l.ReadLoremByCharacterCount(fs, "fake-file.txt", 3)
		if err != nil {
			t.Fatal(err)
		}
		want := "Thi"

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("ReadLoremByWordCount reads file by word", func(t *testing.T) {
		fs := fstest.MapFS{
			"fake-file.txt": {Data: []byte("These are words")},
		}

		got, err := l.ReadLoremByWordCount(fs, "fake-file.txt", 2)

		if err != nil {
			t.Fatal(err)
		}

		want := "These are"

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("ReadLoremByWordCount reads file by paragraph", func(t *testing.T) {
		fs := fstest.MapFS{
			"fake-file.txt": {Data: []byte(`These are
my final
words`)},
		}

		got, err := l.ReadLoremByParagraph(fs, "fake-file.txt", 2)
		if err != nil {
			t.Fatal(err)
		}

		want := `These are
my final`
		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})
}
