package fixture_test

import (
	"reflect"
	"testing"

	ff "github.com/AlexLeeTowns/testers-toolbox/pkg/fixture"
)

func TestModelMerges(t *testing.T) {
	t.Run("Return with values from user params", func(t *testing.T) {
		f := ff.Factory{}
		got := f.CreateUser(ff.User{FirstName: "Alex"})
		want := ff.User{
			FirstName: "Alex",
			LastName:  "Lorem ipsum",
			Country:   "Lorem ipsum",
		}

		if !reflect.DeepEqual(got, &want) {
			t.Errorf("got %v, want %v", got, &want)
		}
	})
}
