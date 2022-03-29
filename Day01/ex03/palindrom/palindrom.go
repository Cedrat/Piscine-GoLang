package palindrom

import (
	"fmt"
)

func palindrom(word string) string {
	var rev_str string;

	for i := len(word) - 1 ; i >= 0; i-- {
		rev_str += string(word[i]);
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
			fmt.Println(str_reversed , len(str_reversed));
		}
	}
}