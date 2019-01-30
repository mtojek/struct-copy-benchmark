package copy

type Dog struct {
	name string
	age int
	breed string
	friends []string
	visited []string

	brain map[string]string
}

func (d *Dog) Copy() *Dog {
	copied := &Dog{
		name: d.name,
		age: d.age,
		breed: d.breed,
		friends: make([]string, len(d.friends)),
		visited: make([]string, len(d.visited)),
		brain: make(map[string]string, len(d.brain)),
	}
	copy(copied.friends, d.friends)
	copy(copied.visited, d.visited)

	for k, v := range d.brain {
		copied.brain[k] = v
	}
	return copied
}