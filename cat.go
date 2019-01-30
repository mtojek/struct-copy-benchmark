package copy

type Cat struct {
	name string
	age int
	breed string
	friends []string
	visited []string
}

func (c *Cat) Copy() *Cat {
	copied := &Cat{
		name: c.name,
		age: c.age,
		breed: c.breed,
		friends: make([]string, len(c.friends)),
		visited: make([]string, len(c.visited)),
	}
	copy(copied.friends, c.friends)
	copy(copied.visited, c.visited)
	return copied
}