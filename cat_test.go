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
	assert.Equal(t, *cat, *result)
}

func TestCat_Copy_Equal_Set_SlicesUntouched(t *testing.T) {
	// given
	cat := &Cat{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
	}

	// when
	result := *cat

	// then
	assert.Equal(t, *cat, result)
}

func TestCat_Copy_NotEqual_JustSet_SlicesTouched(t *testing.T) {
	// given
	cat := &Cat{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
	}

	// when
	result := *cat
	result.friends = append(result.friends, "James")
	cat.visited[0] = "foo"

	// then
	assert.NotEqual(t, *cat, result)
	assert.Len(t, cat.friends, 2)
	assert.Len(t, result.friends,3)
	assert.Equal(t, "foo", cat.visited[0])
	assert.Equal(t, "foo", result.visited[0]) // But expected "garden"
}

func TestCat_Copy_NotEqual_Copy_SlicesTouched(t *testing.T) {
	// given
	cat := &Cat{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
	}

	// when
	var result = cat.Copy()
	result.friends = append(result.friends, "James")
	cat.visited[0] = "foo"

	// then
	assert.NotEqual(t, *cat, *result)
	assert.Len(t, cat.friends, 2)
	assert.Len(t, result.friends,3)
	assert.Equal(t, "foo", cat.visited[0])
	assert.Equal(t, "garden", result.visited[0]) // But expected "garden"
}

func TestCat_Copy_NotEqual_CopierCopy_SlicesTouched(t *testing.T) {
	// given
	cat := &Cat{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
	}

	// when
	var result = &Cat{}
	//reflect.ValueOf(result).Elem().Set(reflect.ValueOf(cat).Elem())
	copier.Copy(result, cat)
	result.friends = append(result.friends, "James")
	cat.visited[0] = "foo"

	// then
	assert.NotEqual(t, *cat, *result)
	assert.Len(t, cat.friends, 2)
	assert.Len(t, result.friends,3)
	assert.Equal(t, "foo", cat.visited[0])
	assert.Equal(t, "foo", result.visited[0]) // But expected "garden"
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
	assert.NotEqual(t, *cat, *result)
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

func BenchmarkCat_Copy_JustSet_ButItsWrong(b *testing.B) {
	cat := &Cat{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
	}

	for n := 0; n < b.N; n++ {
		dest := &Cat{}
		*dest = *cat
	}
}

func BenchmarkCat_Copy_CopierCopy(b *testing.B) {
	cat := &Cat{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
	}

	for n := 0; n < b.N; n++ {
		dest := &Cat{}
		copier.Copy(dest, cat)
	}
}