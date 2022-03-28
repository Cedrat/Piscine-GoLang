
package main

import "bufio"
import "os"
import "ex03/salutations"

func main() {

	var f *os.File;
	
	f = os.Stdout;
	
	writer := bufio.NewWriter(f);
	
	writer.WriteString(salutations.English());
	writer.WriteString(salutations.Francais());
	writer.WriteString(salutations.Chinois());
	writer.Flush();

}
