package faker_test

import (
	"testing"

	"github.com/AlexLeeTowns/testers-toolbox/pkg/faker"
)

var (
	f = faker.CreateNewFaker()
)

func TestPerson(t *testing.T) {
	t.Run("Should retrieve firstname", func(t *testing.T) {
		got := f.Person.GetFirstName()

		if got == "" {
			t.Error("expected first name to be a value")
		}
	})

	t.Run("Should retrieve a last name", func(t *testing.T) {
		got := f.Person.GetFirstName()

		if got == "" {
			t.Error("expected last name to have a not-empty value")
		}
	})
}

func TestCountry(t *testing.T) {
	t.Run("Should retrieve a country code", func(t *testing.T) {
		got := f.Country.GetCountryCode()

		if got == "" {
			t.Error("expected country code to not be empty")
		}
	})
}
