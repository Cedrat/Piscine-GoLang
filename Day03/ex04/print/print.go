package print

import (
	"fmt"
	"io"
)

type Printer struct {
	Name    string
	doPrint func(w io.Writer)
	cnt     int
}

func NewPrinter(name string) *Printer {
	return &Printer{name, Waiting, 0}
}
func (p *Printer) Print(w io.Writer) {
	p.cnt++
	defer p.doPrint(w)
	if p.cnt == 4 {
		panic("machine broken")
	}
}
func (p *Printer) SetPrint(f func(w io.Writer)) {
	p.doPrint = f
}
func Printing(w io.Writer) {
	fmt.Fprintln(w, "Printing...")
}
func Waiting(w io.Writer) {
	fmt.Fprintln(w, "Waiting...")
}

func (p *Printer) Wrapper(f func(w io.Writer)) (s func(w io.Writer)) {
	nb_fix := 0

	return func(w io.Writer) {
		if err := recover(); err != nil {
			p.cnt = 0
			nb_fix++
			fmt.Fprintln(w, p.Name+": "+"printer is fixed", nb_fix, "times")
		}
		fmt.Fprint(w, p.Name+": ")
		f(w)
	}
}
