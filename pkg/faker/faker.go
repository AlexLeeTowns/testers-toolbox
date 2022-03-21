package faker

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed data.json
var dataFile []byte

type Faker struct {
	Person  Person
	Country Country
}

func CreateNewFaker() Faker {
	fml := Faker{}
	if err := json.Unmarshal(dataFile, &fml); err != nil {
		fmt.Printf("unexpected error: %v", err)
	}

	return fml
}
