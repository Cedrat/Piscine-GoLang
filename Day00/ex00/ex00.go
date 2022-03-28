package main

import "os"
import "bufio"

func main() {
	var f *os.File;
	
	f = os.Stdout;
	
	writer := bufio.NewWriter(f);
	
	writer.WriteString("Welcome to GoLang\n");

	writer.Flush();
}