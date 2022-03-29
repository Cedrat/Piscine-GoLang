package main
import (
	"ex02/deepcopy"
)
func main() {
	a := []int{};
	deepcopy.DeepCopy(a, 3, 2);
}