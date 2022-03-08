package fixture

import (
	"fmt"

	"github.com/AlexLeeTowns/testers-toolbox/pkg/loremipsum"
)

// TODO: Make a fixture factory

type Model interface {
	Fml()
}

type FixtureFactory interface {
	Create(m Model) Model
}

type User struct {
	FirstName string
	LastName  string
	Country   string
}

func (u User) Fml() {
	fmt.Print("fml")
}

type Factory struct {
}

func isMissing(param string) string {
	if param != "" {
		return param
	}
	return loremipsum.GetLorem("word", 2)
}

func (f *Factory) CreateUser(u User) *User {

	return &User{
		FirstName: isMissing(u.FirstName),
		LastName:  isMissing(u.LastName),
		Country:   isMissing(u.Country),
	}
}
