package dict

import (
	"fmt"
)

type Dict struct {
	dict map[string]string
}

func NewDict() *Dict {
	var dict Dict

	dict.dict = make(map[string]string)

	return &dict
}

func (dict *Dict) Create(word string, definition string) error {
	if _, ok := dict.dict[word]; ok {
		return ErrorAlreadyExist
	} else {
		dict.dict[word] = definition
		fmt.Println("Create the entry for the word :", word)
	}
	return nil
}

func (dict *Dict) Read(word string) (string, error) {
	if _, ok := dict.dict[word]; ok {
		return "The definition of the word " + word + " is : " + dict.dict[word], nil
	} else {
		return "The word " + word + " isn't in the dictionnary", ErrorNotFound
	}
}

func (dict *Dict) Update(word string, definition string) error {
	if _, ok := dict.dict[word]; ok {
		dict.dict[word] = definition
	} else {
		return ErrorDoesNotExist
	}
	return nil
}

func (dict *Dict) Delete(word string) error {
	if _, ok := dict.dict[word]; ok {
		delete(dict.dict, word)
	} else {
		return ErrorNoRemovable
	}
	return nil
}
