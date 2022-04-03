package dict

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func stringDico(word string, define string) string {
	return ("The definition of the word " + word + " is : " + define)
}

func TestDict(T *testing.T) {
	assert := assert.New(T)

	var dict map[string]string

	Create(dict, "Earth", "is a planet")
	Create(dict, "Earth", "is a cat")

	assert.Equal(Read(dict, "Earth"),
		stringDico("Earth", "is a planet"),
		"Need to be equal")

	Update(dict, "Earth", "Earth is a blue planet")
	assert.Equal(Read(dict, "Earth"),
		stringDico("Earth", "Earth is a blue planet"),
		"Need to be equal")

	Delete(dict, "Earth")
	assert.Equal(Read(dict, "Earth"),
		"The word Earth isn't in the dictionnary",
		"Need to be equal")
}
