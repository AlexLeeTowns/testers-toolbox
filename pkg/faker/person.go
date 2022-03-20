package faker

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	FirstName []string
	LastName  []string
}

func (p Person) GetFirstName() string {
	rand.Seed(time.Now().Unix())
	fname := fmt.Sprint("", p.FirstName[rand.Intn(len(p.FirstName))])

	return fname
}
