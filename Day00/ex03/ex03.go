
package main

import "bufio"
import "os"
import "ex03/salutations"
import "ex03/lottery"

func main() {

	var f *os.File;
	
	f = os.Stdout;
	
	writer := bufio.NewWriter(f);
	
	writer.WriteString(salutations.English());
	writer.WriteString(salutations.Francais());
	writer.WriteString(salutations.Chinois());
	writer.Flush();

	ticket := lottery.Number("numbers.txt");

	if lottery.Match(lottery.Answer() , &ticket) {
		writer.WriteString("I have a golden ticket !!\n");
	} else {
		writer.WriteString("I'm going to eat cabbage\n");
	}
	writer.Flush();
}
