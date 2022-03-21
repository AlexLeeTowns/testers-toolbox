package faker

import (
	"fmt"
	"math/rand"
	"time"
)

type Country struct {
	Code []string
}

func (c Country) GetCountryCode() string {
	rand.Seed(time.Now().Unix())
	code := fmt.Sprint("", c.Code[rand.Intn(len(c.Code))])

	return code
}
