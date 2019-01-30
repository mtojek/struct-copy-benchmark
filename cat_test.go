package copy

import (
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCat_Copy_Equal(t *testing.T) {
	// given
	cat := &Cat{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
	}

	// when
	result := cat.Copy()

	// then
	assert.Equal(t, cat, result)
}

func TestCat_Copy_NotEqual(t *testing.T) {
	// given
	cat := &Cat{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
	}

	// when
	result := cat.Copy()
	cat.visited = append(cat.visited, "forest")

	// then
	assert.NotEqual(t, cat, result)
}

func BenchmarkCat_Copy_CopyMethod(b *testing.B) {
	cat := &Cat{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
	}

	for n := 0; n < b.N; n++ {
		cat.Copy()
	}
}


func BenchmarkCat_Copy_WithCopier(b *testing.B) {
	cat := &Cat{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
	}


	for n := 0; n < b.N; n++ {
		dest := &Cat{}
		copier.Copy(cat, dest)
	}
}