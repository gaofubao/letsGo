package example

type Animal interface {
	getName() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) getName() string {
	return d.name
}

func (c Cat) getName() string {
	return c.name
}

func compare(d, c Animal) bool {
	return d == c
}


