package mr

import "fmt"

type Any interface{}
type Node struct {
	Value Any
	Next  *Node
}

func NewNode(value Any) *Node {
	return &Node{value, nil}
}
func (n *Node) Insert(value Any) {
	tmp := n.Next
	n.Next = NewNode(value)
	if n.Next != nil {
		n.Next.Next = tmp
	}
}
func (n *Node) Remove() {
	if n.Next != nil {
		n.Value = n.Next.Value
		n.Next = n.Next.Next
	} else {
		n.Value = nil
		n.Next = nil
	}
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil}
}
func (l *LinkedList) Append(value Any) {
	if l.Head == nil {
		l.Head = NewNode(value)
		l.Tail = l.Head
	} else {
		l.Tail.Insert(value)
		l.Tail = l.Tail.Next
	}
}

func (l *LinkedList) Print() {
	for tmp := l.Head; tmp != nil; tmp = tmp.Next {
		fmt.Println(tmp.Value)
	}
	fmt.Println()
}

func (l *LinkedList) Map(f func(value Any) Any) (ret *LinkedList) {

	a := NewLinkedList()
	ptr_node := l.Head

	for ptr_node != nil {
		a.Append(f(ptr_node.Value))
		ptr_node = ptr_node.Next
	}
	return a
}

func (l *LinkedList) Reduce(f func(value1, value2 Any) Any) (ret Any) {

	// ptr_node := l.Head
	ptr_node := l.Head
	var b Any

	if ptr_node != nil {
		b = ptr_node.Value
		ptr_node = ptr_node.Next

	}

	for ptr_node != nil {
		b = f(ptr_node.Value, b)
		ptr_node = ptr_node.Next
	}
	return b
}
