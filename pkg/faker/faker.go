package faker

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

//go:embed data.json
var dataFile []byte

type Faker struct {
	Person  *Person
	Country *Country
}

func CreateNewFaker() Faker {
	fml := Faker{}
	rand.Seed(time.Now().Unix())
	if err := json.Unmarshal(dataFile, &fml); err != nil {
		fmt.Printf("cannot unmarshal file into %v\nerror: %v", fml, err)
	}

	return fml
}

func selectString(options []string) string {
	result := fmt.Sprint("", options[rand.Intn(len(options))])

	return result
}
