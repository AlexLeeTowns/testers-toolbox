package fixture_test

import (
	"reflect"
	"testing"

	ff "github.com/AlexLeeTowns/testers-toolbox/pkg/fixture"
)

func TestModelMerges(t *testing.T) {
	t.Run("Return User with values from user params", func(t *testing.T) {
		got := ff.User{FirstName: "Alex"}
		got.Fill()
		want := ff.User{
			FirstName: "Alex",
			LastName:  "Lorem ipsum",
			Country:   "Lorem ipsum",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, &want)
		}
	})

	t.Run("Fill with integers", func(t *testing.T) {
		got := ff.Case{Title: "First case"}
		got.Fill()
		want := ff.Case{
			Id:    1,
			Title: "First case",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Handle dependencies", func(t *testing.T) {
		got := ff.Case{Title: "First case"}
		got.Fill()
		want := ff.User{
			FirstName: "Lorem Ipsum",
			LastName:  "Lorem Ipsum",
			Country:   "Lorem Ipsum",
		}

		if !reflect.DeepEqual(got.User, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
}
