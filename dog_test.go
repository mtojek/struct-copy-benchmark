package copy

import (
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDog_Copy_Equal(t *testing.T) {
	// given
	dog := &Dog{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
		brain: map[string]string{
			"food": "meat",
			"play": "ball",
			"walk": "joy",
		},
	}

	// when
	result := dog.Copy()

	// then
	assert.Equal(t, dog, result)
}

func TestDog_Copy_NotEqual(t *testing.T) {
	// given
	dog := &Dog{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
		brain: map[string]string{
			"food": "meat",
			"play": "ball",
			"walk": "joy",
		},
	}

	// when
	result := dog.Copy()
	dog.brain["play"] = "running"

	// then
	assert.NotEqual(t, dog, result)
}

func BenchmarkDog_Copy_CopyMethod(b *testing.B) {
	dog := &Dog{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
		brain: map[string]string{
			"food": "meat",
			"play": "ball",
			"walk": "joy",
		},
	}

	for n := 0; n < b.N; n++ {
		dog.Copy()
	}
}


func BenchmarkDog_Copy_WithCopier(b *testing.B) {
	dog := &Dog{
		name: "Mr Wiggles",
		age: 3,
		breed: "Persian",
		friends: []string{"Jerry", "Tom"},
		visited: []string{"garden", "park", "shop"},
		brain: map[string]string{
			"food": "meat",
			"play": "ball",
			"walk": "joy",
		},
	}

	for n := 0; n < b.N; n++ {
		dest := &Dog{}
		copier.Copy(dog, dest)
	}
}