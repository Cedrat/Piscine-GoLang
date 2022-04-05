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

	dict := make(map[string]string)

	Create(dict, "Earth", "is a planet")
	Create(dict, "Earth", "is a cat")

	assert.Equal(stringDico("Earth", "is a planet"),
		Read(dict, "Earth"),
		"Need to be equal")

	Update(dict, "Earth", "Earth is a blue planet")
	assert.Equal(stringDico("Earth", "Earth is a blue planet"),
		Read(dict, "Earth"),
		"Need to be equal")

	Delete(dict, "Earth")
	assert.Equal("The word Earth isn't in the dictionnary",
		Read(dict, "Earth"),
		"Need to be equal")
}
