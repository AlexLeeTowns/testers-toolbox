package faker

type Person struct {
	FirstName []string `json:"firstName"`
	LastName  []string `json:"lastName"`
}

func (p *Person) GetFirstName() string {
	fname := selectString(p.FirstName)
	return fname
}

func (p *Person) GetLastName() string {
	lname := selectString(p.LastName)
	return lname
}
