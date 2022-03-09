package fixture

import (
	"github.com/AlexLeeTowns/testers-toolbox/pkg/loremipsum"
)

type Model interface {
	Fill()
}

type User struct {
	FirstName string
	LastName  string
	Country   string
}

type Case struct {
	User
	Id    int
	Title string
}

func (c *Case) Fill() {
	u := User{}
	if c.User == (u) {
		u.Fill()
		c.User = u
	}
	c.Id = intIsMissing(c.Id)
	c.Title = isMissing(c.Title)
}

func (u *User) Fill() {
	u.FirstName = isMissing(u.FirstName)
	u.LastName = isMissing(u.LastName)
	u.Country = isMissing(u.Country)
}

func isMissing(param string) string {
	if param != "" {
		return param
	}
	return loremipsum.GetLorem("word", 2)
}

func intIsMissing(param int) int {
	if param != 0 {
		return param
	}

	return 1
}

func Create(m Model) *Model {
	m.Fill()
	return &m
}
