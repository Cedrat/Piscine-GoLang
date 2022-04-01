package fibo

type Fibo struct {
	actual_value int
	next_value   int
}

func NewFibo() (*Fibo, error) {
	var val Fibo
	val.actual_value = 0
	val.next_value = 1

	return &val, nil
}

func (self *Fibo) Next(fibo *Fibo) int {
	current_value := fibo.actual_value
	self.next_value, self.actual_value = self.actual_value+self.next_value, self.next_value

	return current_value
}
