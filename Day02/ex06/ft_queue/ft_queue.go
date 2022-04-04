package ft_queue

import "fmt"

type Ft_queue struct {
	under_queue []little_element
}

type little_element struct {
	element Any
}

func Newlittle_element(brick Any) (little_element, error) {
	var a little_element

	a.element = brick

	return a, nil
}

func (queue *Ft_queue) Push(a Any) {
	b, _ := Newlittle_element(a)
	queue.under_queue = append(queue.under_queue, b)

}

func (queue *Ft_queue) Empty() bool {
	if len(queue.under_queue) == 0 {
		return true
	}
	return false
}

func (queue *Ft_queue) Len() int {
	return len(queue.under_queue)
}

func (queue *Ft_queue) Pop() Any {
	if queue.Len() == 0 {
		return nil
	}
	queue.under_queue = queue.under_queue[:queue.Len()-1]
	if queue.Len() == 0 {
		return nil
	}
	return queue.under_queue[0].element
}

func (queue *Ft_queue) Display() {
	for i := 0; i < len(queue.under_queue); i++ {
		fmt.Println(queue.under_queue[i].element)
	}
}
