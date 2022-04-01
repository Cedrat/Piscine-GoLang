package main

import (
	"ex06/ft_queue"
)

func main() {

	a := ft_queue.Ft_queue{}

	a.Push(2)
	a.Push("ceci est une string")
	a.Push(-5)

	a.Display()

	a.Pop()
	a.Display()
	a.Pop()
	a.Pop()
	a.Pop()
	a.Display()

}
