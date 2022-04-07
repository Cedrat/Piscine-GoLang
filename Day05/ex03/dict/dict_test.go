package dict

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func stringDico(word string, define string) string {
	return ("The definition of the word " + word + " is : " + define)
}

func TestDict(T *testing.T) {
	assert := assert.New(T)

	dict := NewDict()

	dict.Create("Earth", "is a planet")
	err := dict.Create("Earth", "is a cat")

	assert.Equal("Word Already Exist", err.Error(), "Msg error need to be Word Already Exist")
	definition, err := dict.Read("Earth")

	assert.Equal(stringDico("Earth", "is a planet"),
		definition,
		"Need to be equal")

	dict.Update("Earth", "Earth is a blue planet")

	definition, err = dict.Read("Earth")

	assert.Equal(stringDico("Earth", "Earth is a blue planet"),
		definition,
		"Need to be equal")

	assert.Equal(stringDico("Earth", "Earth is a blue planet"),
		definition,
		"Need to be equal")

	dict.Delete("Earth")

	definition, err = dict.Read("Earth")

	assert.Equal("Word Not Found", err.Error())
	assert.Equal("dict.DictError", fmt.Sprintf("%T", err), "Need to be dict.DictError type")
	assert.Equal("The word Earth isn't in the dictionnary",
		definition,
		"Need to be equal")
}
