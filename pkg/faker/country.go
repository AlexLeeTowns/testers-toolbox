package faker

type Country struct {
	Code []string
}

func (c *Country) GetCode() string {
	code := selectString((c.Code))
	return code
}
