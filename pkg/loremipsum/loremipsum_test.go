package loremipsum_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	l "github.com/AlexLeeTowns/testers-toolbox/pkg/loremipsum"
)

func TestReadLorem(t *testing.T) {
	cases := []struct {
		input    string
		quantity int
		want     string
	}{
		{
			input:    "word",
			quantity: 1,
			want:     "These",
		},
		{
			input:    "char",
			quantity: 1,
			want:     "T",
		},
		{
			input:    "paragraph",
			quantity: 2,
			want: `These
are`,
		},
	}
	fs := fstest.MapFS{
		"fake.txt": {Data: []byte(`These
are
words`)},
	}

	for _, tc := range cases {
		got, err := l.ReadLorem(fs, "fake.txt", tc.input, tc.quantity)
		if err != nil {
			t.Errorf("Threw and error, but shouldn't: %v", err)
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Got %q, want %q", got, tc.want)
		}
	}
}
