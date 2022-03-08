package loremipsum_test

import (
	"testing"

	l "github.com/AlexLeeTowns/testers-toolbox/pkg/loremipsum"
)

func TestNewRead(t *testing.T) {
	cases := []struct {
		input    string
		quantity int
		want     string
	}{
		{
			input:    "word",
			quantity: 1,
			want:     "Lorem",
		},
		{
			input:    "char",
			quantity: 1,
			want:     "L",
		},
		{
			input:    "paragraph",
			quantity: 1,
			want:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam rhoncus convallis turpis in sagittis. Curabitur eget iaculis elit. Vestibulum faucibus tellus non nisi dapibus, a semper turpis venenatis. Maecenas volutpat elit eu erat consectetur suscipit. Donec at nunc elementum, volutpat dui vitae, vehicula sem. Sed ac diam ac mauris aliquam sollicitudin ac nec mi. In consequat odio vitae mollis mattis.",
		},
	}

	for _, tc := range cases {
		got := l.GetLorem(tc.input, tc.quantity)

		if got != tc.want {
			t.Errorf("Got %q, want %q", got, tc.want)
		}
	}
}
