// main.go
package main

import (
	"ex00/mr"

	"fmt"
)

func main() {
	l := mr.NewLinkedList()
	l.Append("1")
	l.Append(2)
	l.Append("3")
	l.Append(4)
	l.Print()
	l2 := l.Map(mr.AddOneInt)
	l2.Print()
	l3 := l.Map(mr.AddOneString)
	l3.Print()
	fmt.Println(l.Reduce(mr.SumInt))
	fmt.Println(l2.Reduce(mr.SumInt))
	fmt.Println(l3.Reduce(mr.SumInt))
}
