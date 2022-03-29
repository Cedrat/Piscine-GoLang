package palindrom

import (
	"fmt"
	"unicode/utf8"
)

func palindrom(word string) string {
	var rev_str string;

	var runes []rune;

	// for i := len(word) - 1 ; i >= 0; i-- {
	// 	rev_str += string(word[i]);
	// }

	for _, char := range word {
		runes = append(runes, char);
	}

	for i := len(runes) - 1; i >= 0; i-- {
		rev_str += string(runes[i]);
	}


	for i:= 0; i < len(word); i++ {
		for p:= 0; p < len(rev_str); p++ {
			if word[i + p] != rev_str[p] {
				break;
			}
			if i + p  == len(word) - 1 {
				return (word + rev_str[p + 1:]);
			}
		}
	}

	return word + rev_str;
}

func Palindrom(words ...string) {
	for _, word := range words {
		if len(word) != 0 {
			str_reversed := palindrom(word);
			fmt.Println(str_reversed , utf8.RuneCountInString(str_reversed));
		}
	}
}