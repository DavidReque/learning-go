package methods

type Car struct {
	Name string
	Year int
}

func (c Car) IsLatest() bool {
	return c.Year == 2023
}

func (c *Car) UpdateName(name string) {
	c.Name = name
}