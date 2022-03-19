package fixture

import (
	"encoding/json"
	"fmt"

	"github.com/AlexLeeTowns/testers-toolbox/pkg/loremipsum"
)

type UserData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type User struct {
	UserData
}

// TODO: Figure out how to handle 0 values
func CreateUser(firstName, lastName string) (*User, error) {
	user := User{}
	data, err := Create(user)
	if err != nil {
		fmt.Printf("error during User creation: %v", err)
		return nil, err
	}

	json.Unmarshal(data, &user.UserData)
	return &user, nil
}

func (u User) path() string {
	return "/user"
}

func (u User) json() ([]byte, error) {
	uData := UserData{
		FirstName: loremipsum.GetLorem("word", 1),
		LastName:  loremipsum.GetLorem("word", 1),
	}

	data, err := json.Marshal(uData)
	if err != nil {
		return nil, err
	}

	return data, nil
}
