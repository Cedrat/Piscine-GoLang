package main

import (
	"ex04/print"
	"os"
)

func main() {
	p := print.NewPrinter("printer1")
	p.SetPrint(p.Wrapper(print.Waiting))
	p.Print(os.Stdout)
	p.Print(os.Stdout)
	p.SetPrint(p.Wrapper(print.Printing))
	p.Print(os.Stdout)
	p.Print(os.Stdout)
	p.Print(os.Stdout)
	p.Print(os.Stdout)
	p.Print(os.Stdout)
	p.Print(os.Stdout)
}
