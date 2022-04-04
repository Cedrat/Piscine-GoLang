package dict

import "fmt"

func Create(dictionary map[string]string, word string, definition string) {
	if _, ok := dictionary[word]; ok {
		fmt.Println(word, "is currently in the dictionnary")
	} else {
		dictionary[word] = definition
		fmt.Println("Create the entry for the word :", word)
	}
}

func Read(dictionary map[string]string, word string) string {
	if _, ok := dictionary[word]; ok {
		return "The definition of word is " + dictionary[word]
	} else {
		return "The word " + word + "isn't in the dictionnary"
	}

}

func Update(dictionary map[string]string, word string, definition string) {
	if _, ok := dictionary[word]; ok {
		dictionary[word] = definition
	} else {
		fmt.Println("The word doesn't exist, can't be updated")
	}
}

func Delete(dictionary map[string]string, word string) {
	if _, ok := dictionary[word]; ok {
		delete(dictionary, word)
	} else {
		fmt.Println("The word can't be deleted because he doesn't exist")
	}
}
