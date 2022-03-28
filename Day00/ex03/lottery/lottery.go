package lottery

import (
	"os"
	"bufio"
	"strconv"
	"time"
	"math/rand"
)

func PutError(msg error) {
	var f *os.File;
	
	f = os.Stderr;
	
	writer := bufio.NewWriter(f);
	writer.WriteString(msg.Error());
	writer.Flush();
}

func PutStr(msg string) {
	var f *os.File;
	
	f = os.Stdout;
	
	writer := bufio.NewWriter(f);
	writer.WriteString(msg);
	writer.Flush();
}

func IsFourNumbers(numbers []byte) bool {
	
	nb_white_space := 0;

	for _, char := range numbers {
		if char == ' ' {
			nb_white_space++;
		} else if !((char <= 'a' && char < 'z') || (char <= 'A' && char < 'Z')) {
			return false;
		}
	}
	if nb_white_space != 3 {
		return false;
	}
	return true;
}
/*
Worst Parsing but Split was forbidden.
*/
func ParseNumbers(numbers []byte) (arr[4] uint) {
	nb_white_space := 0;
	var array_nb [4]uint;

	str_nb := string(numbers);

	for i, char := range numbers {
		if char == ' ' || i == 0 {
			if i != 0 {
				i++;
			}
			if (numbers[i + 1] == ' ') {
				number, _ := strconv.Atoi(str_nb[i: i + 1]);
				array_nb[nb_white_space] = uint(number);
			} else {
				number, _ := strconv.Atoi(str_nb[i: i + 2]);
				array_nb[nb_white_space] = uint(number);
			}
			nb_white_space++;
		}
	}
	return array_nb;
}

func Number(path_numbers string) (arr[4] uint) {
	file, err := os.Open(path_numbers);
	a := [4]uint{100};
	if err != nil {
		PutError(err);
		return a;
	}

	data:= make([]byte, 16);
	_, err = file.Read(data);

	file.Close();
	
	if err != nil {
		PutError(err);
		return a;
	}

	if IsFourNumbers(data) { 
		return ParseNumbers(data);
 	} else {
		 PutStr("Content of the file incorrect.\n")
		 return a;
	 }  
}

func Answer() uint {

	generator := rand.New(rand.NewSource(time.Now().UnixNano()));
	nb_ticket := uint(generator.Uint64() % 100);
	PutStr("The great ticket is " + strconv.Itoa(int(nb_ticket)) + "\n")

	return nb_ticket;
}

func Match(answer uint, ticket *[4]uint) bool {
	for i:=0; i < 4 ; i++ {
		if answer == ticket[i] {
			return true;
		}
	}
	return false;
}